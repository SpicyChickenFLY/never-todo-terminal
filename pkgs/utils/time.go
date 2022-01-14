package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/colorful"
)

const (
	minPerSec   = 60
	hourPerSec  = minPerSec * 60
	dayPerSec   = hourPerSec * 24
	weekPerSec  = dayPerSec * 7
	monthPerSec = dayPerSec * 30
	yearPerSec  = dayPerSec * 365
)

// GetOrdinalFormat transfer number to ordinary
func GetOrdinalFormat(num uint) string {
	if num > 10 && num < 20 {
		return fmt.Sprint(num, "th")
	}
	tail := num % 10
	switch tail {
	case 1:
		return fmt.Sprint(num, "st")
	case 2:
		return fmt.Sprint(num, "nd")
	case 3:
		return fmt.Sprint(num, "rd")
	default:
		return fmt.Sprint(num, "th")
	}
}

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
		case d/yearPerSec > 0:
			result += fmt.Sprintf("%.1fy", float64(d)/float64(yearPerSec))
		case d/monthPerSec > 0:
			result += fmt.Sprintf("%.1fmon", float64(d)/float64(monthPerSec))
		case d/weekPerSec > 0:
			result += fmt.Sprintf("%.1fw", float64(d)/float64(weekPerSec))
		case d/dayPerSec > 0:
			result += fmt.Sprintf("%.1fd", float64(d)/float64(dayPerSec))
		case d/hourPerSec > 0:
			result += fmt.Sprintf("%.1fh", float64(d)/float64(hourPerSec))
		case d/minPerSec > 0:
			result += fmt.Sprintf("%.1fmin", float64(d)/float64(minPerSec))
		default:
			result = colorful.GetStartMark("default", "default", "yellow") + "(now"
		}
	} else {
		switch {
		case d/yearPerSec > 0:
			result += fmt.Sprintf("%.1fyear", float64(d)/float64(yearPerSec))
		case d/monthPerSec > 0:
			result += fmt.Sprintf("%.1fmonth", float64(d)/float64(monthPerSec))
		case d/weekPerSec > 0:
			result += fmt.Sprintf("%.1fweek", float64(d)/float64(weekPerSec))
		case d/dayPerSec > 0:
			result += fmt.Sprintf("%.1fday", float64(d)/float64(dayPerSec))
		case d/hourPerSec > 0:
			result += fmt.Sprintf("%.1fhour", float64(d)/float64(hourPerSec))
		case d/minPerSec > 0:
			result += fmt.Sprintf("%.1fminute", float64(d)/float64(minPerSec))
		default:
			result = colorful.GetStartMark("default", "default", "yellow") + "(now"
		}
	}
	result += ")" + colorful.GetEndMark()
	return result
}

// Time format
const (
	TimeFormatDate = 0 + iota
	TimeFormatTime
	TimeFormatDateTime
)

func CompleteDateTime(str string, dtType int) string {
	now := time.Now()
	d, t := now.Format("2006/01/02"), "00:00:00"
	switch dtType {
	case TimeFormatDate:
		d = CompleteDate(str)
	case TimeFormatTime:
		t = CompleteTime(str)
	case TimeFormatDateTime:
		dateTimeParts := strings.Split(str, " ")
		d, t = CompleteDate(dateTimeParts[0]), CompleteTime(dateTimeParts[1])
	}
	return d + " " + t
}

func CompleteDate(d string) string {
	now := strings.Split(time.Now().Format("2006/01/02"), "/")
	var year, month, day string
	dateParts := strings.Split(d, "/")
	switch len(dateParts) {
	case 3: // year/month/day
		year, month, day = dateParts[0], dateParts[1], dateParts[2]
	case 2: // month/day
		year, month, day = now[0], dateParts[0], dateParts[1]
	case 1: // day
		year, month, day = now[0], now[1], dateParts[0]
	}
	if len(year) == 2 { // complete year
		year = now[0][:2] + year
	}
	if len(month) == 1 { // complete month
		month = "0" + month
	}
	if len(day) == 1 { // complete day
		day = "0" + day
	}
	return fmt.Sprintf("%s/%s/%s", year, month, day)
}

func CompleteTime(t string) string {
	now := strings.Split(time.Now().Format("15:04:05"), ":")
	var hour, minute, second string
	dateParts := strings.Split(t, ":")
	switch len(dateParts) {
	case 3: // hour/minute/second
		hour, minute, second = dateParts[0], dateParts[1], dateParts[2]
	case 2: // minute/second
		hour, minute, second = dateParts[0], dateParts[1], now[2]
	case 1: // second
		hour, minute, second = dateParts[0], now[1], now[2]
	}
	if len(hour) == 1 { // complete hour
		hour = "0" + hour
	}
	if len(minute) == 1 { // complete minute
		minute = "0" + minute
	}
	if len(second) == 1 { // complete second
		second = "0" + second
	}
	return fmt.Sprintf("%s:%s:%s", hour, minute, second)
}
