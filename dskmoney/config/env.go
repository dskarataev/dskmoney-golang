package config

import (
	"os"
)

type EnvConfig struct {
	Deploy
	DB
}

func ReadEnvConfig() EnvConfig {
	return EnvConfig{
		Deploy: Deploy{
			Env:  os.Getenv("ENV"),
			Port: os.Getenv("PORT"),
		},
		DB: DB{
			Addr:   os.Getenv("DB_ADDR"),
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWD"),
			Name:   os.Getenv("DB_NAME"),
		},
	}
}
