package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "todo-go/docs/swagger"
	"todo-go/internal/config"
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
	app_config, err := config.FromEnv()
	if err != nil {
		logger.Error("Failed to load app config with", "error", err)
		os.Exit(0)
	}

	fmt.Println(app_config)

	mux.HandleFunc("/healthcheck", handlers.HealthCheckHandler)
	mux.HandleFunc("/swagger/", httpSwagger.Handler())

	address := app_config.Address()
	http_address := "https://" + address

	logger.Info("Starting to listen at", "address", http_address)
	if err := http.ListenAndServe(address, mux); err != nil {
		logger.Error("Failed to listen", "address", address, "error", err)
		os.Exit(1)
	}
}
