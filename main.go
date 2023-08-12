package main

import (
	"commercetools-ms-product/config"
	"commercetools-ms-product/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + config.Getenv("PORT")))
}
