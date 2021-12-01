package utils

import (
	"math"
	"time"
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

// LessInAbs compare absolute value of two int
func LessInAbs(a, b int) bool {
	return Abs(a) < Abs(b)
}

// LessInID compare ID value
func LessInID(a, b int) bool {
	return LessInAbs(a, b)
}

// LessInTime compare time
func LessInTime(a, b time.Time) bool {
	return a.Before(b)
}
