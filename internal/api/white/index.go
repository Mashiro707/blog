package white

import (
	"blog/internal/service/white"
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/msg"
	"blog/pkg/request"

	"github.com/gin-gonic/gin"
)

func ArticleIndexList(c *gin.Context) {
	var param request.Pagination
	if err := c.ShouldBind(&param); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}

	res, total, err := white.ArticleIndexList(c, param)
	if err != nil {
		app.Error(c, code.DBGetError, msg.DBGetError)
		return
	}
	app.PageOK(c, res, total, param.PageNum, param.PageSize, msg.DBGetSuccess)
}
