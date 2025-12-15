package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Serverv  ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
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

func MustLoad() *Config {
	var cfg Config

	env := "dev"
	if e := os.Getenv("APP_ENV"); e != "" {
		env = e
	}

	configPath := fmt.Sprintf("config/config.%s.yaml", env)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		panic(fmt.Sprintf("cannot read config file: %w", err))
	}
	return &cfg
}
