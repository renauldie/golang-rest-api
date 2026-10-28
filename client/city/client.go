package cityclient

import (
	"context"
	"fmt"
	"golang-rest-api/client/common"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	CityURL     string
	CityByIDURL string
	Timeout     time.Duration
}

type Client interface {
	GetCities(ctx context.Context, id string) ([]City, error)
	GetCityByID(ctx context.Context, id string) (*City, error)
}

type client struct {
	cfg        Config
	httpClient *http.Client
}

func NewClient(cfg Config) Client {
	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}

	cfg.CityURL = strings.TrimRight(cfg.CityURL, "/")
	cfg.CityByIDURL = strings.TrimRight(cfg.CityByIDURL, "/")

	return &client{
		cfg:        cfg,
		httpClient: common.NewPooledHTTPClient(cfg.Timeout),
	}
}

func (c *client) GetCities(ctx context.Context, id string) ([]City, error) {
	targetURL := c.cfg.CityURL
	if id != "" {
		targetURL = fmt.Sprintf("%s/%s", c.cfg.CityURL, id)
	}

	var cities []City
	err := common.Exchange(ctx, c.httpClient, http.MethodGet, targetURL, nil, &cities)
	if err != nil {
		return nil, fmt.Errorf("failed to call city external api (%s): %w", targetURL, err)
	}

	return cities, nil
}

func (c *client) GetCityByID(ctx context.Context, id string) (*City, error) {
	parsedURL, err := url.Parse(c.cfg.CityByIDURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse city by id url: %w", err)
	}

	q := parsedURL.Query()
	q.Set("id", id)
	parsedURL.RawQuery = q.Encode()

	targetURL := parsedURL.String()

	var city City
	err = common.Exchange(ctx, c.httpClient, http.MethodGet, targetURL, nil, &city)
	if err != nil {
		return nil, fmt.Errorf("failed to call city by id external api (%s): %w", targetURL, err)
	}

	return &city, nil
}
