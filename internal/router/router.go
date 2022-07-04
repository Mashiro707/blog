package router

import (
	"blog/internal/api/black"
	"blog/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminManagement(g *gin.RouterGroup) {
	g.POST("/login", black.Login)
	administrator := g.Group("/admin", middlewares.Authentication())
	{
		article := administrator.Group("/article")
		{
			article.POST("/create", black.CreateArticle)
			article.POST("/update", black.UpdateArticle)
			article.POST("/delete", black.DeleteArticle)
			article.POST("/getArticle", black.GetArticleByID)
			article.POST("/getArticleList", black.ArticleList)
		}
		category := administrator.Group("/category")
		{
			category.POST("/create", black.CreateCategory)
			category.POST("/update", black.UpdateCategory)
			category.POST("/delete", black.DeleteCategory)
			category.POST("/getCategoryList", black.GetCategoryList)
		}
	}

}

func PageDisplay(g *gin.RouterGroup) {

}
