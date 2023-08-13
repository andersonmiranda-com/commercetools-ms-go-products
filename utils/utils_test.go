package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func mockFiberCtx(query string) *fiber.Ctx {
	app := fiber.New()
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.SetRequestURI("/?" + query)
	return app.AcquireCtx(ctx)
}

func TestGetWhereClause(t *testing.T) {
	ctx := mockFiberCtx("name=testName&slug=testSlug&key=testKey&extraParam=extraValue")

	whereClause := GetWhereClause(ctx)

	assert.Contains(t, whereClause, "masterData(current(name(en=\"testName\")))")
	assert.Contains(t, whereClause, "masterData(current(slug(en=\"testSlug\")))")
	assert.Contains(t, whereClause, "key=\"testKey\"")
	assert.Contains(t, whereClause, "extraParam=extraValue")
}

func TestGenerateWhereClause(t *testing.T) {
	whereClause := generateWhereClause("testName", "testSlug", "testKey", "en")

	assert.Contains(t, whereClause, "masterData(current(name(en=\"testName\")))")
	assert.Contains(t, whereClause, "masterData(current(slug(en=\"testSlug\")))")
	assert.Contains(t, whereClause, "key=\"testKey\"")
}

func TestConvertToSlice(t *testing.T) {
	data := "[test1,test2]"
	res, err := ConvertToSlice(data)
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1", "test2"}, res)

	data = "test1"
	res, err = ConvertToSlice(data)
	assert.Nil(t, err)
	assert.Equal(t, []string{"test1"}, res)
}

func TestResponse(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return Response("testdata", 200, nil, c)
	})

	req, _ := http.NewRequest("GET", "/", bytes.NewBuffer(nil))
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	app.Get("/error", func(c *fiber.Ctx) error {
		return Response(nil, 500, fmt.Errorf("test error"), c)
	})

	req, _ = http.NewRequest("GET", "/error", bytes.NewBuffer(nil))
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
