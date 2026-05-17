package application

import (
	"context"
	"fmt"
	"time"

	"github.com/NormoBoat/shortlink/internal/config"
	"github.com/NormoBoat/shortlink/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	recover2 "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var (
	errDontLoadHttp    = "Не удалось настроить http"
	errHttpReturnFatal = "Http упал с ошибкой"
)

type Application struct {
	web    *fiber.App
	config *config.Config
	store  *storage.Storage
}

func New() *Application {

	cfg := config.New()
	web := setUpHTTP()
	store := storage.New()

	dbPool, err := pgxpool.New(context.Background(), )

	sess := session.New(session.Config{
		Storage: store.DB.,
		CookieSecure:    true,
		CookieHTTPOnly:  false, // TODO add to config
		CookieSameSite:  "Lax",
		IdleTimeout:     30 * time.Minute, // TODO too
		AbsoluteTimeout: 24 * time.Hour,   // TODO too
		Extractor:       extractors.FromCookie("__Host-session_id"),
	})

	web.Use(sess)
	web.Use(recover2.New())
	routeRetistartion(web)

	return &Application{
		web,
		cfg,
		store,
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
	return app
}
