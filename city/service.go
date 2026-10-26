package city

import (
	"context"
	cityclient "golang-rest-api/client/city"
)

type Service interface {
	GetCities(ctx context.Context, id string) ([]cityclient.City, error)
}

type service struct {
	cityClient cityclient.Client
}

func NewService(cityClient cityclient.Client) Service {
	return &service{cityClient: cityClient}
}

func (s *service) GetCities(ctx context.Context, id string) ([]cityclient.City, error) {
	return s.cityClient.GetCities(ctx, id)
}
