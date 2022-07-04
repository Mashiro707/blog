package middlewares

import (
	"blog/pkg/common/auth"
	"blog/pkg/redis"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authentication")

		if headerToken == "" || !strings.HasPrefix(headerToken, "Bearer ") {
			c.Abort()
			return
		}
		tokenString := headerToken[7:]
		token, clamis, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.Abort()
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		result := redis.RedisClient.Get(ctx, clamis.Username)
		if result.Err() != nil {
			c.Abort()
			cancel()
			return
		}
		c.Next()
		defer cancel()
	}
}
