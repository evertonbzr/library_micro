package config

import (
	"log"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

var (
	PORT         string
	ENV          string
	NAME         string
	REDIS_URL    string
	DATABASE_URL string
	NATS_URI     string
	JWT_SECRET   string
)

func Load(env string) {
	slog.Info("Loading config...", "env", env)

	if env == "" {
		viper.Set("ENV", "development")
		viper.Set("PORT", "3000")
		viper.Set("NAME", "module_user")
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error loading .env file", "error", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	PORT = viper.GetString("PORT")
	ENV = viper.GetString("ENV")
	NAME = viper.GetString("NAME")
	REDIS_URL = viper.GetString("REDIS_URL")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	JWT_SECRET = viper.GetString("JWT_SECRET")

	NATS_URI = nats.DefaultURL

	if viper.GetString("NATS_URI") != "" {
		NATS_URI = viper.GetString("NATS_URI")
	}
}

func IsDevelopment() bool {
	return ENV == "development"
}

func IsProduction() bool {
	return ENV == "production"
}

func IsTest() bool {
	return ENV == "test"
}
