package utils

import "time"

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
