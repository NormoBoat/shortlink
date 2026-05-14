package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	dontLoadENV = errors.New("Не удалось загрузить переменные из окружения")
)

type WebConfig struct {
	Port string
	Host string
}

type Config struct {
	Web WebConfig
}

func New() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf(dontLoadENV.Error())
	}

	return &Config{
		Web: WebConfig{
			Port: loadKeyOrDefault("HTTP_PORT", "8080"),
			Host: loadKeyOrDefault("HTTP_HOST", "localhost"),
		},
	}
}

func loadKeyOrDefault(key string, def string) string {
	result := os.Getenv(key)
	if result == "" {
		log.Info().Msgf("Не удалось загрузить переменную окружения \"%s\", будет установленно значение поумолчанию \"%s\"", key, def)
		result = def
	}
	return result
}
