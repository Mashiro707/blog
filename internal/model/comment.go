package model

import "time"

type TimeModel struct {
	CreatedTime time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_time"`
}
