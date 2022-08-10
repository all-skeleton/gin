package middleware

import (
	"github.com/all-skeleton/gin-skeleton/app/api"
	"github.com/all-skeleton/gin-skeleton/app/library"
	"github.com/all-skeleton/gin-skeleton/config"
	"github.com/gin-gonic/gin"
	"time"
)

func Api() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		if token == "" {
			api.Response(c, library.ErrorAuthCheckTokenFail, "")
			c.Abort()
			return
		}

		claims, err := library.ParseApiToken(token)
		if err != nil {
			api.Response(c, library.ErrorAuthCheckTokenFail, "")
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			api.Response(c, library.ErrorAuthCheckTokenTimeout, "")
			c.Abort()
			return
		}

		if claims.Version != config.App.Version {
			api.Response(c, library.SystemVersionUpdate, "")
			c.Abort()
			return
		}

		c.Set("ApiClaims", claims)

		// todo
		c.Next()
	}
}
