package utils

import (
	"fmt"
	"testing"
)

func TestGetOrdinalFormat(t *testing.T) {
	for i := 1; i < 100; i++ {
		fmt.Print(GetOrdinalFormat(uint(i)), " ")
	}
}
