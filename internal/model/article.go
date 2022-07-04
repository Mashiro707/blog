package model

type Article struct {
	ID          int    `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	Title       string `gorm:"column:title;type:varchar(64);NOT NULL" json:"title"`
	Cover       string `gorm:"column:cover;type:varchar(255);NOT NULL" json:"cover"`
	Content     string `gorm:"column:content;type:text" json:"content"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Status      int8   `gorm:"column:status;type:tinyint(4);default:0;NOT NULL" json:"status"`
	IsTop       int8   `gorm:"column:is_top;type:tinyint(4);default:0;NOT NULL" json:"is_top"`
	IsComment   int8   `gorm:"column:is_comment;type:tinyint(4);default:0;NOT NULL" json:"is_comment"`
	CategoryID  int    `gorm:"column:category_id;type:int(11);NOT NULL" json:"category_id"`
	Deleted     int8   `gorm:"column:deleted;type:tinyint(4);default:0;NOT NULL" json:"deleted"`
	TimeModel
}

func (m *Article) TableName() string {
	return "article"
}
