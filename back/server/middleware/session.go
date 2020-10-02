package middleware

import (
	"net/http"

	"back/shared/dep"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const GoApiSessionName = "go_api_session"

func Session() gin.HandlerFunc {
	store, err := redis.NewStore(10, "tcp", dep.RedisHost+":"+dep.RedisPort, dep.RedisPassword, []byte(dep.GoAppKey))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   60 * 60 * 2, // 2 hours
		HttpOnly: true,
		Secure:   true,
	})

	return sessions.Sessions(GoApiSessionName, store)
}
