package config

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoadConfig struct {
	App struct {
		Debug      bool   `env:"APP_DEBUG"`
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET_KEY"`
	}

	Database struct {
		Host     string `env:"DB_HOST"`
		Name     string `env:"DB_NAME"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     string `env:"DB_PORT"`
	}
}

type Routing struct {
	Router   *fiber.App
	Database *gorm.DB
}
