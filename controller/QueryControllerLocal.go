package controller

import (
	"fmt"
	"go_deliver/model"
	"go_deliver/util"

	"github.com/jinzhu/gorm"
)

func ShowQuery() string {
	leading := "查询系统使用方法如下：输入你想查询的用户uid，我们会根据您所提供的uid显示订单并计算总金额"
	return leading
}

func CheckQuery(uid int, DB *gorm.DB) bool {
	var deliverlist model.DeliverList
	//判断ID是否存在
	DB.Where("uid = ?", uid).First(&deliverlist)
	if deliverlist.ID != 0 {
		return true
	}
	return false

}

func QueryLists(uid int, DB *gorm.DB) {

	util.ClearScreen()
	// var deliverlist model.DeliverList
	var deliverlists []model.DeliverList

	DB.Where("uid = ?", uid).Find(&deliverlists)
	// DB.Where(model.DeliverList{Uid: uid}).Find(&deliverlist)
	fmt.Println("查询信息如下:\n", deliverlists)

	Count := SumRow(DB, uid)
	fmt.Println("合计金额为:%d", Count)
	fmt.Println("请继续选择需要执行的操作")
}

func SumRow(db *gorm.DB, uid int) int {
	rows, err := db.Table("deliver_lists").Select(" sum(price) AS total").Where("uid=?", uid).Rows()
	if err != nil {
		fmt.Println("报错: ", err)
		//return 0, err
	}
	defer rows.Close()
	total := 0
	if rows.Next() {

		err := rows.Scan(&total)
		if err != nil {
			fmt.Println("报错:", err)
			//return 0, err
		}

	}
	return total
}
