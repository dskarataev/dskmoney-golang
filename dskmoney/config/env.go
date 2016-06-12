package config

import (
	"os"
)

type EnvConfig struct {
	Port       string
	Env        string
}

func ReadEnvConfig() EnvConfig {
	return EnvConfig{
		Port: os.Getenv("PORT"),
		Env: os.Getenv("ENV"),
	}
}
