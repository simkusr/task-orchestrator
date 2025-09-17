package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const (
	defaultEnvFileName = ".env"
)

type Config struct {
	Port string `env:"PORT"`
}

func (c *Config) NewConfig() error {
	loadEnvironmentVariables()

	err := env.Parse(c)

	return err
}

func loadEnvironmentVariables() {
	envFileName, ok := os.LookupEnv("ENV_FILE_NAME")
	if !ok {
		envFileName = defaultEnvFileName
	}

	godotenv.Load(envFileName)
}
