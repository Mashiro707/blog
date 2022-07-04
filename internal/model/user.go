package model

import (
	"blog/pkg/mysql"
	"time"
)

type User struct {
	ID        int       `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	UserName  string    `gorm:"column:user_name;NOT NULL" json:"user_name"`
	Password  string    `gorm:"column:password;NOT NULL" json:"password"`
	NickName  string    `gorm:"column:nick_name" json:"nick_name"`
	Email     string    `gorm:"column:email" json:"email"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Deleted   int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
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
