package main

import (
	"shortened-url/config"
	routeHandler "shortened-url/handler/routes"
	"shortened-url/middleware"
	routeRepo "shortened-url/repository/routes"

	"github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func main() {
	cfg := config.Get()
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.CustomErrorHandler,
	})

	routesRepo := routeRepo.New(cfg)
	routesHandler := routeHandler.New(routesRepo, validate)

	app.Get("/all", routesHandler.All)
	app.Get("/:url", routesHandler.Get)
	app.Post("/shorten", routesHandler.Shorten)

	app.Listen(":3000")
}
