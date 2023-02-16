package model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

// "gorm.io/driver/mysql"
// "gorm.io/gorm"

type DeliverList struct {
	ID      uint      `gorm:"primarykey;not null"`
	Created time.Time `gorm:"datatime(0);not null"`
	Uid     int       `gorm:"type:int(4);not null"`
	Weight  int       `gorm:"type:int(4);not null"`
	Price   int       `gorm:"type:int(4);not null"`
}
