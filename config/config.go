package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		Bot  `yaml:"bot"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	}

	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL" env-default:"debug"`
	}

	PG struct {
		Host     string `yaml:"host"     env:"PG_HOST"     env-default:"localhost"`
		Port     string `yaml:"port"     env:"PG_PORT"     env-default:"5432"`
		Username string `yaml:"username" env:"PG_DB"       env-default:"postgres"`
		Password string `                env:"PG_PASSWORD" env-default:"qwerty123"`
		DBName   string `yaml:"dbname"   env:"PG_NAME"     env-default:"postgres"`
		SSLMode  string `yaml:"sslmode"  env:"PG_SSL"      env-default:"disable"`
	}

	Bot struct {
		Token string `yaml:"token" env:"TG_TOKEN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
