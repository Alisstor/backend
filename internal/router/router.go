package router

import (
	"github.com/alisstor/backend/internal/config"
	"github.com/alisstor/backend/internal/handlers"
	"github.com/alisstor/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)
	r.GET("/ping", handlers.Ping)

	api := r.Group("/api")
	api.Use(middleware.Auth(cfg.JWTSecret))
	{
		api.GET("/protected", handlers.Protected)
	}

	return r
}
