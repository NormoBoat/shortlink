package views

import "github.com/gofiber/fiber/v3"

func Root(ctx fiber.Ctx) error {
	return ctx.SendString("keep")
}
