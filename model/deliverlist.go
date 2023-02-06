package model

import "github.com/jinzhu/gorm"

type DeliverList struct {
	gorm.Model
	Weight int `gorm:"type:int(4);not null"`
	Uid    int `gorm:"type:int(4);not null"`
	Price  int `gorm:"type:int(4);not null"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP` // ;<-:create" json:"created_at,omitempty"
}
