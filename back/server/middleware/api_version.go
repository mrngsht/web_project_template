package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BackendApiVersion = "0.1"

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		frontendApiVersion := c.GetHeader("x-api-version")
		if frontendApiVersion != BackendApiVersion {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid api version"})
			return
		}
		c.Next()
	}
}
