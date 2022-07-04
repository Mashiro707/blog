package white

import (
	"blog/internal/model"
	common "blog/pkg/common/gorm"
	"blog/pkg/mysql"
	"blog/pkg/request"
	"blog/pkg/response"
	"blog/pkg/zap"

	"github.com/gin-gonic/gin"
)

func ArticleIndexList(c *gin.Context, param request.Pagination) ([]response.ArticleIndexList, int64, error) {
	var (
		total            int64
		articleModel     model.Article
		articleIndexList []response.ArticleIndexList
	)

	if err := mysql.DB.Table(articleModel.TableName()).
		Where("deleted = 0").
		Count(&total).Error; err != nil {
		zap.ErrorLog(err)
		return articleIndexList, total, err
	}

	if err := mysql.DB.Table(articleModel.TableName() + " a").
		Select("a.id, a.title, a.description, a.is_top, a.category_id, c.id, c.name").
		Joins("join category c on c.id = a.categoryId").
		Where("deleted = 0").
		Scopes(common.Paginate(param.PageNum, param.PageSize)).
		Scan(&articleIndexList).Error; err != nil {
		zap.ErrorLog(err)
		return articleIndexList, total, err
	}

	return articleIndexList, total, nil
}
