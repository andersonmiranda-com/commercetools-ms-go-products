package service

import (
	"commercetools-ms-product/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (cts *ctService) Get(c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodGetInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	productResults, err := cts.Connection.Products().WithId(id).Get().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func (cts *ctService) Find(c *fiber.Ctx) error {

	queryArgs := platform.ByProjectKeyProductsRequestMethodGetInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	productResults, err := cts.Connection.Products().Get().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func (cts *ctService) Create(c *fiber.Ctx) error {

	queryArgs := platform.ByProjectKeyProductsRequestMethodPostInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productDraft := platform.ProductDraft{}

	if err := c.BodyParser(&productDraft); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	productResults, err := cts.Connection.Products().Post(productDraft).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func (cts *ctService) Update(c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodPostInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productUpdate := platform.ProductUpdate{}

	if err := c.BodyParser(&productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ProductUpdate",
		})
	}

	ctx := context.Background()
	productResults, err := cts.Connection.Products().WithId(id).Post(productUpdate).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func (cts *ctService) Remove(c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()

	if queryArgs.Version == 0 {
		// Get oldProduct version
		oldProduct, err := cts.Connection.Products().WithId(id).Get().Execute(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		queryArgs.Version = oldProduct.Version
	}

	productResults, err := cts.Connection.Products().WithId(id).Delete().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func (cts *ctService) SetPublishStatus(action string, c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodPostInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()

	// Get oldProduct version
	oldProduct, err := cts.Connection.Products().WithId(id).Get().Execute(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	actionPayload := `{
		"version" : ` + strconv.Itoa(oldProduct.Version) + `,
		"actions" : [ {
		  "action" : "` + action + `"
		} ]
	  }`

	productUpdate := platform.ProductUpdate{}

	if err := json.Unmarshal([]byte(actionPayload), &productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productResults, err := cts.Connection.Products().WithId(id).Post(productUpdate).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}
