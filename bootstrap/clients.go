package bootstrap

import (
	cityclient "golang-rest-api/client/city"
	"golang-rest-api/config"
)

type Clients struct {
	City cityclient.Client
}

func InitClients(cfg *config.ClientConfig) *Clients {
	return &Clients{
		City: cityclient.NewClient(cfg.City),
	}
}
