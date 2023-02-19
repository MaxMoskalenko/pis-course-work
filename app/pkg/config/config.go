package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
}

type Database struct {
	Name     string `env:"DATABASE_NAME" envDefault:"postgres"`
	Schema   string `env:"DATABASE_SCHEMA" envDefault:"postgres"`
	Username string `env:"DATABASE_USERNAME" envDefault:"postgres"`
	Password string `env:"DATABASE_PASSWORD"`
	Host     string `env:"DATABASE_HOST" envDefault:"0.0.0.0"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
}

func Load(ConfigFilePath string) (*Config, error) {
	godotenv.Load(ConfigFilePath)
	cnf := &Config{
		Database: Database{},
	}

	fmt.Println("Loading Config")
	if err := env.Parse(cnf); err != nil {
		return nil, err
	}

	return cnf, nil
}
