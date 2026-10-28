package city

import (
	"context"
	"errors"
	"strings"

	cityclient "golang-rest-api/client/city"
)

type Service interface {
	GetCities(ctx context.Context, id string) ([]cityclient.City, error)
	GetCityByID(ctx context.Context, id string) (*cityclient.City, error)
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

func (s *service) GetCityByID(ctx context.Context, id string) (*cityclient.City, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("city id is required")
	}
	return s.cityClient.GetCityByID(ctx, id)
}
