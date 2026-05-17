package application

import (
	"github.com/NormoBoat/shortlink/internal/views"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func Render(c fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}

func routeRetistartion(app *fiber.App) {
	component := views.Links()
	app.Get("/", func(ctx fiber.Ctx) error {
		s := session.FromContext(ctx)
		res := s.Session.Get("foo")
		return Render(ctx, component)
	})
}
