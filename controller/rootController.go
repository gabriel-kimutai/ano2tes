package controller

import (

	"github.com/gabriel-kimutai/ano2tes/views"
	"github.com/gofiber/fiber/v2"
)

func Root(c *fiber.Ctx) error {
	return views.Render(c, views.Base(nil))
}