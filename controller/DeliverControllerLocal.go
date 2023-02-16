package controller

import (
	"go_deliver/common"
	"go_deliver/model"
	"go_deliver/util"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func MakeListLocal() {
	DB := common.GetDB()

	// DeleteAll(DB)
	w := util.RandomWeight()
	//生成随机id和重量，并生成订单
	uid := util.RandomId()
	// w := int(math.Ceil(weight))
	p := CalPrice2(w)
	newDeliverList := model.DeliverList{
		Uid:    uid,
		Weight: w,
		Price:  p,
	}
	DB.Create(&newDeliverList)
}

func CalPrice2(w int) int {
	// 计算订单价格
	if w == 1 {
		return 18
	} else {
		var p1 int = 13 + 5*w
		var p2 float64
		for i := 2; i < w; i++ {
			var pre_p2 float64 = 0.0
			p2 = 0.01 * (float64(8+5*i) + pre_p2)
			pre_p2 = p2
		}
		return util.Throw4Add5(float64(p1) + p2)
	}
}

// func DeleteAll(db *gorm.DB) {
// 	db = common.GetDB()
// 	dl := &model.DeliverList{}
// 	err := db.First(&dl).Error
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.Model(&model.DeliverList{}).Delete("id = ?", dl.Uid).Error
// 	if err != nil {
// 		panic(err)
// 	}
// }
