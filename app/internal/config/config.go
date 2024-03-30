package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`

	GRPC struct {
		IP   string `env:"GRPC_IP" env-default:"0.0.0.0"`
		Port string `env:"GRPC_PORT" env-default:"25819"`
	}

	PostgreSQL struct {
		PostgreUsername string `env:"PSQL_USERNAME" env-required:"true"`
		Password        string `env:"PSQL_PASSWORD" env-required:"true"`
		Database        string `env:"PSQL_DATABASE" env-required:"true"`
		Host            string `env:"PSQL_HOST" env-default:"localhost"`
		Port            string `env:"PSQL_PORT" env-default:"5432"`
	}
	Jwt struct {
		SecretKey     string `env:"JWT_SECRET_KEY" env-required:"true"`
		TokenDuration string `env:"JWT_TOKEN_DURATION" env-required:"true"`
	}
}

// Singleton: Config should only ever be created once
var instance *Config

// Once is an object that will perform exactly one action.
var once sync.Once

// GetConfig returns pointer to Config
func GetConfig() *Config {
	// Calls the function if and only if Do is being called for the first time for this instance of Once
	once.Do(func() {
		log.Print("collecting config...")

		// Config initialization
		instance = &Config{}

		// Read environment variables into the instance of the Config
		if err := cleanenv.ReadEnv(instance); err != nil {
			// If something is wrong
			helpText := "Environment variables error:"
			// Returns a description of environment variables with a custom header - helpText
			help, err := cleanenv.GetDescription(instance, &helpText)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(help)
			log.Printf("%+v\n", instance)

			log.Fatal(err)
		}
	})
	return instance
}
