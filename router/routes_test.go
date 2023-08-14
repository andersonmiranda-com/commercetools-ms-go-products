package router

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()

	SetupRoutes(app)

	t.Run("check health route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/products/health", nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)

		body := string(bodyBytes)
		expectedBody := `{"health":"ok","status":200}`
		assert.Equal(t, expectedBody, body)
	})

	// TODO: Write tests for other endpoints.
}
