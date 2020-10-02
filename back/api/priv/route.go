package priv

import (
	"back/api/priv/session"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	session.Route(r.Group("/session"))
}
