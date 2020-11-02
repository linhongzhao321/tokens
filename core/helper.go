package core

import (
	"math/rand"
	"time"
)

// x^y
// Do not check for overflow
// please judge the data range according to the usage scenario
func IntPow(x int64, y int64) (z int64) {
	z = 1
	for y > 0 {
		z *= x
		y--
	}
	return
}

// !!!WARNING!!!
// This operation is prohibited on a sequence containing Elements
// that exists in Buffer.mBuf
func Shuffle(buf []string, begin int64, end int64) {
	l := end - begin
	rand.Seed(time.Now().UnixNano())
	for i := begin; i <= end; i++ {
		randIdx := rand.Int63n(l)
		buf[i], buf[randIdx] = buf[randIdx], buf[i]
	}
}
