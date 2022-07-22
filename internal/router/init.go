package router

import (
	"blog/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	middlewares.InitMiddlewares(r)
	pingRouter(r)
	g := r.Group("/api", middlewares.CORS())
	// admin management
	AdminManagement(g)
	// blog show
	BlogShow(g)
	return r
}
func pingRouter(r *gin.Engine) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong!",
		})
	})
}
