package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	Mode string `envconfig:"MODE" default:"development"`
	Port string `envconfig:"PORT" default:":3000"`
	DBConfig
}

type DBConfig struct {
	DSN string `envconig:"DSN"`
}

func (c *Config) IsProduction() bool {
	return c.Mode == "production"
}

func Get() *Config {
	once.Do(func() {
		// load env vars from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		err = envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
	})

	return &config
}
