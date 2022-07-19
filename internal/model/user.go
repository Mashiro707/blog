package model

import (
	"blog/pkg/mysql"
)

type User struct {
	ID       int64  `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	UserName string `gorm:"column:user_name;NOT NULL" json:"user_name"`
	Password string `gorm:"column:password;NOT NULL" json:"password"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
	Email    string `gorm:"column:email" json:"email"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Deleted  int64  `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`
	TimeModel
}

type UserInfo struct {
	ID           int64  `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	UserName     string `gorm:"column:user_name;NOT NULL" json:"user_name"`
	Password     string `gorm:"column:password;NOT NULL" json:"password"`
	NickName     string `gorm:"column:nick_name" json:"nick_name"`
	Email        string `gorm:"column:email" json:"email"`
	Avatar       string `gorm:"column:avatar" json:"avatar"`
	Deleted      int64  `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`
	CategoryID   int64  `gorm:"column:category_id;type:bigint(20)" json:"category_id"`
	CategoryName string `gorm:"column:cagegory_name;type:varchar(64)" json:"category_name"`
	TimeModel
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) GetUserInfoByID(id int) (User, error) {
	var user User
	var err error
	if err = mysql.DB.Table(m.TableName()).
		Where("id = ? AND deleted = 0", id).
		First(&user).Error; err != nil {
		return user, err
	}
	return user, err
}
