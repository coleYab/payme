package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func CacheMiddleware(client *redis.Client, ctx *gin.Context, next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Next()
	}
}
