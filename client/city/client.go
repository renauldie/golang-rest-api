package cityclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client interface {
	GetCities(ctx context.Context, id string) ([]City, error)
}

type client struct {
	cityURL    string
	httpClient *http.Client
}

func NewClient(cityURL string, timeout time.Duration) Client {
	if cityURL == "" {
		cityURL = os.Getenv("CITY_API_URL")
	}

	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	return &client{
		cityURL: strings.TrimRight(cityURL, "/"),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *client) GetCities(ctx context.Context, id string) (cities []City, err error) {
	targetURL := fmt.Sprintf("%s/%s", c.cityURL, id)

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call city external api (%s): %w", targetURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external api returned status code: %d from %s", resp.StatusCode, targetURL)
	}

	if decodeErr := json.NewDecoder(resp.Body).Decode(&cities); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode response from %s: %w", targetURL, decodeErr)
	}

	return cities, nil
}
