package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/internal/handlers"
)

func TestHealthCheckHandler(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	// Act
	handlers.HealthCheckHandler(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
