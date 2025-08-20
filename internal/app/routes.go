package app

import (
	// "github.com/labstack/echo/v4"
	"myapp/internal/modules/countries" // ← cambia module path
)

func (s *Server) RegisterRoutes() {
	api := s.E.Group("/api/v1") // CODE VIEJO
	// var api *echo.Group = s.E.Group("/api/v1")

	// módulos
	countries.Register(api, s.DB)
	_ = api
}
