package main

import (
	"fmt"
	"go_deliver/common"
	"go_deliver/controller"
	"go_deliver/util"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "go_deliver/model"
)

func main() {

	//程序启动时调用数据库状态并保持数据库关闭状态
	db := common.InitDB()
	fmt.Printf("欢迎来到订单系统!")
	fmt.Println("请问您想执行以下何种操作？")
	var f int
	for {
		controller.MenuItems()

		fmt.Scanf("%d\n", &f)
		if f == 1 {
			controller.Button1(db)
		} else if f == 2 {
			fmt.Printf("%v\n", controller.ShowQuery())
			var uid int
			fmt.Scanf("%d\n", &uid)
			idExist := controller.CheckQuery(uid, db)
			if idExist == false {
				util.ClearScreen()
				fmt.Println("用户不存在，退回初始节点!")
				fmt.Println("请重新选择需要执行的操作")
			} else {
				controller.QueryLists(uid, db)
			}
		} else if f == 3 {
			// controller.DeleteAll(db)
			defer db.Close()
			db.Exec("DELETE FROM deliver_lists")
			db.Exec("DELETE FROM sqlite_sequence WHERE name = 'deliver_lists'")
			return
		} else {
			util.ClearScreen()
			fmt.Println("输入错误，请重新输入！")
			fmt.Scanf("%d\n", &f)
		}
	}

}
