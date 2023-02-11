package util

import (
	"math"
	"math/rand"
	"time"
)

func RandomId(n int) int {
	min := 1000
	max := 2000
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min
	return id
}

func probOf(weight int) float64 {
	return 1 / float64(weight)
}

func RandomWeight(f float64) int {
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())

	val := rand.Float64()
	total_prob := 0.0
	for i := min; i <= max; i++ {
		total_prob += probOf(i)
	}

	val *= total_prob

	for i := min; i <= max; i++ {
		val -= probOf(i)
		if val < 0 {
			return i
		}
	}

	return 0
}

func Throw4Add5(f float64) int {
	// 四舍五入函数
	return int(math.Floor(f + 0.5))
}
