package black

import (
	"blog/internal/service/black"
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/msg"
	"blog/pkg/request"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params request.UserLogin
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	result, err := black.UserLogin(c, params)
	if err != nil {
		app.Error(c, code.LoginError, msg.LoginError)
		return
	}
	app.OK(c, msg.LoginSuccess, result)
}
