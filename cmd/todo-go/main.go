package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "todo-go/docs/swagger"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo API
// @version 1.0
// @description This is a sample server for a todo application.
// @BasePath /
func main() {
	mux := http.NewServeMux()
	logger := slog.Default()

	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/swagger/", httpSwagger.Handler())

	appAddress := "localhost:8000"
	httpAppAddress := fmt.Sprintf("http://%s", appAddress)
	logger.Info("Starting to listen at", "address", httpAppAddress)
	if err := http.ListenAndServe(appAddress, mux); err != nil {
		logger.Error("Failed to listen", "address", appAddress, "error", err)
		os.Exit(1)
	}
}

// healthCheckHandler godoc
// @Summary Health Check
// @Description Returns OK if service is up
// @Tags health
// @Success 200
// @Router /healthcheck [get]
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
