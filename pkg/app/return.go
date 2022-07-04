package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, msg string) {
	var res Response
	res.Msg = msg
	res.Data = []string{}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

func OK(c *gin.Context, msg string, data interface{}) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int64, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageNo = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
func PageError(c *gin.Context, result interface{}, count int64, pageIndex int, pageSize int, code int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageNo = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}
