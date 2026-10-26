package city_test

import (
	"context"
	"errors"
	"testing"

	"golang-rest-api/city"
	cityclient "golang-rest-api/client/city"
)

type mockCityClient struct {
	getCitiesFn func(ctx context.Context, id string) ([]cityclient.City, error)
}

func (m *mockCityClient) GetCities(ctx context.Context, id string) ([]cityclient.City, error) {
	if m.getCitiesFn != nil {
		return m.getCitiesFn(ctx, id)
	}
	return nil, nil
}

func TestCityService_GetCities_Success(t *testing.T) {
	mockClient := &mockCityClient{
		getCitiesFn: func(ctx context.Context, id string) ([]cityclient.City, error) {
			return []cityclient.City{
				{ID: "11", Name: "Jakarta Selatan"},
			}, nil
		},
	}

	svc := city.NewService(mockClient)
	cities, err := svc.GetCities(context.Background(), "12")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cities) != 1 {
		t.Fatalf("expected 1 city, got %d", len(cities))
	}
	if cities[0].Name != "Jakarta Selatan" {
		t.Errorf("expected Jakarta Selatan, got %s", cities[0].Name)
	}
}

func TestCityService_GetCities_Error(t *testing.T) {
	mockClient := &mockCityClient{
		getCitiesFn: func(ctx context.Context, id string) ([]cityclient.City, error) {
			return nil, errors.New("connection failed")
		},
	}

	svc := city.NewService(mockClient)
	_, err := svc.GetCities(context.Background(), "12")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
