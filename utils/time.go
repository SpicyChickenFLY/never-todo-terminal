package utils

import (
	"fmt"
	"sort"
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

var monthName = []string{
	"month",
	"January", "February", "March", "April",
	"May", "June", "July", "August",
	"September", "October", "November", "December",
}

var dayName = []string{
	"day", "Monday", "Tuesday", "Wednesday",
	"Thursday", "Friday", "Saturday", "Sunday",
}

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

type field struct {
	values      []uint
	explanation string
}
type fieldConf struct {
	min, max uint
	units    []string
}

var fieldConfs = []fieldConf{
	{0, 60, []string{"second"}},
	{0, 60, []string{"minute"}},
	{0, 23, []string{"hour"}},
	{1, 31, []string{"day"}},
	{1, 12, monthName},
	{1, 7, dayName},
}

func calcDatesByDOW(year, month uint, dowList []uint) []uint {
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

// CalcNextSchedule after last schedule by your plan
func CalcNextSchedule(cronStr string, lastSchedule time.Time) time.Time {
	result := time.Time{}
	// TODO: CalcNextSchedule
	fields, err := parseCronString(cronStr, true)
	if err != nil {
		return result
	}

	carry := false
	// calculate available month
	lsMonth := uint(lastSchedule.Month())
	nextMonth :=
	for _, month := range fields[monthIdx].values {
		if lsMonth < month {
			carry = true
		} else if lsMonth == month {
			break
		}
	}
	// calculate time
	if carry == true {

	}
	return result
}

// ExplainSchedule bu return a string
func ExplainSchedule(cronStr string, needAbbrev bool) string {
	fields, err := parseCronString(cronStr, needAbbrev)
	if err != nil {
		fmt.Printf("Invalid Crontab string: %s\n", err.Error())
		return cronStr
	}
	fmt.Printf("this task loop AT %s FOR %s ON %s and %s IN %s\n",
		fields[minIdx].explanation,
		fields[hourIdx].explanation,
		fields[domIdx].explanation,
		fields[dowIdx].explanation,
		fields[monthIdx].explanation,
	)
	return cronStr
}

//parseCronString return its explaination and calculate next schedule
func parseCronString(cronStr string, needAbbrev bool) ([]field, error) {
	if len(cronStr) == 0 {
		return nil, fmt.Errorf("invalid format")
	}

	// Split on whitespace.
	fieldExprs := strings.Fields(cronStr)
	switch len(fieldExprs) {
	case 5:
		fieldExprs = append([]string{"0"}, fieldExprs...) // fill second field
		fieldExprs = append(fieldExprs, "*")              // fill year field
	default:
		return nil, fmt.Errorf("too few fields")

	}

	fields := make([]field, fieldCount)
	var err error
	for i := minIdx; i <= dowIdx; i++ {
		fields[i], err = parseField(fieldExprs[i], fieldConfs[i].min, fieldConfs[i].max, fieldConfs[i].units, needAbbrev)
		if err != nil {
			return nil, err
		}
	}

	return fields, nil
}

func parseField(expr string, min, max uint, units []string, needAbbrev bool) (field, error) {
	f := field{make([]uint, 0), ""}
	parts := strings.Split(expr, ",")
	explanations := make([]string, 0)
	for _, partExpr := range parts {
		value, explaination, err := parseRangeAndStep(partExpr, min, max, units, needAbbrev)
		if err != nil {
			return field{}, err
		}
		f.values = append(f.values, value...)
		explanations = append(explanations, explaination)
	}
	f.explanation = strings.Join(explanations, " and ")
	return f, nil
}

// parseRangeAndStep returns the bits indicated by the given expression:
//   number | number "-" number [ "/" number ]
// or error parsing range.
func parseRangeAndStep(expr string, min, max uint, units []string, needAbbrev bool) ([]uint, string, error) {
	var start, end, step uint = min, max, 1
	var err error
	rangeAndStep := strings.Split(expr, "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")

	values := make([]uint, 0)
	explanation := ""

	// parse range
	switch len(lowAndHigh) {
	case 1:
		if lowAndHigh[0] != "*" {
			if start, err = mustParseInt(lowAndHigh[0]); err != nil {
				return nil, "", err
			}
			end = start
		}
	case 2:
		if start, err = mustParseInt(lowAndHigh[0]); err != nil {
			return nil, "", err
		}
		if end, err = mustParseInt(lowAndHigh[1]); err != nil {
			return nil, "", err
		}
	default:
		return nil, "", fmt.Errorf("too many hyphens: %s", expr)
	}

	if start < min || end > max {
		return nil, "", fmt.Errorf("range(%d-%d) is out of bound(%d-%d): %s", start, end, min, max, expr)
	}
	if start > end {
		start, end = end, start
	}

	// parse step
	switch len(rangeAndStep) {
	case 1:
	case 2:
		step, err = mustParseInt(rangeAndStep[1])
		if err != nil {
			return nil, "", err
		}
		if step <= 0 {
			return nil, "", fmt.Errorf("step must be positive: %s", expr)
		}
	default:
		return nil, "", fmt.Errorf("too many slashes: %s", expr)
	}

	if start != end && end-start < step {
		return nil, "", fmt.Errorf("step(%d) out ot of range(%d-%d): %s", step, start, end, expr)
	}

	// assemable result
	for i := start; i <= end; i += step {
		values = append(values, i)
	}

	if len(lowAndHigh) == 1 {
		if lowAndHigh[0] != "*" {
			if len(units) == 1 {
				explanation = fmt.Sprintf("every %s %s", GetOrdinalFormat(start), processUnit(units[0], needAbbrev))
			} else {
				explanation = fmt.Sprintf("every %s", processUnit(units[start], needAbbrev))
			}
		} else {
			if step == 1 {
				explanation = fmt.Sprintf("every %s", processUnit(units[0], needAbbrev))
			} else {
				explanation = fmt.Sprintf("every %d %s", step, processUnit(units[0], needAbbrev))
			}
		}
	} else {
		if len(units) == 1 {
			if step > 1 {
				explanation = fmt.Sprintf("every %d %s ", step, processUnit(units[0], needAbbrev))
			}
			explanation += fmt.Sprintf("from every %s %s to %s %s",
				GetOrdinalFormat(start), units[0], GetOrdinalFormat(end), processUnit(units[0], needAbbrev))
		} else {
			explanations := make([]string, 0)
			for i := start; i <= end; i += step {
				explanations = append(explanations, processUnit(units[i], needAbbrev))
			}
			explanation = "every " + strings.Join(explanations, ", ")
		}
	}

	return values, explanation, nil
}

func processUnit(unit string, needAbbrev bool) string {
	if needAbbrev {
		return unit[0:3]
	}
	return unit
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
