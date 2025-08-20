package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	E  *echo.Echo
	DB *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	e := echo.New()
	return &Server{E: e, DB: db}
}

func (s *Server) Start(addr string) error {
	go func() {
		_ = s.E.Start(addr)
	}()
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.E.Shutdown(ctx)
}

func (s *Server) HealthRoutes() {
	s.E.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	s.E.GET("/health/db", func(c echo.Context) error {
		sqlDB, err := s.DB.DB()
		if err != nil {
			return c.String(http.StatusInternalServerError, "db error")
		}
		ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Second)
		defer cancel()
		if err := sqlDB.PingContext(ctx); err != nil {
			return c.String(http.StatusServiceUnavailable, "db down")
		}
		return c.String(http.StatusOK, "db up")
	})
}
