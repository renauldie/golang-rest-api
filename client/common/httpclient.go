package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang-rest-api/pkg/correlation"
	"io"
	"net"
	"net/http"
	"time"
)

type RequestOption func(*http.Request)

func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

func WithHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
}

func WithContentType(contentType string) RequestOption {
	return WithHeader("Content-Type", contentType)
}

func WithAccept(accept string) RequestOption {
	return WithHeader("Accept", accept)
}

func NewPooledHTTPClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   25,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

func Exchange[Resp any](
	ctx context.Context,
	httpClient *http.Client,
	method string,
	targetURL string,
	requestPayload any,
	responseDest *Resp,
	opts ...RequestOption,
) error {
	var bodyReader io.Reader

	if requestPayload != nil {
		switch p := requestPayload.(type) {
		case []byte:
			bodyReader = bytes.NewBuffer(p)
		case string:
			bodyReader = bytes.NewBufferString(p)
		case io.Reader:
			bodyReader = p
		default:
			bodyBytes, err := json.Marshal(requestPayload)
			if err != nil {
				return fmt.Errorf("failed to marshal request payload: %w", err)
			}
			bodyReader = bytes.NewBuffer(bodyBytes)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, targetURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if requestPayload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	if corrID := correlation.FromContext(ctx); corrID != "" {
		req.Header.Set(correlation.HeaderCorrelationID, corrID)
	}

	for _, opt := range opts {
		if opt != nil {
			opt(req)
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute external request to %s: %w", targetURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("external api returned unexpected status code: %d from %s", resp.StatusCode, targetURL)
	}

	if responseDest != nil {
		if decodeErr := json.NewDecoder(resp.Body).Decode(responseDest); decodeErr != nil {
			return fmt.Errorf("failed to decode response from %s: %w", targetURL, decodeErr)
		}
	}

	return nil
}
