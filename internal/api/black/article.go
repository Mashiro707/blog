package black

import (
	"blog/internal/service/black"
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/msg"
	"blog/pkg/request"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var params request.CreateArticle
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	if err := black.CreateArticle(c, params); err != nil {
		app.Error(c, code.DBCreateError, msg.DBCreateError)
		return
	}
	app.OK(c, msg.DBCreateSuccess, nil)
}

func UpdateArticle(c *gin.Context) {
	var params request.UpdateArticle
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	if err := black.UpdateArticle(c, params); err != nil {
		app.Error(c, code.DBUpdateError, msg.DBUpdateError)
		return
	}
	app.OK(c, msg.DBUpdateSuccess, nil)
}

func DeleteArticle(c *gin.Context) {
	var params request.DeleteArticle
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	if err := black.DeleteArticle(c, params); err != nil {
		app.Error(c, code.DBDelError, msg.DBDelError)
		return
	}
	app.OK(c, msg.DBDelSuccess, nil)
}

func GetArticleByID(c *gin.Context) {
	var params request.GetArticleByID
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	res, err := black.GetArticleByID(params)
	if err != nil {
		app.Error(c, code.DBGetError, msg.DBGetError)
		return
	}
	app.OK(c, msg.DBGetSuccess, res)
}

func ArticleList(c *gin.Context) {
	var params request.Pagination
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Error(c, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	res, err := black.ArticleList(&params)
	if err != nil {
		app.Error(c, code.DBGetError, msg.DBGetError)
		return
	}
	app.OK(c, msg.DBGetSuccess, res)
}
