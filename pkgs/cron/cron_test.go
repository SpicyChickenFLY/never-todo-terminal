package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/imroc/biu"
)

func Test_Plan_Explain(t *testing.T) {
	testCases := []string{
		"15 7 * * 1-5",     // get up for work
		"50 21 * * *",      // time for sleep
		"0 8-19/2 * * 1-5", // stand up and drink water
		"10 10 * 7,8 *",    // finish homework in summer holiday
		"0 0 31 4 *",       // bad expression
		"0 0 30 2 *",       // bad expression
	}

	for _, testCase := range testCases {
		p, err := NewPlan(testCase)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		p.Explain()
	}
}

func Test_Plan_Next(t *testing.T) {
	testCases := [][]interface{}{
		{"15 7 * * 1-5", "2022/01/15 12:00"},     // get up for work
		{"50 21 * * *", "2020/01/18 20:59"},      // time for sleep
		{"0 8-19/2 * * 1-5", "2022/01/15 12:00"}, // stand up and drink water
		{"10 10 * 7,8 *", "2022/01/15 12:00"},    // finish homework in summer holiday
		{"0 0 31 4 *", "2022/01/15 12:00"},       // bad expression
		{"0 0 30 2 *", "2022/01/15 12:00"},       // bad expression
		{"0 0 29 2 *", "2022/02/28 23:00"},
	}

	for _, testCase := range testCases {
		p, err := NewPlan(testCase[0].(string))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		t := p.Next(time.Now().Add(time.Second * time.Duration(1)))
		fmt.Println(t)
	}
}

func Test_calcDatesByDOM(t *testing.T) {
	testCases := [][]interface{}{
		{2021, 8, uint64(1<<1 | 1<<7)},
		{2021, 12, uint64(1<<3 | 1<<5)},
		{2022, 1, uint64(1<<2 | 1<<4)},
		{2022, 2, uint64(1<<3 | 1<<5)},
	}
	for _, testCase := range testCases {
		result := dowToDOM(uint(testCase[0].(int)), uint(testCase[1].(int)), testCase[2].(uint64))
		fmt.Println(biu.ToBinaryString(result))
	}
}

func Test_a(t *testing.T) {
	a := ^uint64(0)
	fmt.Println(biu.ToBinaryString(a))

	fmt.Println(time.Date(2001, 2, 0, 0, 0, 0, 0, time.UTC))
}
