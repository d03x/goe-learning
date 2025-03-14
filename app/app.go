package app

import (
	"elearning/app/config"
	"elearning/app/connection"
	"embed"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"log/slog"
)

type RouteSetupCallback func(a *fiber.App, db *goqu.Database)
type FuncMigrationHandler func(db *goqu.Database) error

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
	db     *goqu.Database
	server *fiber.App
	config *config.Config
}

func NewApp(configParam Config) App {
	newConfig := config.GetConfig()
	db := connection.DatabaseConnection(newConfig.DB)
	return &app{
		db:     goqu.New("mysql", db),
		server: fiber.New(configParam.Fiber),
		config: newConfig,
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
			slog.Error(`html index not found! Please build /client using pnpm run build and recompile application`)
			return ctx.Status(fiber.StatusNotFound).SendString("public/index.html not found")
		}
		return ctx.Type("html").Send(html)
	})
}
func (a *app) InitMigrate(handler FuncMigrationHandler) {
	err := handler(a.db)
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
	routeCallback(a.server, a.db)
}

func (a *app) Run() error {
	return a.server.Listen(":8080")
}
