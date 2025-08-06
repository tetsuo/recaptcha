package recaptcha_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/tetsuo/recaptcha"
)

type mockTransport struct {
	RoundTripFunc func(req *http.Request) *http.Response
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req), nil
}

func TestVerify_Success(t *testing.T) {
	responseBody := `{
		"success": true,
		"challenge_ts": "2025-08-06T12:00:00Z",
		"hostname": "example.com",
		"error-codes": []
	}`

	client := &http.Client{
		Transport: &mockTransport{
			RoundTripFunc: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
					Header:     make(http.Header),
				}
			},
		},
	}

	ctx := context.Background()
	challengeResponse := "dummy-challenge"
	secretKey := "dummy-secret"
	remoteIP := "127.0.0.1"

	resp, err := recaptcha.Verify(ctx, client, challengeResponse, secretKey, remoteIP)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("expected success true, got false")
	}
	if resp.Hostname != "example.com" {
		t.Errorf("expected hostname 'example.com', got '%s'", resp.Hostname)
	}
}
