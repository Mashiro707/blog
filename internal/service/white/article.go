package white

import (
	"blog/internal/model"
	"blog/pkg/mysql"
	"blog/pkg/request"
	"blog/pkg/response"
	"blog/pkg/zap"

	"github.com/gin-gonic/gin"
)

func ArticleDetail(c *gin.Context, param request.GetArticleByID) (articleDetail response.ArticleDetail, err error) {
	var (
		articleModel model.Article
	)
	if err = mysql.DB.Table(articleModel.TableName()).
		Select("article.title, article.cover, article.content, article.is_top, article.is_comment, article.category_id, category.name category_name").
		Joins("join category on category.id = article.category_id").
		Where("id = ? AND deleted = 0", param.ID).Scan(&articleDetail).Error; err != nil {
		zap.ErrorLog(err)
		return articleDetail, err
	}
	return articleDetail, nil
}
