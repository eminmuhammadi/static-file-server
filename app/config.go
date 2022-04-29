package app

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

var APP_CONFIG = fiber.Config{
	Prefork:                   true,
	StrictRouting:             true,
	CaseSensitive:             true,
	GETOnly:                   true,
	DisableDefaultContentType: true,
	EnablePrintRoutes:         true,
	ReduceMemoryUsage:         true,
	BodyLimit:                 4 * 1024 * 1024,
	Concurrency:               256 * 1024,
	ReadTimeout:               time.Second * 15,
	WriteTimeout:              time.Second * 15,
	IdleTimeout:               time.Second * 60,
	ReadBufferSize:            4096,
	WriteBufferSize:           4096,
	DisableStartupMessage:     true,
}
