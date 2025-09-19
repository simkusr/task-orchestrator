package config

import (
	"errors"
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
	err := loadEnvironmentVariables()
	if err != nil {
		return errors.Join(err, errors.New("load env vars"))
	}

	return env.Parse(c)

}

func loadEnvironmentVariables() error {
	envFileName, ok := os.LookupEnv(envFileName)
	if !ok {
		envFileName = defaultEnvFileName
	}

	return godotenv.Load(envFileName)
}
