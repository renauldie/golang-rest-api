package config

import "os"

type ClientConfig struct {
	CityApiURL string
}

func LoadClientConfig() *ClientConfig {
	return &ClientConfig{
		CityApiURL: os.Getenv("CITY_API_URL"),
	}
}
