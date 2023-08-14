package controller

import (
	"commercetools-ms-product/service"
	"commercetools-ms-product/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
)

func Get(c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodGetInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectClient, ctx := service.Connector()
	productResults, err := projectClient.Products().WithId(id).Get().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Find(c *fiber.Ctx) error {

	queryArgs := platform.ByProjectKeyProductsRequestMethodGetInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectClient, ctx := service.Connector()
	productResults, err := projectClient.Products().Get().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Create(c *fiber.Ctx) error {

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

	projectClient, ctx := service.Connector()
	productResults, err := projectClient.Products().Post(productDraft).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Update(c *fiber.Ctx) error {

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

	projectClient, ctx := service.Connector()
	productResults, err := projectClient.Products().WithId(id).Post(productUpdate).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func SetPublishStatus(action string, c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodPostInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Get oldProduct version
	projectClient, ctx := service.Connector()
	oldProduct, err := projectClient.Products().WithId(id).Get().Execute(ctx)
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

	productResults, err := projectClient.Products().WithId(id).Post(productUpdate).WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Remove(c *fiber.Ctx) error {

	id := c.Params("id")

	queryArgs := platform.ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	if err := c.QueryParser(&queryArgs); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if queryArgs.Version == 0 {
		// Get oldProduct version
		projectClient, ctx := service.Connector()
		oldProduct, err := projectClient.Products().WithId(id).Get().Execute(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		queryArgs.Version = oldProduct.Version
	}

	projectClient, ctx := service.Connector()
	productResults, err := projectClient.Products().WithId(id).Delete().WithQueryParams(queryArgs).Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}
