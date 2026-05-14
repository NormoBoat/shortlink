package application

import (
	"github.com/NormoBoat/shortlink/internal/views"

	"github.com/gofiber/fiber/v3"
)

func routeRetistartion(app *fiber.App) {
	app.Get("/", views.Root)
}
