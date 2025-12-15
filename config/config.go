package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Host         string        `yaml:"Host"`
	Port         int           `yaml:"Port"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type PostgresConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBname   string `yaml:"DBname"`
}

type RedisConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

type LoggingConfig struct {
	Level string `yaml:"Level"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	env := "dev"
	if e := os.Getenv("APP_ENV"); e != "" {
		env = e
	}

	configPath := fmt.Sprintf("config/config.%s.yaml", env)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, errors.New("config file not found")
	}
	return &cfg, nil
}
