package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetPostID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.IncrementViewCount(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	post, err := h.service.GetPostByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get post",
		})
	}

	accept := c.Get("Accept")

	if accept == "text/html" {
		if accept == "text/html" {
			return c.Render("postsEdit", fiber.Map{
				"Post": post,
			})
		}
	}

	return c.JSON(post)
}
