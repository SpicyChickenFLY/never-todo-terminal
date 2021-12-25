package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/utils/colorful"
)

const (
	minPerSec   = 60
	hourPerSec  = minPerSec * 60
	dayPerSec   = hourPerSec * 24
	weekPerSec  = dayPerSec * 7
	monthPerSec = dayPerSec * 30
	yearPerSec  = dayPerSec * 365
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
		case d/yearPerSec > 0:
			result += fmt.Sprintf("%dy", d/yearPerSec)
		case d/monthPerSec > 0:
			result += fmt.Sprintf("%dmon", d/monthPerSec)
		case d/weekPerSec > 0:
			result += fmt.Sprintf("%dw", d/weekPerSec)
		case d/dayPerSec > 0:
			result += fmt.Sprintf("%dd", d/dayPerSec)
		case d/hourPerSec > 0:
			result += fmt.Sprintf("%dh", d/hourPerSec)
		case d/minPerSec > 0:
			result += fmt.Sprintf("%dmin", d/minPerSec)
		default:
			result = colorful.GetStartMark("default", "default", "yellow") + "(now"
		}
	} else {
		switch {
		case d/yearPerSec > 0:
			result += fmt.Sprintf("%dyear", d/yearPerSec)
		case d/monthPerSec > 0:
			result += fmt.Sprintf("%dmonth", d/monthPerSec)
		case d/weekPerSec > 0:
			result += fmt.Sprintf("%dweek", d/weekPerSec)
		case d/dayPerSec > 0:
			result += fmt.Sprintf("%dday", d/dayPerSec)
		case d/hourPerSec > 0:
			result += fmt.Sprintf("%dhour", d/hourPerSec)
		case d/minPerSec > 0:
			result += fmt.Sprintf("%dminute", d/minPerSec)
		default:
			result = colorful.GetStartMark("default", "default", "yellow") + "(now"
		}
	}
	result += ")" + colorful.GetEndMark()
	return result
}

var monthName = []string{
	"Error",
	"January", "February", "March", "April",
	"May", "June", "July", "August",
	"September", "October", "November", "December",
}

var dayName = []string{
	"Error", "Monday", "Tuesday", "Wednesday",
	"Thursday", "Friday", "Saturday", "Sunday",
}

// field index
const (
	secIdx = iota
	minIdx
	hourIdx
	domIdx
	monthIdx
	dowIdx
	yearIdx
	fieldCount
)

type Field struct {
	values   []uint
	min, max uint
	units    []string
}

func explain() string {
	return ""
}

// CalcNextSchedule after last schedule by your plan
func CalcNextSchedule(cronStr string, lastSchedule time.Time) time.Time {
	return time.Now()
}

// ExplainSchedule bu return a string
func ExplainSchedule(cronStr string) string {
	return ""
}

//ParseCronString return its explaination and calculate next schedule
func ParseCronString(cronStr string) ([]Field, error) {
	if len(cronStr) == 0 {
		return nil, fmt.Errorf("invalid format")
	}

	// Split on whitespace.
	fieldExprs := strings.Fields(cronStr)
	if len(fieldExprs) == 5 {
		fieldExprs = append([]string{"0"}, fieldExprs...) // fill second field
		fieldExprs = append(fieldExprs, "*")              // fill year field
	}

	fields := make([]Field, fieldCount)
	fields[minIdx] = parseField(fieldExprs[minIdx], 0, 60, []string{"minute"})
	fields[hourIdx] = parseField(fieldExprs[hourIdx], 0, 60, []string{"hour"})
	fields[domIdx] = parseField(fieldExprs[domIdx], 0, 60, []string{"day"})
	fields[monthIdx] = parseField(fieldExprs[monthIdx], 0, 60, monthName)
	fields[dowIdx] = parseField(fieldExprs[dowIdx], 0, 60, dayName)

	return fields, nil
}

func parseField(expr string, min, max uint, units []string) Field {
	field := Field{min: min, max: max, units: units}
	// slices := strings.Split(expr, ",")
	// for _, slice := range slices {
	// 	part :=
	// }
	return field
}

// getField returns an Int with the bits set representing all of the times that
// the field represents or error parsing field value.  A "field" is a comma-separated
// list of "ranges".
func getField(field string, min, max uint) (uint64, error) {
	var bits uint64
	ranges := strings.FieldsFunc(field, func(r rune) bool { return r == ',' })
	for _, expr := range ranges {
		bit, err := getRange(expr, min, max)
		if err != nil {
			return bits, err
		}
		bits |= bit
	}
	return bits, nil
}

// getRange returns the bits indicated by the given expression:
//   number | number "-" number [ "/" number ]
// or error parsing range.
func getRange(expr string, min, max uint) (uint64, error) {
	var start, end, step uint
	var err error
	rangeAndStep := strings.Split(expr, "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")
	singleDigit := len(lowAndHigh) == 1

	var extra uint64
	if lowAndHigh[0] == "*" || lowAndHigh[0] == "?" {
		start = min
		end = max
		extra = starBit
	} else {
		start, err = mustParseInt(lowAndHigh[1])
		if err != nil {
			return 0, err
		}
		switch len(lowAndHigh) {
		case 1:
			end = start
		case 2:
			end, err = mustParseInt(lowAndHigh[2])
			if err != nil {
				return 0, err
			}
		default:
			return 0, fmt.Errorf("too many hyphens: %s", expr)
		}
	}

	switch len(rangeAndStep) {
	case 1:
		step = 1
	case 2:
		step, err = mustParseInt(rangeAndStep[1])
		if err != nil {
			return 0, err
		}

		// Special handling: "N/step" means "N-max/step".
		if singleDigit {
			end = max
		}
		if step > 1 {
			extra = 0
		}
	default:
		return 0, fmt.Errorf("too many slashes: %s", expr)
	}

	if start < min {
		return 0, fmt.Errorf("beginning of range (%d) below minimum (%d): %s", start, min, expr)
	}
	if end > max {
		return 0, fmt.Errorf("end of range (%d) above maximum (%d): %s", end, max, expr)
	}
	if start > end {
		return 0, fmt.Errorf("beginning of range (%d) beyond end of range (%d): %s", start, end, expr)
	}
	if step == 0 {
		return 0, fmt.Errorf("step of range should be a positive number: %s", expr)
	}

	return getBits(start, end, step) | extra, nil
}

// mustParseInt parses the given expression as an int or returns an error.
func mustParseInt(expr string) (uint, error) {
	num, err := strconv.Atoi(expr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int from %s: %s", expr, err)
	}
	if num < 0 {
		return 0, fmt.Errorf("negative number (%d) not allowed: %s", num, expr)
	}

	return uint(num), nil
}

// getBits sets all bits in the range [min, max], modulo the given step size.
func getBits(min, max, step uint) uint64 {
	var bits uint64

	// If step is 1, use shifts.
	if step == 1 {
		return ^(math.MaxUint64 << (max + 1)) & (math.MaxUint64 << min)
	}

	// Else, use a simple loop.
	for i := min; i <= max; i += step {
		bits |= 1 << i
	}
	return bits
}

// all returns all bits within the given bounds.  (plus the star bit)
func all(min, max uint) uint64 {
	return getBits(min, max, 1) | starBit
}
