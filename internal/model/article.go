package model

type Article struct {
	ID          int64  `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"id"`
	Title       string `gorm:"column:title;type:varchar(64);NOT NULL" json:"title"`
	Cover       string `gorm:"column:cover;type:varchar(255);NOT NULL" json:"cover"`
	Content     string `gorm:"column:content;type:text" json:"content"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Status      int8   `gorm:"column:status;type:tinyint(4);default:0;NOT NULL" json:"status"`
	IsTop       int8   `gorm:"column:is_top;type:tinyint(4);default:0;NOT NULL" json:"is_top"`
	IsComment   int8   `gorm:"column:is_comment;type:tinyint(4);default:0;NOT NULL" json:"is_comment"`
	CategoryID  int64  `gorm:"column:category_id;type:bigint(20);NOT NULL" json:"category_id"`
	Deleted     int8   `gorm:"column:deleted;type:tinyint(4);default:0;NOT NULL" json:"deleted"`
	TimeModel
}

type ArticleInfo struct {
	ID           int64  `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"id"`
	Title        string `gorm:"column:title;type:varchar(64);NOT NULL" json:"title"`
	Cover        string `gorm:"column:cover;type:varchar(255);NOT NULL" json:"cover"`
	Content      string `gorm:"column:content;type:text" json:"content"`
	Description  string `gorm:"column:description;type:text" json:"description"`
	Status       int8   `gorm:"column:status;type:tinyint(4);default:0;NOT NULL" json:"status"`
	IsTop        int8   `gorm:"column:is_top;type:tinyint(4);default:0;NOT NULL" json:"is_top"`
	IsComment    int8   `gorm:"column:is_comment;type:tinyint(4);default:0;NOT NULL" json:"is_comment"`
	CategoryID   int64  `gorm:"column:category_id;type:bigint(20);NOT NULL" json:"category_id"`
	CategoryName string `gorm:"column:cagegory_name;type:varchar(64)" json:"category_name"`
	Deleted      int8   `gorm:"column:deleted;type:tinyint(4);default:0;NOT NULL" json:"deleted"`
	TimeModel
}

func (m *Article) TableName() string {
	return "article"
}
