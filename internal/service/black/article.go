package black

import (
	"blog/internal/model"
	common "blog/pkg/common/gorm"
	"blog/pkg/mysql"
	"blog/pkg/request"
	"blog/pkg/response"
	"blog/pkg/zap"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context, params request.CreateArticle) (err error) {
	var articleModel model.Article
	article := model.Article{
		Title:       params.Title,
		Cover:       params.Cover,
		Content:     params.Content,
		Description: params.Description,
		IsTop:       params.IsTop,
		IsComment:   params.IsComment,
		CategoryID:  params.CategoryID,
	}
	// 新增文章
	if err = mysql.DB.Table(articleModel.TableName()).Create(&article).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func UpdateArticle(c *gin.Context, params request.UpdateArticle) (err error) {
	var articleModel model.Article
	articleTagRelationMap := map[string]interface{}{
		"title":       params.Title,
		"cover":       params.Cover,
		"content":     params.Content,
		"description": params.Description,
		"is_top":      params.IsTop,
		"is_comment":  params.IsComment,
		"category_id": params.CategoryID,
	}
	// 更新文章
	if err = mysql.DB.Table(articleModel.TableName()).
		Where("id = ? AND deleted = 0", params.ID).
		Updates(articleTagRelationMap).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func DeleteArticle(c *gin.Context, params request.DeleteArticle) (err error) {
	var articleModel model.Article
	// 维护文章表
	if err = mysql.DB.Table(articleModel.TableName()).
		Where("id = ? AND deleted = 0", params.ID).
		Updates(map[string]int{
			"deleted": 1,
		}).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return err
}

func GetArticleByID(param request.GetArticleByID) (articleDetail response.ArticleDetail, err error) {
	var (
		articleModel model.Article
	)
	if err = mysql.DB.Table(articleModel.TableName()).
		Select("article.title, article.cover, article.content, article.description, article.is_top, article.is_comment, article.category_id, category.name category_name").
		Joins("join category on category.id = article.category_id").
		Where("article.id = ? AND article.deleted = 0", param.ID).Scan(&articleDetail).Error; err != nil {
		zap.ErrorLog(err)
		return articleDetail, err
	}
	return articleDetail, nil
}

func ArticleList(params *request.Pagination) (articleList []response.ArticleList, err error) {
	var (
		articleModel     model.Article
		articleInfoArray []model.ArticleInfo
	)
	if err = mysql.DB.Table(articleModel.TableName()).
		Select("article.id, article.title, article.cover, article.content, article.is_top, article.is_comment, article.category_id, category.name category_name, article.created_time, article.updated_time").
		Joins("join category on category.id = article.category_id").
		Scopes(common.Paginate(params.PageNum, params.PageSize)).
		Where("article.deleted = 0").Scan(&articleInfoArray).Error; err != nil {
		zap.ErrorLog(err)
		return articleList, err
	}
	for _, v := range articleInfoArray {
		articleList = append(articleList, *convArticle(&v))
	}
	return articleList, nil
}

func convArticle(articleInfo *model.ArticleInfo) *response.ArticleList {
	return &response.ArticleList{
		ID:           articleInfo.ID,
		Title:        articleInfo.Title,
		Status:       articleInfo.Status,
		IsTop:        articleInfo.IsTop,
		IsComment:    articleInfo.IsComment,
		CategoryID:   articleInfo.CategoryID,
		CategoryName: articleInfo.CategoryName,
		CreatedTime:  articleInfo.CreatedTime.Unix(),
		UpdatedTime:  articleInfo.UpdatedTime.Unix(),
	}
}
