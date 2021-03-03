package main

import (
	"math"
	"math/rand"
	"time"
)

var Rand *rand.Rand

func init() {
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// calculate growth ratio from a to b in percentage
func GrowthRatioInPerc(a, b int) (ratio float64) {
	if a == 0 {
		a, b = 1, b+1
	}

	return math.Round((float64(b)-float64(a))*1e6/float64(a)) / 1e4
}

func Divmod(a, b int) (d, m int) {
	return a / b, a % b
}

func Percentage(a, b int64) (per float64) {
	return math.Round(float64(a)*1e4/float64(b)) / 1e2
}
