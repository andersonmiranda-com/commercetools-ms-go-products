package main

import (
	"commercetools-ms-product/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	ctService := service.NewService()
	app := fiber.New()
	app.Use(logger.New())
	service.SetupRoutes(app, ctService)
	log.Fatal(app.Listen(":4444"))
}
