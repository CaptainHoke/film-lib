package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env:"NK_ENV" env-default:"local" env-required:"true"`
	DbConfig   `yaml:"db_config" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:3239"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"15s"`
}

// DbConfig TODO: fancy stuff
type DbConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SSLMode      string
	PoolMaxConns int
}

func MustLoadCfg(cfgPath string) Config {
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Config not found: %s", cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Config read error: %s", err)
	}

	return cfg
}
