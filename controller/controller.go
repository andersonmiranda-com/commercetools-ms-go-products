package controller

import (
	"commercetools-ms-product/service"
	"commercetools-ms-product/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
)

func Get(c *fiber.Ctx) error {

	id := c.Params("id")

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().WithId(id).Get()

	expand := c.Query("expand")

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Find(c *fiber.Ctx) error {

	where := c.Query("where")
	expand := c.Query("expand")
	sort := c.Query("sort")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	withTotalStr := c.Query("withTotal")

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().Get()

	if where != "" {
		if whereSlice, err := utils.ConvertToSlice(where); err == nil {
			productsGetter = productsGetter.Where(whereSlice)
		}
	}

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	if sort != "" {
		if sortSlice, err := utils.ConvertToSlice(sort); err == nil {
			productsGetter = productsGetter.Sort(sortSlice)
		}
	}

	if limit, err := strconv.Atoi(limitStr); err == nil {
		productsGetter = productsGetter.Limit(limit)
	}

	if offset, err := strconv.Atoi(offsetStr); err == nil {
		productsGetter = productsGetter.Offset(offset)
	}

	if withTotal, err := strconv.ParseBool(withTotalStr); err == nil {
		productsGetter = productsGetter.WithTotal(withTotal)
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Create(c *fiber.Ctx) error {

	expand := c.Query("expand")

	productDraft := platform.ProductDraft{}

	if err := c.BodyParser(&productDraft); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ProductDraft",
		})
	}

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().Post(productDraft)

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	expand := c.Query("expand")

	productUpdate := platform.ProductUpdate{}

	if err := c.BodyParser(&productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ProductUpdate",
		})
	}

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().WithId(id).Post(productUpdate)

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Publish(c *fiber.Ctx) error {

	id := c.Params("id")
	expand := c.Query("expand")

	// Get oldProduct version
	projectClient, ctx := service.Connector()
	oldProduct, err := projectClient.Products().WithId(id).Get().Execute(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ProductUpdate",
		})
	}

	actionPayload := `{
		"version" : ` + strconv.Itoa(oldProduct.Version) + `,
		"actions" : [ {
		  "action" : "publish"
		} ]
	  }`

	productUpdate := platform.ProductUpdate{}

	if err := json.Unmarshal([]byte(actionPayload), &productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse actionPayload",
		})
	}

	productsGetter := projectClient.Products().WithId(id).Post(productUpdate)

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Unpublish(c *fiber.Ctx) error {

	id := c.Params("id")
	expand := c.Query("expand")

	// Get oldProduct version
	projectClient, ctx := service.Connector()
	oldProduct, err := projectClient.Products().WithId(id).Get().Execute(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ProductUpdate",
		})
	}

	actionPayload := `{
		"version" : ` + strconv.Itoa(oldProduct.Version) + `,
		"actions" : [ {
		  "action" : "unpublish"
		} ]
	  }`

	productUpdate := platform.ProductUpdate{}

	if err := json.Unmarshal([]byte(actionPayload), &productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse actionPayload",
		})
	}

	productsGetter := projectClient.Products().WithId(id).Post(productUpdate)

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}

func Remove(c *fiber.Ctx) error {

	id := c.Params("id")
	versionStr := c.Query("version")
	expand := c.Query("expand")

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().WithId(id).Delete()

	if version, err := strconv.Atoi(versionStr); err == nil {
		productsGetter = productsGetter.Version(version)
	}

	if expand != "" {
		if expandSlice, err := utils.ConvertToSlice(expand); err == nil {
			productsGetter = productsGetter.Expand(expandSlice)
		}
	}

	productResults, err := productsGetter.Execute(ctx)
	return utils.Response(productResults, http.StatusOK, err, c)
}
