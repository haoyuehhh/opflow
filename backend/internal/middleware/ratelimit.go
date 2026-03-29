package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

type RateLimiter struct {
	limiter ratelimit.Limiter
}

func NewRateLimiter(rate int) *RateLimiter {
	return &RateLimiter{
		limiter: ratelimit.New(rate),
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl.limiter.Take()
		c.Next()
	}
}

func RateLimitMiddleware(rate int) gin.HandlerFunc {
	limiter := NewRateLimiter(rate)
	return limiter.Middleware()
}