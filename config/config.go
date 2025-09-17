package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const (
	envFileName        = "ENV_FILE_NAME"
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
	envFileName, ok := os.LookupEnv(envFileName)
	if !ok {
		envFileName = defaultEnvFileName
	}

	godotenv.Load(envFileName)
}
