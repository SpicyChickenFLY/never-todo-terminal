package cron

import (
	"fmt"
	"time"
)

func mergeDOWToDOM(year, month uint, dowVal, domVal uint64) uint64 {
	isDOMWithStar, isDOWWithStar := dowVal&maskOptStar != 0, domVal&maskOptStar != 0
	if !isDOWWithStar {
		domValByDOW := dowToDOM(year, month, dowVal)
		if isDOMWithStar {
			return domValByDOW
		}
		return domValByDOW & domVal
	}
	return domVal
}

func dowToDOM(year, month uint, dowVal uint64) uint64 {
	result := uint64(0)
	t, err := time.Parse("2006/01/02", fmt.Sprintf("%4d/%02d/01", year, month))
	if err != nil {
		// TODO: handle error
	}
	fmt.Println(t.Format("2006/01/02"))
	offset := 1 - int(t.Weekday()) - 7
	for dow := fieldConfs[dowType].min; dow <= uint(daysIn(year, month)) && 1<<dow&dowVal != 0; dow++ {
		for i := int(dow) + offset; i <= 31; i += 7 {
			if i <= 0 {
				continue
			}
			result |= 1 << i
		}
	}
	return result
}

// daysIn returns the number of days in a month for a given year.
func daysIn(year, month uint) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(int(year), time.Month(month)+1, 0, 0, 0, 0, 0, nil).Day()
}
