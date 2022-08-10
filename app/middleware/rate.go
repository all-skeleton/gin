package middleware

import (
	"context"
	"github.com/all-skeleton/gin-skeleton/app/api"
	"github.com/all-skeleton/gin-skeleton/app/library"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

func RateLimiter() gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(c *gin.Context) {
		l, _ := limiters.LoadOrStore(c.ClientIP(), rate.NewLimiter(3, 10))

		ctx, cancel := context.WithTimeout(c, 500*time.Millisecond)
		defer cancel()

		if err := l.(*rate.Limiter).Wait(ctx); err != nil {
			api.Response(c, library.SystemRateLimitErr, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
