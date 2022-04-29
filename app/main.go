package app

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Init app
func Init() *fiber.App {
	app := fiber.New(APP_CONFIG)

	return useMiddleware(app)
}

// Middleware
func useMiddleware(app *fiber.App) *fiber.App {
	// Etag middleware
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))

	// Compress
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS",
	}))

	// Recover middleware
	app.Use(recover.New())

	return app
}

// Run app
func Run(port string, staticPath string) error {
	app := Init()

	// Static files
	app.Static("/", staticPath, fiber.Static{
		Browse:    true,
		Download:  false,
		Compress:  true,
		ByteRange: true,
	})

	return app.Listen(fmt.Sprintf(":%s", port))
}
