package white

import (
	"blog/internal/service/white"
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/msg"
	"blog/pkg/request"

	"github.com/gin-gonic/gin"
)

func ArticleDetailByID(c *gin.Context) {
	var param request.GetArticleByID
	if err := c.ShouldBind(&param); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}

	res, err := white.ArticleDetail(c, &param)
	if err != nil {
		app.Error(c, code.DBGetError, msg.DBGetError)
		return
	}
	app.OK(c, msg.DBGetSuccess, res)
}
