package controller

import (
	"go_deliver/common"
	"go_deliver/model"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MakeList(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	weight := ctx.PostForm("weight")
	uid := ctx.PostForm("uid")
	//判断是否合规
	w0, _ := (strconv.ParseFloat(weight, 64))
	w := int(math.Ceil(w0))
	if w < 0 || w > 100 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "重量必须在0-100kg之间!"})
		return
	}

	u, _ := strconv.ParseInt(uid, 10, 32)
	if u < 1000 || u > 2000 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "用户不存在!"})
		return
	}
	//log.Println(w, u)
	//计算订单价格
	p := CalPrice(w)
	//创建订单
	newDeliverList := model.DeliverList{
		Weight: w,
		Uid:    int(u),
		Price:  p,
	}
	DB.Create(&newDeliverList)
	//返回结果
	ctx.JSON(200, gin.H{
		"message ": "成功下单"})
}

func CalPrice(w int) int {
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
		return Throw4Add5(float64(p1) + p2)
	}
}
func Throw4Add5(f float64) int {
	// 四舍五入函数
	return int(math.Floor(f + 0.5))
}
