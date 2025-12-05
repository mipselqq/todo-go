package handlers

import "net/http"

// healthCheckHandler godoc
// @Summary Health Check
// @Description Returns OK if service is up
// @Tags health
// @Success 200
// @Router /healthcheck [get]
func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
