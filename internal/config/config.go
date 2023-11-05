package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	DbConfig   `yaml:"db_config"`
}

type HttpServer struct {
	Address string `yaml:"address"`
	Host    string `yaml:"host"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
}

func MustLoad() *Config {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %w", err)
	}
	return &cfg
}
