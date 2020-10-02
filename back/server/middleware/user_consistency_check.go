package middleware

import (
	"back/server/context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserConsistencyCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.FullPath(), "getLoginUser") {
			c.Next()
			return
		}

		loginUserID := context.GetLoginUser(c).ID
		frontUserID := c.GetHeader("x-user-id")
		if strconv.Itoa(loginUserID) != frontUserID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "ログイン中のユーザと異なるユーザでの操作が行われました。"})
			return
		}
		c.Next()
	}
}
