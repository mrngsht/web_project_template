package api

import (
	"back/api/priv"
	"back/api/pub"
	"back/server/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	priv.Route(
		r.Group(
			"/priv",
			middleware.AuthCheck(),
			middleware.UserConsistencyCheck(),
		))
	pub.Route(r.Group("/pub"))
}
