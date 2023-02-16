package controller

import (
	"fmt"
	"go_deliver/model"
	"go_deliver/util"
	"time"

	"github.com/jinzhu/gorm"
)

func MenuItems() {
	fmt.Println("1.生成订单")
	fmt.Println("2.查询订单")
	fmt.Println("3.退出系统")
}

func Button1(db *gorm.DB) {

	//初始化数据库
	db.Exec("DELETE FROM deliver_lists")
	db.Exec("DELETE FROM sqlite_sequence WHERE name = 'deliver_lists'")
	// 设置订单数量
	list_length := 1000000

	tx := db.Begin()

	for i := 1; i <= list_length; i++ {
		w := util.RandomWeight()
		//生成随机id和重量，并生成订单
		uid := util.RandomId()
		// w := int(math.Ceil(weight))
		p := CalPrice2(w)
		newDeliverList := model.DeliverList{
			ID:      uint(i),
			Created: time.Now(),
			Uid:     uid,
			Weight:  w,
			Price:   p,
		}
		tx.Create(&newDeliverList)
		// err := tx.Create(&newDeliverList).Error
		// if err != nil {
		// 	fmt.Println(err)
		// 	tx.Rollback()
		// }
	}

	tx.Commit()
	fmt.Println("订单已生成！")
	fmt.Println("请继续选择需要执行的操作")
}

// func Button2(db *gorm.DB,uid int) {

// }
