package black

import (
	"blog/internal/model"
	"blog/pkg/mysql"
	"blog/pkg/request"
	"blog/pkg/zap"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context, params request.CreateCategoryRequest) (err error) {
	var (
		categoryModel model.Category
	)
	if err = mysql.DB.Table(categoryModel.TableName()).Create(model.Category{
		Name: params.Name,
	}).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func UpdateCategory(c *gin.Context, params request.UpdateCategoryRequest) (err error) {
	var (
		categoryModel model.Category
	)
	if err = mysql.DB.Table(categoryModel.TableName()).
		Where("id = ? AND deleted = 0", params.ID).
		Update("name", params.Name).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func DeleteCategory(c *gin.Context, params request.DeleteCategoryRequest) (err error) {
	var (
		categoryModel model.Category
	)
	if err = mysql.DB.Table(categoryModel.TableName()).
		Where("id = ? AND deleted = 0", params.ID).
		Update("deleted", 1).Error; err != nil {
		zap.ErrorLog(err)
		return err
	}
	return nil
}

func GetCategoryByID(c *gin.Context, params int) (categoryName string, err error) {
	var (
		categoryModel model.Category
	)
	if err = mysql.DB.Table(categoryModel.TableName()).
		Where("id = ? AND deleted = 0", params).
		Scan(&categoryName).Error; err != nil {
		zap.ErrorLog(err)
		return categoryName, err
	}
	return categoryName, nil
}

func GetCategoryList(ctx *gin.Context) (res []model.Category, err error) {
	var categoryModel model.Category
	if err = mysql.DB.Table(categoryModel.TableName()).
		Where("deleted = 0").Scan(&res).Error; err != nil {
		zap.ErrorLog(err)
		return res, err
	}
	return res, err
}
