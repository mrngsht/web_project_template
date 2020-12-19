package middleware

import (
	"back/conf"
	"back/internal/httperr"
	"golang.org/x/xerrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errCount := len(c.Errors)
		if errCount > 1 {
			conf.Log.Warnf("複数のエラーがcontextに登録されました。ハンドリングに使用されなかったエラー: \n%s", c.Errors[:errCount].String())
		}

		lastErr := c.Errors.Last()
		if lastErr == nil {
			return
		}

		var msgResp *httperr.StrMsgResponse
		if xerrors.As(lastErr, &msgResp) {
			c.AbortWithStatusJSON(msgResp.HttpStatus, gin.H{"message": msgResp.Message})
			return
		}

		// 共通的にハンドリングしたいエラー型を追加する場合はこのあたりに書く

		conf.Log.Errorf("unhandled server error: %+v", lastErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
}
