package model

type Category struct {
	ID      int64  `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"id"`
	Name    string `gorm:"column:name;type:varchar(64);NOT NULL" json:"name"`
	Deleted int64  `gorm:"column:deleted;type:tinyint(4);default:0;NOT NULL" json:"deleted"`
	TimeModel
}

type CategoryInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (m *Category) TableName() string {
	return "category"
}
