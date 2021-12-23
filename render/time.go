package render

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/utils/colorful"
)

const (
	second = 1
	minute = second * 60
	hour   = minute * 60
	day    = hour * 24
	week   = day * 7
	month  = day * 30
	year   = day * 365
)

// EstimateTime between two time
func EstimateTime(t1, t2 time.Time, needAbbrev bool) string {
	result := ""
	d := int(t1.Sub(t2).Seconds())
	if d < 0 {
		result = colorful.GetStartMark("default", "default", "red") + "(-"
		d *= -1
	} else {
		result = colorful.GetStartMark("default", "default", "green") + "("
	}
	if needAbbrev {
		switch {
		case d/year > 0:
			result += fmt.Sprintf("%dy", d/year)
		case d/month > 0:
			result += fmt.Sprintf("%dmon", d/month)
		case d/week > 0:
			result += fmt.Sprintf("%dw", d/week)
		case d/day > 0:
			result += fmt.Sprintf("%dd", d/day)
		case d/hour > 0:
			result += fmt.Sprintf("%dh", d/hour)
		case d/minute > 0:
			result += fmt.Sprintf("%dmin", d/minute)
		default:
			result = "now"
		}
	} else {
		switch {
		case d/year > 0:
			result += fmt.Sprintf("%dyear", d/year)
		case d/month > 0:
			result += fmt.Sprintf("%dmonth", d/month)
		case d/week > 0:
			result += fmt.Sprintf("%dweek", d/week)
		case d/day > 0:
			result += fmt.Sprintf("%dday", d/day)
		case d/hour > 0:
			result += fmt.Sprintf("%dhour", d/hour)
		case d/minute > 0:
			result += fmt.Sprintf("%dminute", d/minute)
		default:
			result = "now"
		}
	}
	result += ")" + colorful.GetEndMark()
	return result
}
