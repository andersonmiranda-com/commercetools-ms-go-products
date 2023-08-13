package utils

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
)

func GetWhereClause(c *fiber.Ctx) []string {

	name := c.Query("name")
	slug := c.Query("slug")
	key := c.Query("key")
	locale := c.Query("locale")
	if locale == "" {
		locale = "en"
	}

	where := generateWhereClause(name, slug, key, locale)

	// Collecting other query parameters as strings
	whereClause := []string{}
	excludedKeys := map[string]bool{
		"name":      true,
		"slug":      true,
		"key":       true,
		"locale":    true,
		"expand":    true,
		"sort":      true,
		"limit":     true,
		"offset":    true,
		"withTotal": true,
	}

	c.Request().URI().QueryArgs().VisitAll(func(k, v []byte) {
		key := string(k)
		if !excludedKeys[key] {
			whereClause = append(whereClause, key+"="+string(v))
		}
	})

	if len(where) > 0 {
		whereClause = append(whereClause, where...)
	}

	return whereClause
}

func generateWhereClause(name, slug, key, locale string) []string {
	var where []string

	if name != "" {
		where = append(where, "masterData(current(name("+locale+"=\""+name+"\")))")
	}

	if key != "" {
		where = append(where, "key=\""+key+"\"")
	}

	if slug != "" {
		where = append(where, "masterData(current(slug("+locale+"=\""+slug+"\")))")
	}

	return where
}

func ConvertToSlice(input string) ([]string, error) {
	if strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]") {
		trimmed := strings.Trim(input, "[]")
		resultSlice := strings.Split(trimmed, ",")
		for i, val := range resultSlice {
			resultSlice[i] = strings.TrimSpace(val) // remove potential whitespace
		}
		return resultSlice, nil
	}

	return []string{input}, nil
}

func Response(data interface{}, httpStatus int, err error, c *fiber.Ctx) error {

	if err != nil {
		if err.Error() == "resource not found" {
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
