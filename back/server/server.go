package server

import (
	"back/api"
	"back/server/middleware"
	"back/shared/dep"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Session())
	r.Use(middleware.ApiVersion())
	r.Use(middleware.ErrorHandler())

	api.Route(r.Group("/api"))

	if err := r.Run(":9200"); err != nil {
		dep.Log.Errorf("gin serverでエラー: %+v", err)
	}
}
