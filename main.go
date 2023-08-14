package main

import (
	"commercetools-ms-product/router"
	"commercetools-ms-product/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	service.TestConnection()
	app := fiber.New()
	app.Use(logger.New())
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4444"))
}
