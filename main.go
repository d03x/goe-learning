package main

import (
	"elearning/app"
	"elearning/app/api"
	"elearning/app/repositories"
	"embed"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v3"
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
	newApp.SetupRoute(func(a *fiber.App, db *goqu.Database) {
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
