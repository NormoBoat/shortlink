package application

import (
	"fmt"

	"github.com/NormoBoat/shortlink/internal/config"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

var (
	errDontLoadHttp    = "Не удалось настроить http"
	errHttpReturnFatal = "Http упал с ошибкой"
)

type Application struct {
	web    *fiber.App
	config *config.Config
}

func New() *Application {

	return &Application{
		config: config.New(),
		web:    setUpHTTP(),
	}
}

func (a *Application) Start() {
	param := fmt.Sprint(a.config.Web.Host, ":", a.config.Web.Port)
	if err := a.web.Listen(param); err != nil {
		log.Fatal().Msgf("%s: %s", errHttpReturnFatal, err)
	}
}

func setUpHTTP() *fiber.App {
	app := fiber.New()
	if app == nil {
		log.Fatal().Msg(errDontLoadHttp)
	}
	routeRetistartion(app)
	return app
}
