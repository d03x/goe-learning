package app

import (
	"elearning/app/config"
	"elearning/app/connection"
	"embed"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/gorm"
	"log/slog"
)

type RouteSetupCallback func(a *fiber.App, db *gorm.DB)
type FuncMigrationHandler func(db *gorm.DB) error

type Config struct {
	Fiber fiber.Config
}
type App interface {
	Run() error
	InitMigrate(handler FuncMigrationHandler)
	SetupRoute(routeCallback RouteSetupCallback)
	SetupFrontEnd(publicPath embed.FS)
}
type app struct {
	database *gorm.DB
	server   *fiber.App
	config   *config.Config
}

func NewApp(configParam Config) App {
	newConfig := config.GetConfig()
	db := connection.DatabaseConnection(newConfig.DB)
	return &app{
		database: db,
		server:   fiber.New(configParam.Fiber),
		config:   newConfig,
	}
}
func (a *app) SetupFrontEnd(publicPath embed.FS) {
	a.server.Get("/*", static.New("", static.Config{
		FS:     publicPath,
		Browse: false,
	}))
	//run frontend html file
	a.server.Get("/", func(ctx fiber.Ctx) error {
		html, err := publicPath.ReadFile("public/index.html")
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString("Public/index.html not frontend")
		}
		return ctx.Type("html").Send(html)
	})
}
func (a *app) InitMigrate(handler FuncMigrationHandler) {
	err := handler(a.database)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (a *app) SetupRoute(routeCallback RouteSetupCallback) {
	a.server.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone:   "Asia/Jakarta",
		TimeFormat: "02-Jan-2006",
	}))
	routeCallback(a.server, a.database)
}

func (a *app) Run() error {
	return a.server.Listen(":8080")
}
