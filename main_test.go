package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if resp["status"] != "ok" {
		t.Errorf("status field = %q, want %q", resp["status"], "ok")
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{"default", "/hello", "hello, world"},
		{"with name", "/hello?name=Deepi", "hello, Deepi"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			var resp map[string]string
			json.NewDecoder(w.Body).Decode(&resp)
			if resp["message"] != tc.want {
				t.Errorf("message = %q, want %q", resp["message"], tc.want)
			}
		})
	}
}
