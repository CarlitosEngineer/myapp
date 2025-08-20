package countries

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(g *echo.Group, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewController(svc)

	r := g.Group("/countries")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
