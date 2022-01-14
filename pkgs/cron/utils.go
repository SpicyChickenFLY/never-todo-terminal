package cron

import (
	"fmt"
	"sort"
	"time"
)

func convertDOWToDOM(year, month uint, dowList []uint) []uint {
	result := []uint{}
	t, err := time.Parse("2006/01/02", fmt.Sprintf("%4d/%02d/01", year, month))
	if err != nil {

	}
	fmt.Println(t.Format("2006/01/02"))
	offset := 1 - int(t.Weekday()) - 7
	for _, dom := range dowList {
		for i := int(dom) + offset; i <= 31; i += 7 {
			if i <= 0 {
				continue
			}
			result = append(result, uint(i))
		}
	}
	sort.SliceStable(result, func(i, j int) bool { return int(result[i]) < int(result[j]) })
	return result
}

func mergeDOMWithDOW(domList, dowList []uint) []uint {
	result := []uint{}
	j := 0
	for i := 0; i < len(domList) && j < len(dowList); i++ {
		for ; j < len(dowList); j++ {
			if dowList[j] > domList[i] {
				break
			} else if dowList[j] == domList[i] {
				result = append(result, dowList[j])
			}
		}
	}
	return result
}

func searchNextMatch(year, month int, lastSchedule time.Time) (time.Time, bool) {
	// TODO: Calc Next Matched day&time in this month
	return time.Time{}, false
}
