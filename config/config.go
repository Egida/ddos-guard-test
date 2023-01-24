package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		Bot  `yaml:"bot"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	PG struct {
		Host     string `env-required:"true" yaml:"host"    env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port"    env:"PG_PORT"`
		Username string `env-required:"true"                env:"PG_USERNAME"`
		Password string `env-required:"true"                env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"dbname"  env:"PG_DBNAME"`
		SSLMode  string `env-required:"true" yaml:"sslmode" env:"PG_SSLMODE"`
	}

	Bot struct {
		Token string `env-required:"true" env:"TG_TOKEN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("config - NewConfig - cleanenv.ReadConfig: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("config - NewConfig - cleanenv.ReadEnv: %w", err)
	}

	return cfg, nil
}
