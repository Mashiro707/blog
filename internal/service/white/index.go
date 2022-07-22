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

func ArticleIndexList(c *gin.Context, param *request.Pagination) (articleList []response.ArticleList, total int64, err error) {
	var (
		articleModel model.Article
		articleInfo  []model.ArticleInfo
	)

	if err = mysql.DB.Table(articleModel.TableName()).
		Where("deleted = 0").
		Count(&total).Error; err != nil {
		zap.ErrorLog(err)
		return articleList, total, err
	}

	if err = mysql.DB.Table(articleModel.TableName() + " a").
		Select("a.id, a.title, a.description, a.is_top, a.category_id, a.create_time, a.update_time, a.category_id, c.name").
		Joins("join category c on c.id = a.category_id").
		Where("a.deleted = 0").
		Order("a.is_top desc").
		Order("a.id desc").
		Scopes(common.Paginate(param.PageNum, param.PageSize)).
		Scan(&articleInfo).Error; err != nil {
		zap.ErrorLog(err)
		return articleList, total, err
	}
	for _, a := range articleInfo {
		articleList = append(articleList, *convArticle(&a))
	}

	return articleList, total, nil
}

func convArticle(articleInfo *model.ArticleInfo) *response.ArticleList {
	return &response.ArticleList{
		ID:           articleInfo.ID,
		Title:        articleInfo.Title,
		Description:  articleInfo.Description,
		Content:      articleInfo.Content,
		IsTop:        articleInfo.IsTop,
		IsComment:    articleInfo.IsComment,
		CategoryID:   articleInfo.CategoryID,
		CategoryName: articleInfo.CategoryName,
		CreateTime:   articleInfo.CreateTime.Unix(),
		UpdateTime:   articleInfo.UpdateTime.Unix(),
	}
}
