package service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, service Service) {

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

	api.Get("/:id", service.Get)
	api.Get("/", service.Find)
	api.Post("/", service.Create)
	api.Put("/:id", service.Update)
	api.Delete("/:id", service.Remove)

	// api.Patch("/publish/:id", func(c *fiber.Ctx) error {
	// 	return controller.SetPublishStatus("publish", c)
	// })

	// api.Patch("/unpublish/:id", func(c *fiber.Ctx) error {
	// 	return controller.SetPublishStatus("unpublish", c)
	// })

}
