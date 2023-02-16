package controller

import (
	"fmt"
	"go_deliver/util"

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

	// 设置订单区间
	list_begin := 0
	list_last := 10000
	for i := list_begin; i < list_last; i++ {
		MakeListLocal()
	}

	util.ClearScreen()
	fmt.Println("订单已生成！")
	fmt.Println("请继续选择需要执行的操作")
}

// func Button2(db *gorm.DB,uid int) {

// }
