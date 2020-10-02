package context

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const ContextKeyLoginUser = "login_user"

type LoginUser struct {
	ID      int
	LoginID string
	Name    string
}

func GetLoginUser(c *gin.Context) LoginUser {
	val, ok := c.Get(ContextKeyLoginUser)
	if !ok {
		panic("contextからlogin_userが取得できませんでした")
	}

	loginUser, ok := val.(LoginUser)
	if !ok {
		panic(fmt.Errorf("contextから取得したlogin_user[%v]が既知の形式ではありませんでした", val))
	}

	return loginUser
}
