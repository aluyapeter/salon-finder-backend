package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	rr := httptest.NewRecorder()

	Health(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	expected := "Salon Finder API v0.1 - OK"
	if string(body) != expected {
		t.Errorf("expected body %q, got %q", expected, string(body))
	}
}