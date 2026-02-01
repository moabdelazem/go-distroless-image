package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	handleRoot(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("status code = %d; want %d", recorder.Code, http.StatusOK)
	}

	var response Response
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedMessage := "Hello from a Configurable Distroless API!"
	if response.Message != expectedMessage {
		t.Errorf("Message = %q; want %q", response.Message, expectedMessage)
	}
}

func TestHandleHealth(t *testing.T) {
	tests := []struct {
		name            string
		appName         string
		expectedStatus  string
		expectedService string
	}{
		{
			name:            "returns healthy status with default app name",
			appName:         "Distroless API",
			expectedStatus:  "healthy",
			expectedService: "Distroless API",
		},
		{
			name:            "returns healthy status with custom app name",
			appName:         "Custom Service",
			expectedStatus:  "healthy",
			expectedService: "Custom Service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				SrvPort: ":8080",
				AppName: tt.appName,
			}

			handler := handleHealth(cfg)
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			recorder := httptest.NewRecorder()

			handler(recorder, req)

			if recorder.Code != http.StatusOK {
				t.Errorf("status code = %d; want %d", recorder.Code, http.StatusOK)
			}

			var response HealthResponse
			if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if response.Status != tt.expectedStatus {
				t.Errorf("Status = %q; want %q", response.Status, tt.expectedStatus)
			}
			if response.Service != tt.expectedService {
				t.Errorf("Service = %q; want %q", response.Service, tt.expectedService)
			}
		})
	}
}
