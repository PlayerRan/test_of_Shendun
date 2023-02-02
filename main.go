package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	//"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DeliverList struct {
	gorm.Model
	Weight int `gorm:"type:int(4);not null"`
	Uid    int `gorm:"type:int(4);not null"`
	Price  int `gorm:"type:int(4);not null"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP` // ;<-:create" json:"created_at,omitempty"
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/deliver", func(ctx *gin.Context) {
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
		newDeliverList := DeliverList{
			Weight: w,
			Uid:    int(u),
			Price:  p,
		}
		db.Create(&newDeliverList)
		//返回结果
		ctx.JSON(200, gin.H{
			"message ": "成功下单"})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}
func RandomInt(n int) int {
	min := 1000000
	max := 2000000
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min
	return id
}

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
	db.AutoMigrate(&DeliverList{})
	return db
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
