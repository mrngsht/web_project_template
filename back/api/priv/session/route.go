package session

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.POST("/getLoginUser", getLoginUser)
	r.POST("/logout", logout)
	r.POST("/hello", func(c *gin.Context) { c.JSON(200, gin.H{"message": "hello"}) })
}
