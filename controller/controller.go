package controller

import (
	"commercetools-ms-product/service"
	"commercetools-ms-product/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	productsGetter := projectClient.Products().WithId("").

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
