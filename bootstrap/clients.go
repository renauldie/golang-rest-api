package bootstrap

import (
	"time"

	cityclient "golang-rest-api/client/city"
	"golang-rest-api/config"
)

type Clients struct {
	City cityclient.Client
}

func InitClients(cfg *config.ClientConfig) *Clients {
	return &Clients{
		City: cityclient.NewClient(cfg.CityApiURL, 10*time.Second),
	}
}
