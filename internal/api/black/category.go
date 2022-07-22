package black

import (
	"blog/internal/service/black"
	"blog/pkg/app"
	"blog/pkg/code"
	"blog/pkg/msg"
	"blog/pkg/request"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var prarm request.CreateCategoryRequest
	if err := ctx.ShouldBind(&prarm); err != nil {
		app.Error(ctx, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	err := black.CreateCategory(ctx, prarm)
	if err != nil {
		app.Error(ctx, code.DBCreateError, msg.DBCreateError)
		return
	}
	app.OK(ctx, msg.DBCreateSuccess, nil)
}

func UpdateCategory(ctx *gin.Context) {
	var prarm request.UpdateCategoryRequest
	if err := ctx.ShouldBind(&prarm); err != nil {
		app.Error(ctx, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	err := black.UpdateCategory(ctx, prarm)
	if err != nil {
		app.Error(ctx, code.DBUpdateError, msg.DBUpdateError)
		return
	}
	app.OK(ctx, msg.DBUpdateSuccess, nil)
}

func DeleteCategory(ctx *gin.Context) {
	var prarm request.DeleteCategoryRequest
	if err := ctx.ShouldBind(&prarm); err != nil {
		app.Error(ctx, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	err := black.DeleteCategory(ctx, prarm)
	if err != nil {
		app.Error(ctx, code.DBDelError, msg.DBDelError)
		return
	}
	app.OK(ctx, msg.DBDelSuccess, nil)
}

func GetCategoryList(ctx *gin.Context) {
	var params request.Pagination
	if err := ctx.ShouldBind(&params); err != nil {
		app.Error(ctx, code.ParamsNotValid, msg.ParamsNotValid)
		return
	}
	res, err := black.GetCategoryList(ctx, &params)
	if err != nil {
		app.Error(ctx, code.DBGetError, msg.DBGetError)
	}
	app.OK(ctx, msg.DBGetSuccess, res)
}
