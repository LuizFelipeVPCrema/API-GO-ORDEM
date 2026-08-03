package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	// TODO: Implementar autenticação e logging
	// Auth     AuthConfig
	// Log      LogConfig
}

type DatabaseConfig struct {
	Path string
}

type ServerConfig struct {
	Port string
}

func Carregar() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf(
			"erro ao carregar o arquivo .env: %w",
			err,
		)
	}

	cfg := &Config{
		Database: DatabaseConfig{
			Path: os.Getenv("DB_PATH"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}

	if err := cfg.validar(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validar() error {
	if c.Database.Path == "" {
		return fmt.Errorf(
			"variavel de ambiente DB_PATH e obrigatoria",
		)
	}

	if c.Server.Port == "" {
		return fmt.Errorf(
			"variavel de ambiente SERVER_PORT e obrigatoria",
		)
	}

	return nil
}
