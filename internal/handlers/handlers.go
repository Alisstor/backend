package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "backend",
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "you are authenticated"})
}
