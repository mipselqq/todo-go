package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "todo-go/docs/swagger"
	"todo-go/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo API
// @version 1.0
// @description This is a sample server for a todo application.
// @BasePath /
func main() {
	mux := http.NewServeMux()
	logger := slog.Default()

	mux.HandleFunc("/healthcheck", handlers.HealthCheckHandler)
	mux.HandleFunc("/swagger/", httpSwagger.Handler())

	appAddress := "localhost:8000"
	httpAppAddress := fmt.Sprintf("http://%s", appAddress)
	logger.Info("Starting to listen at", "address", httpAppAddress)
	if err := http.ListenAndServe(appAddress, mux); err != nil {
		logger.Error("Failed to listen", "address", appAddress, "error", err)
		os.Exit(1)
	}
}

