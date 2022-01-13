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
		fields, err := parseCronStr(testCase, false)
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

func Test_calcDatesByDOM(t *testing.T) {
	testCases := [][]interface{}{
		{2021, 8, []uint{1, 7}},
		{2021, 12, []uint{3, 5}},
		{2022, 1, []uint{2, 4}},
		{2022, 2, []uint{3, 5}},
	}
	for _, testCase := range testCases {
		result := convertDOWToDOM(uint(testCase[0].(int)), uint(testCase[1].(int)), testCase[2].([]uint))
		fmt.Println(result)
	}
}
