package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(r *gin.Engine) {
	r.Use(CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
