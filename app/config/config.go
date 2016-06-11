package config

import (
	"dskmoney-golang/app/utils"
	"github.com/go-ini/ini"
	"fmt"
)

const (
	ConfigPath = "etc/"
	BaseConfig = "base.ini"

	ProductionEnv = "production"
	DevEnv = "dev"
	DefaultEnv = DevEnv
)

var (
	AllowedEnvs = []string{ProductionEnv, DevEnv}
)

type App struct {
	Env string `ini:"env"`
	Port string `ini:"port"`
}

type DB struct {
	Host      string `ini:"db_host"`
	Name      string `ini:"db_name"`
	User      string `ini:"db_user"`
	Password  string `ini:"db_password"`
}

type Config struct {
	App `ini:"DEFAULT"`
	DB `ini:"DEFAULT"`
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Init() error {
	cfg := ini.Empty()
	// it makes reading 50-70% faster, but we should not write to config file in that case
	cfg.BlockMode = false

	err := cfg.Append(ConfigPath + BaseConfig)
	if err != nil {
		return err
	}

	// settings from environment variables are more important
	envCfg := ReadEnvConfig()
	if utils.StrInSlice(envCfg.Env, AllowedEnvs) {
		this.Env = envCfg.Env

		err := cfg.Append(ConfigPath + envCfg.Env + ".ini")
		if err != nil {
			fmt.Println(envCfg.Env + " config does not exist. Only base config was used.")
		}
	} else {
		this.Env = DefaultEnv
	}

	err = cfg.MapTo(this)
	if err != nil {
		return err
	}

	// Port from environment variable is more important
	if envCfg.Port != "" {
		this.Port = envCfg.Port
	}

	fmt.Printf("Config: %#v\n", this)
	return nil
}
