package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware adds CORS headers to every response
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			log.Println("Inside CORSMiddleware preflight request from", c.ClientIP())
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
