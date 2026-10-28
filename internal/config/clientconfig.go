package config

import (
	"os"
	"time"

	cityclient "golang-rest-api/client/city"
)

type ClientConfig struct {
	City cityclient.Config
}

func LoadClientConfig() *ClientConfig {
	return &ClientConfig{
		City: cityclient.Config{
			CityURL:     os.Getenv("CITY_API_URL"),
			CityByIDURL: os.Getenv("CITY_API_BY_ID_URL"),
			Timeout:     10 * time.Second,
		},
	}
}
