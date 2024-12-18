package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	config Environment
}

func CreateServer(cfg Environment) *Server {
	// Create server
	s := &Server{
		Echo:   echo.New(),
		config: cfg,
	}

	s.HideBanner = true

	// Health Check
	s.GET("/api/v1/health", HealthCheck)

	return s
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
