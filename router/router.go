package router

import (
	"github.com/all-skeleton/gin-skeleton/app/api"
	"github.com/all-skeleton/gin-skeleton/app/middleware"
	"github.com/all-skeleton/gin-skeleton/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger()).Use(gin.Recovery())
	gin.SetMode(config.App.RunMode)

	r.GET("/api/wx-auth", api.WxMiNiLogin)
	
	apiGroup := r.Group("/api").Use(middleware.RateLimiter()).Use(middleware.Api())
	{
		// todo
		apiGroup.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Hello~~~",
			})
		})
	}
	return r
}
