package black

import (
	"blog/internal/model"
	"blog/pkg/mysql"
	"blog/pkg/request"
	"blog/pkg/response"
	"blog/pkg/zap"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context, params request.CreateArticle) (err error) {
	var articleModel model.Article
	var articleTagRelationModel model.ArticleTagRelation
	var articleTagRelationArray []model.ArticleTagRelation
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
	// 维护文章标签表
	for _, tagID := range params.Tags {
		articleTagRelationArray = append(articleTagRelationArray, model.ArticleTagRelation{
			ArticleID: article.ID,
			TagID:     tagID,
		})
	}
	if err = mysql.DB.Table(articleTagRelationModel.TableName()).Create(&articleTagRelationArray).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func UpdateArticle(c *gin.Context, params request.UpdateArticle) (err error) {
	var articleModel model.Article
	var articleTagRelationModel model.ArticleTagRelation
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
	// 判断标签是否更新，如果需要更新标签则需要删除原来所有标签再新增新的数据
	if len(params.Tags) == 0 {
		return nil
	} else {
		// delete tags
		if err = mysql.DB.Table(articleTagRelationModel.TableName()).
			Where("article_id = ? AND deleted = 0", params.ID).
			Updates(map[string]int{
				"deleted": 1,
			}).Error; err != nil {
			zap.ErrorLog(err)
			return err
		}
		// 维护文章标签表
		var articleTagRelationArray []model.ArticleTagRelation
		for _, tagID := range params.Tags {
			articleTagRelationArray = append(articleTagRelationArray, model.ArticleTagRelation{
				ArticleID: params.ID,
				TagID:     tagID,
			})
		}
		if err = mysql.DB.Table(articleTagRelationModel.TableName()).
			Create(&articleTagRelationArray).Error; err != nil {
			zap.ErrorLog(err)
			return err
		}
	}
	return nil
}

func DeleteArticle(c *gin.Context, params request.DeleteArticle) (err error) {
	var articleModel model.Article
	var articleTagRelationModel model.ArticleTagRelation
	// 维护文章表
	if err = mysql.DB.Table(articleModel.TableName()).
		Where("id = ? AND deleted = 0", params.ID).
		Updates(map[string]int{
			"deleted": 1,
		}).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	// 维护文章标签表
	if err = mysql.DB.Table(articleTagRelationModel.TableName()).
		Where("article_id = ? AND deleted = 0", params).
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
		Select("article.title, article.cover, article.content, article.is_top, article.is_comment, article.category_id, category.name category_name").
		Joins("join category on category.id = article.category_id").
		Where("id = ? AND deleted = 0", param.ID).Scan(&articleDetail).Error; err != nil {
		zap.ErrorLog(err)
		return articleDetail, err
	}
	return articleDetail, nil
}

func ArticleList() (articleList []response.ArticleList, err error) {
	var (
		articleModel model.Article
	)
	if err = mysql.DB.Table(articleModel.TableName()).
		Select("article.id, article.title, article.cover, article.content, article.is_top, article.is_comment, article.category_id, category.name category_name, article.created_at, article.update_at").
		Joins("join category on category.id = article.category_id").
		Where("deleted = 0").Scan(&articleList).Error; err != nil {
		zap.ErrorLog(err)
		return articleList, err
	}
	return articleList, nil
}
