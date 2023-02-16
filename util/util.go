package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func RandomId() int {
	min := 1000
	max := 2000
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min
	return id
}

func probOf(weight int) float32 {
	return 1 / float32(weight)
}

func RandomWeight() int {
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())

	val := rand.Float32()
	total_prob := float32(0.0)
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

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
