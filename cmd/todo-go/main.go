// @title Todo API
// @version 1.0
// @description Simple Todo API
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	logger := slog.Default()

	// GET /healthcheck
	// @Summary Health Check
	// @Description Returns OK if service is up
	// @Success 200
	// @Router /healthcheck [get]
	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	appAddress := "localhost:8000"
	httpAppAddress := fmt.Sprintf("http://%s", appAddress)
	logger.Info("Starting to listen at", "address", httpAppAddress);
	if err := http.ListenAndServe(appAddress, mux); err != nil {
		logger.Error("Failed to listen", "address", appAddress, "error", err);
		os.Exit(1)
	}
}
