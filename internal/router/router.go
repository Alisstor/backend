package router

import (
	"github.com/alisstor/backend/internal/config"
	"github.com/alisstor/backend/internal/handlers"
	"github.com/alisstor/backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/health", handlers.Health)
	r.GET("/ping", handlers.Ping)

	api := r.Group("/api")
	api.Use(middleware.Auth(cfg.JWTSecret))
	{
		api.GET("/protected", handlers.Protected)
	}

	return r
}
