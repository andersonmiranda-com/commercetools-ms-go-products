package controller

import (
	"commercetools-ms-product/service"
	"commercetools-ms-product/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {

	where := utils.GetWhereClause(c)
	fmt.Println(where)

	projectClient, ctx := service.Connector()
	productsGetter := projectClient.Products().Get().Where(where)

	expand := c.Query("expand")
	sort := c.Query("sort")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	withTotalStr := c.Query("withTotal")

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

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return utils.Response(productResults, http.StatusOK, err, c)
}
