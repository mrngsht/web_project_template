package pub

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
