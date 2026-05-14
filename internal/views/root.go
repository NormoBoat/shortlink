package views

import (
	"github.com/NormoBoat/shortlink/internal/views/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func Root(ctx fiber.Ctx) error {

	component := templates.Header("Главная страница")
	return Render(ctx, component)
}

// TODO перенести в другое место
func Render(c fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}
