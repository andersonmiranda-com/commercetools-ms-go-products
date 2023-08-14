package router

import (
	"commercetools-ms-product/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())

	api := app.Group("/api/products")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/", controller.Find)
	api.Get("/:id", controller.Get)
	api.Post("/", controller.Create)
	api.Put("/:id", controller.Update)

	api.Patch("/publish/:id", func(c *fiber.Ctx) error {
		return controller.SetPublishStatus("publish", c)
	})

	api.Patch("/unpublish/:id", func(c *fiber.Ctx) error {
		return controller.SetPublishStatus("unpublish", c)
	})

	api.Delete("/:id", controller.Remove)

}
