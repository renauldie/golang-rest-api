package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Port string
	Env  string
}

func LoadAppConfig() *AppConfig {
	port := os.Getenv("APP_PORT")
	env := os.Getenv("ENV")
	return &AppConfig{
		Port: port,
		Env:  env,
	}
}

func (c *AppConfig) Address() string {
	return fmt.Sprintf(":%s", c.Port)
}
