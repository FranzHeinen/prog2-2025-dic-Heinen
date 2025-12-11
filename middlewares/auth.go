package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var condition = "weather-secure-xyz-2025"

		header := c.GetHeader("X-API-KEY")

		if header != condition {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: missing or invalid header",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
