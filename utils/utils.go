package utils

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
)

func Response(data interface{}, httpStatus int, err error, c *fiber.Ctx) error {

	if err != nil {
		if errors.Is(err, platform.ErrNotFound) {
			return c.Status(404).JSON(map[string]string{
				"error": err.Error(),
			})
		} else if reqErr, ok := err.(platform.ErrorResponse); ok {
			return c.Status(reqErr.StatusCode).JSON(map[string]string{
				"error": err.Error(),
			})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": err.Error(),
			})
		}

	} else {
		if data != nil {
			return c.Status(httpStatus).JSON(data)
		} else {
			c.Status(httpStatus)
			return nil
		}
	}
}
