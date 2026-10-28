package common_test

import (
	"context"
	"encoding/json"
	"golang-rest-api/client/common"
	"golang-rest-api/pkg/correlation"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type sampleRequest struct {
	Message string `json:"message"`
}

type sampleResponse struct {
	Status string `json:"status"`
}

func TestExchange_GET_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept application/json, got %s", r.Header.Get("Accept"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := common.NewPooledHTTPClient(2 * time.Second)

	var resp sampleResponse
	err := common.Exchange(context.Background(), client, http.MethodGet, server.URL, nil, &resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", resp.Status)
	}
}

func TestExchange_POST_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		bodyBytes, _ := io.ReadAll(r.Body)
		var req sampleRequest
		_ = json.Unmarshal(bodyBytes, &req)

		if req.Message != "hello" {
			t.Errorf("expected message 'hello', got '%s'", req.Message)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"created"}`))
	}))
	defer server.Close()

	client := common.NewPooledHTTPClient(2 * time.Second)

	var resp sampleResponse
	err := common.Exchange(context.Background(), client, http.MethodPost, server.URL, sampleRequest{Message: "hello"}, &resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status != "created" {
		t.Errorf("expected status 'created', got '%s'", resp.Status)
	}
}

func TestExchange_DynamicHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("expected Content-Type application/x-www-form-urlencoded, got %s", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Accept") != "application/xml" {
			t.Errorf("expected Accept application/xml, got %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("Authorization") != "Bearer secret-token" {
			t.Errorf("expected Authorization Bearer secret-token, got %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("X-Custom-Header") != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", r.Header.Get("X-Custom-Header"))
		}

		bodyBytes, _ := io.ReadAll(r.Body)
		if string(bodyBytes) != "grant_type=client_credentials" {
			t.Errorf("unexpected body: %s", string(bodyBytes))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"authorized"}`))
	}))
	defer server.Close()

	client := common.NewPooledHTTPClient(2 * time.Second)

	var resp sampleResponse
	err := common.Exchange(
		context.Background(),
		client,
		http.MethodPost,
		server.URL,
		"grant_type=client_credentials",
		&resp,
		common.WithContentType("application/x-www-form-urlencoded"),
		common.WithAccept("application/xml"),
		common.WithHeader("Authorization", "Bearer secret-token"),
		common.WithHeaders(map[string]string{
			"X-Custom-Header": "custom-value",
		}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status != "authorized" {
		t.Errorf("expected status 'authorized', got '%s'", resp.Status)
	}
}

func TestExchange_CorrelationID_Forwarding(t *testing.T) {
	testCorrID := "corr-id-9988"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(correlation.HeaderCorrelationID) != testCorrID {
			t.Errorf("expected correlation ID %s, got %s", testCorrID, r.Header.Get(correlation.HeaderCorrelationID))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := common.NewPooledHTTPClient(2 * time.Second)
	ctx := correlation.WithCorrelationID(context.Background(), testCorrID)

	var resp map[string]any
	err := common.Exchange(ctx, client, http.MethodGet, server.URL, nil, &resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
