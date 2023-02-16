package model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

// "gorm.io/driver/mysql"
// "gorm.io/gorm"

type DeliverList struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Uid       int `gorm:"type:int(4);not null"`
	Weight    int `gorm:"type:int(4);not null"`
	Price     int `gorm:"type:int(4);not null"`
}
