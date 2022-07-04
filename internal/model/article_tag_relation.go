package model

type ArticleTagRelation struct {
	ID        int `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	ArticleID int `gorm:"column:article_id;type:int(11);NOT NULL" json:"article_id"`
	TagID     int `gorm:"column:tag_id;type:int(11);NOT NULL" json:"tag_id"`
	Deleted   int `gorm:"column:deleted;type:tinyint(4);default:0;NOT NULL" json:"deleted"`
	TimeModel
}

func (m *ArticleTagRelation) TableName() string {
	return "article_tag_relation"
}
