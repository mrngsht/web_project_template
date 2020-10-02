package session

import (
	"back/server/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie(middleware.GoApiSessionName, "", -1, "/", "", true, true) // delete cookie
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
