package handlers

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetPosts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1
	defaultPublished := true

	limit, err := strconv.Atoi(c.Query("limit", strconv.Itoa(defaultLimit)))
	if err != nil || limit <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	page, err := strconv.Atoi(c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil || page <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page parameter",
		})
	}

	offset := (page - 1) * limit
	searchTitle := c.Query("title", "")

	publishedStr := c.Query("published")
	var published bool
	if publishedStr == "" {
		published = defaultPublished
	} else {
		published = strings.ToLower(publishedStr) == "true"
	}

	articles, err := h.service.GetPostAll(limit, offset, searchTitle, published)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalArticles, err := h.service.GetTotalPostCount(searchTitle, published)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalPages := int(math.Ceil(float64(totalArticles) / float64(limit)))

	responseData := fiber.Map{
		"posts":      articles,
		"count":      totalArticles,
		"limit":      limit,
		"page":       page,
		"total_page": totalPages,
	}

	accept := c.Get("Accept")
	if accept == "text/html" {
		var nextPage, prevPage int
		if page < totalPages {
			nextPage = page + 1
		} else {
			nextPage = 0
		}
		if page > 1 {
			prevPage = page - 1
		} else {
			prevPage = 0
		}

		return c.Render("postsList", fiber.Map{
			"Posts":      articles,
			"Count":      totalArticles,
			"Limit":      limit,
			"Page":       page,
			"NextPage":   nextPage,
			"PrevPage":   prevPage,
			"TotalPages": totalPages,
			"Published":  published,
		})
	}

	return c.JSON(responseData)
}
