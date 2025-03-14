package main

import (
	"elearning/app"
	"elearning/app/api"
	"elearning/app/models"
	"elearning/app/repositories"
	"embed"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"log/slog"
)

//go:embed public/*
var publicPath embed.FS

func main() {
	config := app.Config{
		Fiber: fiber.Config{
			AppName: "E Learning App",
		},
	}
	newApp := app.NewApp(config)
	newApp.InitMigrate(func(db *gorm.DB) error {
		err := db.AutoMigrate(&models.Products{}, &models.Users{})
		return err
	})
	newApp.SetupRoute(func(a *fiber.App, db *gorm.DB) {
		//handler public file
		userRepo := repositories.NewUser(db)
		api.NewAuth(a, userRepo)
	})
	//use embed folder using go:embed
	newApp.SetupFrontEnd(publicPath)
	if err := newApp.Run(); err != nil {
		slog.Error(err.Error())
	}
}
