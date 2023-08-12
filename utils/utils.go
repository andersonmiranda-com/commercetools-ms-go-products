package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	var resultSlice []string
	errSlice := json.Unmarshal([]byte(input), &resultSlice)
	if errSlice == nil {
		return resultSlice, nil
	}

	var resultString string
	errString := json.Unmarshal([]byte(input), &resultString)
	if errString == nil {
		return []string{resultString}, nil
	}

	return nil, fmt.Errorf("input is neither a valid string nor a valid array of strings")
}

func Response(data interface{}, httpStatus int, err error, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		if data != nil {
			return c.Status(httpStatus).JSON(data)
		} else {
			c.Status(httpStatus)
			return nil
		}
	}
}
