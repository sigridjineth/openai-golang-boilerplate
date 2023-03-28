package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const HealthEndpoint = "/health"

type Health struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{
		Name:        "Chatshire EthDenver Backend",
		Description: "Chatshire Backend API for EthDenver Hackathon PoC Stage",
	})
}
