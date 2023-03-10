package common

import (
	"fmt"
	"go_deliver/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "123.56.19.92"
	port := "3316"
	database := "test"
	username := "test"
	password := "114514"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	db.AutoMigrate(&model.DeliverList{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB

}
