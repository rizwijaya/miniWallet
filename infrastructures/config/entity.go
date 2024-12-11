package config

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoadConfig struct {
	App struct {
		Mode       string `env:"APP_MODE"`
		Debug      bool   `env:"APP_DEBUG"`
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET_KEY"`
	}

	Database struct {
		Driver   string `env:"DB_DRIVER"`
		Host     string `env:"DB_HOST"`
		Name     string `env:"DB_NAME"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     string `env:"DB_PORT"`
	}

	Memcache struct {
		Host string `env:"MEMCACHE_HOST"`
		Port string `env:"MEMCACHE_PORT"`
	}
}

type Routing struct {
	Router   *fiber.App
	Database *gorm.DB
	Memcache *memcache.Client
}
