package config

import (
	"errors"
	"os"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

type ServerConfig struct {
	Port     string
	Debug    bool
	Timezone string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func ProvideConfig() (*Config, error) {
	var cfg Config

	if err := loadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv(config *Config) error {
	config.App.Name = os.Getenv("APP_NAME")
	if config.App.Name == "" {
		return errors.New("APP_NAME must be set")
	}

	config.App.Version = os.Getenv("APP_VERSION")
	if config.App.Version == "" {
		return errors.New("APP_VERSION must be set")
	}

	config.App.Environment = os.Getenv("APP_ENV")
	if config.App.Environment == "" {
		return errors.New("APP_ENV must be set")
	}

	config.Server.Port = os.Getenv("SERVER_PORT")
	if config.Server.Port == "" {
		return errors.New("SERVER_PORT must be set")
	}

	config.Server.Debug = os.Getenv("SERVER_DEBUG") == "true"

	config.Server.Timezone = os.Getenv("SERVER_TIMEZONE")
	if config.Server.Timezone == "" {
		return errors.New("SERVER_TIMEZONE must be set")
	}

	config.Database.Host = os.Getenv("DB_HOST")
	if config.Database.Host == "" {
		return errors.New("DB_HOST must be set")
	}

	config.Database.Port = os.Getenv("DB_PORT")
	if config.Database.Port == "" {
		return errors.New("DB_PORT must be set")
	}

	config.Database.Username = os.Getenv("DB_USERNAME")
	if config.Database.Username == "" {
		return errors.New("DB_USERNAME must be set")
	}

	config.Database.Password = os.Getenv("DB_PASSWORD")
	if config.Database.Password == "" {
		return errors.New("DB_PASSWORD must be set")
	}

	config.Database.Name = os.Getenv("DB_NAME")
	if config.Database.Name == "" {
		return errors.New("DB_NAME must be set")
	}

	return nil
}
