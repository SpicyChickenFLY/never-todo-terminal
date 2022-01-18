package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/imroc/biu"
)

// func TestExplainSchedule(t *testing.T) {
// 	testCases := []string{
// 		"15 7 * * 1-5",     // get up for work
// 		"50 21 * * *",      // time for sleep
// 		"0 8-19/2 * * 1-5", // stand up and drink water
// 		"10 10 * 7,8 *",    // finish homework in summer holiday
// 		"0 0 31 4 *",       // bad expression
// 		"0 0 30 2 *",       // bad expression
// 	}

// 	for _, testCase := range testCases {
// 		p, err := NewPlan(testCase)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			continue
// 		}

// 		p.Explain()
// 	}
// }

// func Test_calcDatesByDOM(t *testing.T) {
// 	testCases := [][]interface{}{
// 		{2021, 8, []uint{1, 7}},
// 		{2021, 12, []uint{3, 5}},
// 		{2022, 1, []uint{2, 4}},
// 		{2022, 2, []uint{3, 5}},
// 	}
// 	for _, testCase := range testCases {
// 		result := convertDOWToDOM(uint(testCase[0].(int)), uint(testCase[1].(int)), testCase[2].([]uint))
// 		fmt.Println(result)
// 	}
// }

func Test_a(t *testing.T) {
	a := ^uint64(0)
	fmt.Println(biu.ToBinaryString(a))

	fmt.Println(time.Date(2001, 2, 0, 0, 0, 0, 0, time.UTC))
}
