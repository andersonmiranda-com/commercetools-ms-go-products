package service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cts Service) {

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())

	api := app.Group("/")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/:id", cts.Get)
	api.Get("/", cts.Find)
	api.Post("/", cts.Create)
	api.Put("/:id", cts.Update)
	api.Delete("/:id", cts.Remove)

	api.Patch("/publish/:id", func(c *fiber.Ctx) error {
		return cts.SetPublishStatus("publish", c)
	})

	api.Patch("/unpublish/:id", func(c *fiber.Ctx) error {
		return cts.SetPublishStatus("unpublish", c)
	})

}
