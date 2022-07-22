package black

import (
	"blog/internal/model"
	"blog/pkg/common/auth"
	"blog/pkg/common/bcrypt"
	"blog/pkg/mysql"
	"blog/pkg/redis"
	"blog/pkg/request"
	"blog/pkg/response"
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, params request.UserLogin) (response.UserInfoWithToken, error) {
	var userModel model.User
	if err := mysql.DB.Table(userModel.TableName()).
		Where("user_name = ? AND deleted = 0", params.UserName).
		Scan(&userModel).Error; err != nil {
		return response.UserInfoWithToken{}, err
	}
	if bool := bcrypt.PasswordVerify(params.Password, userModel.Password); !bool {
		return response.UserInfoWithToken{}, errors.New("密码错误")
	}

	// 生成token
	token, err := auth.CreateToken(params)
	if err != nil {
		return response.UserInfoWithToken{}, err
	}
	// 将token放入redis，并设置key的过期时间与token过期时间一致
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	redis.RedisClient.Set(ctx, params.UserName, token, 7*24*time.Hour)
	defer cancel()

	res := response.UserInfoWithToken{
		UserInfo: *conv(&userModel),
		Token:    token,
	}
	return res, err
}

func conv(user *model.User) *response.UserInfo {
	return &response.UserInfo{
		ID:         user.ID,
		UserName:   user.UserName,
		NickName:   user.NickName,
		Email:      user.Email,
		Avatar:     user.Avatar,
		CreateTime: user.CreateTime.Unix(),
		UpdateTime: user.UpdateTime.Unix(),
	}
}
