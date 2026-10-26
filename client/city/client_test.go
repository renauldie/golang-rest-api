package cityclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	cityclient "golang-rest-api/client/city"
)

func TestCityClient_GetCities_DynamicURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/xco/irh/page/migration/city/12" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"11","name":"Jakarta Selatan"}]`))
	}))
	defer server.Close()

	dynamicCityURL := server.URL + "/api/xco/irh/page/migration/city"
	client := cityclient.NewClient(dynamicCityURL, 2*time.Second)

	cities, err := client.GetCities(context.Background(), "12")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cities) != 1 {
		t.Fatalf("expected 1 city, got %d", len(cities))
	}
	if cities[0].ID != "11" || cities[0].Name != "Jakarta Selatan" {
		t.Errorf("unexpected city data: %+v", cities[0])
	}
}

func TestCityClient_GetCities_CatchesPanic(t *testing.T) {
	// Menguji bahwa blok defer-recover (Try-Catch) berhasil menangkap panic
	client := cityclient.NewClient("http://localhost:7024/api/xco/irh/page/migration/city", 1*time.Second)

	// Passing nil context akan memicu error pada NewRequestWithContext
	_, err := client.GetCities(nil, "12")
	if err == nil {
		t.Fatalf("expected error when context is nil, got nil")
	}

	if !strings.Contains(err.Error(), "failed to create request") && !strings.Contains(err.Error(), "caught panic") {
		t.Errorf("unexpected error message: %v", err)
	}
}
