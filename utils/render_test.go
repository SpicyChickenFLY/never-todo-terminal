package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdinalFormat(t *testing.T) {
	for i := 1; i < 100; i++ {
		fmt.Print(GetOrdinalFormat(uint(i)), " ")
	}
}

func TestParseCronString(t *testing.T) {
	testCases := []string{
		"15 7 * * 1-5",     // get up for work
		"50 21 * * *",      // time for sleep
		"0 8-19/2 * * 1-5", // stand up and drink water
		"10 10 * 7,8 *",    // finish homework in summer holiday
	}
	for _, testCase := range testCases {
		fields, err := parseCronString(testCase)
		assert.Nil(t, err)
		for _, field := range fields {
			fmt.Println(field.explanation, field.values)
		}
	}
}

func TestExplainSchedule(t *testing.T) {
	testCases := []string{
		"15 7 * * 1-5",     // get up for work
		"50 21 * * *",      // time for sleep
		"0 8-19/2 * * 1-5", // stand up and drink water
		"10 10 * 7,8 *",    // finish homework in summer holiday
	}
	for _, testCase := range testCases {
		fmt.Println(ExplainSchedule(testCase, false))
	}
}
