package utils

import (
	"math"
)

func min(x int, y int, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

// Abs return absolute value
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}
