package middleware

import (
	"back/server/context"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-contrib/sessions"
)

const SessionKeyLoginUser = "login_user"

func init() {
	gob.Register(context.LoginUser{})
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if loginUser := session.Get(SessionKeyLoginUser); loginUser != nil {
			if _, ok := loginUser.(context.LoginUser); ok {
				c.Set(context.ContextKeyLoginUser, loginUser)
				c.Next()
				return
			}
			// okじゃない場合は非互換かもしれないので作り直し
		}

		// TODO: 認証してcontext.LoginUserを設定する
		c.Set(context.ContextKeyLoginUser, context.LoginUser{})

		c.Next()
		return
	}
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
}
func serverError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
}
