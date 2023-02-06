package util

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	min := 1000000
	max := 2000000
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(max-min) + min
	return id
}
