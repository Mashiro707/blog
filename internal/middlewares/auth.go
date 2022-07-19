package middlewares

import (
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/common/auth"
	"blog/pkg/msg"
	"blog/pkg/redis"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authorization")

		if headerToken == "" || !strings.HasPrefix(headerToken, "Bearer ") {
			c.Abort()
			app.Error(c, code.AuthNotValid, msg.AuthNotValid)
			return
		}
		tokenString := headerToken[7:]
		token, clamis, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.Abort()
			app.Error(c, code.AuthNotValid, msg.AuthNotValid)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		result := redis.RedisClient.Get(ctx, clamis.Username)
		if result.Err() != nil {
			c.Abort()
			cancel()
			app.Error(c, code.AuthNotValid, msg.AuthNotValid)
			return
		}
		c.Next()
		defer cancel()
	}
}
