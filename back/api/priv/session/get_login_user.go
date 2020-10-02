package session

import (
	"back/server/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, context.GetLoginUser(c))
}
