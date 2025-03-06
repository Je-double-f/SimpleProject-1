package config

import (
	"fmt"
)

// Config содержит все конфигурационные параметры приложения
type Config struct {
	DatabaseURL string // URL для подключения к PostgreSQL
	ServerPort  int    // Порт, на котором будет работать сервер

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig загружает конфигурацию
func LoadConfig() *Config {
	cfg := &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "super_user",
		DBPassword: "1234",
		DBName:     "postgres",
		ServerPort: 8080,
	}

	// Формирование строки подключения
	cfg.DatabaseURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	return cfg
}
