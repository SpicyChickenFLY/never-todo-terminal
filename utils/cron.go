package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

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
	values      uint
	expr        string
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

func searchNextMatch(year, month int) (time.Time, bool) {
	// TODO: Calc Next Matched day&time in this month
	return time.Time{}, false
}

// CalcNextSchedule after last schedule by your plan
func CalcNextSchedule(cronStr string, lastSchedule time.Time) time.Time {
	result := time.Time{}
	// TODO: CalcNextSchedule
	fields, err := parseCronStr(cronStr, true)
	if err != nil {
		return result
	}

	// calculate available month
	lastMonth := uint(lastSchedule.Month())
	lastYear := uint(lastSchedule.Year())
	for year := lastYear; ; year++ {
		for month := fields[monthIdx].values[0]; month >= lastMonth; month++ {
			result, ok := searchNextMatch(int(year), int(month))
			if ok {
				return result
			}
		}
	}
}

// Plan contains loop plan
type Plan struct {
	fields [fieldCount]field
}

// NewPlan return *Plan
func NewPlan(cronStr string) *Plan {
	p := &Plan{}
	if err := p.parseCronStr(cronStr, false); err != nil {
	}
	return p
}

func (p *Plan) checkValid() bool {
	// check field rang
	return false
}

// Explain return explanation
func (p *Plan) Explain() string {
	if p.checkValid() {
		fmt.Printf("this task loop AT %s FOR %s ON %s and %s IN %s\n",
			p.fields[minIdx].explanation,
			p.fields[hourIdx].explanation,
			p.fields[domIdx].explanation,
			p.fields[dowIdx].explanation,
			p.fields[monthIdx].explanation,
		)
	} else {
		fmt.Printf("Invalid Plan\n")
	}
	return fmt.Sprintf("%s %s %s %s %s %s %s",
		p.fields[secIdx].expr,
		p.fields[minIdx].expr,
		p.fields[hourIdx].expr,
		p.fields[domIdx].expr,
		p.fields[dowIdx].expr,
		p.fields[monthIdx].expr,
		p.fields[yearIdx].expr,
	)
}

//parseCronStr return its explaination and calculate next schedule
func (p *Plan) parseCronStr(cronStr string, needAbbrev bool) error {
	if len(cronStr) == 0 {
		return fmt.Errorf("invalid format")
	}

	// Split on whitespace.
	fieldExprs := strings.Fields(cronStr)
	switch len(fieldExprs) {
	case 5:
		fieldExprs = append([]string{"0"}, fieldExprs...) // fill second field
		fieldExprs = append(fieldExprs, "*")              // fill year field
	default:
		return fmt.Errorf("too few fields")

	}

	var err error
	for i := secIdx; i <= yearIdx; i++ {
		p.fields[i], err = parseField(fieldExprs[i], fieldConfs[i].min, fieldConfs[i].max, fieldConfs[i].units, needAbbrev)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseField(expr string, min, max uint, units []string, needAbbrev bool) (field, error) {
	f := field{uint(0), "", ""}
	f.expr = expr
	parts := strings.Split(expr, ",")
	explanations := make([]string, 0)
	for _, partExpr := range parts {
		value, explaination, err := parseRangeAndStep(partExpr, min, max, units, needAbbrev)
		if err != nil {
			return field{}, err
		}
		f.values = value
		explanations = append(explanations, explaination)
	}
	f.explanation = strings.Join(explanations, " and ")
	return f, nil
}

// parseRangeAndStep returns the bits indicated by the given expression:
//   number | number "-" number [ "/" number ]
// or error parsing range.
func parseRangeAndStep(expr string, min, max uint, units []string, needAbbrev bool) (uint, string, error) {
	var start, end, step uint = min, max, 1
	var err error
	rangeAndStep := strings.Split(expr, "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")

	explanation := ""

	// parse range
	switch len(lowAndHigh) {
	case 1:
		if lowAndHigh[0] != "*" {
			if start, err = mustParseUint(lowAndHigh[0]); err != nil {
				return 0, "", err
			}
			end = start
		}
	case 2:
		if start, err = mustParseUint(lowAndHigh[0]); err != nil {
			return 0, "", err
		}
		if end, err = mustParseUint(lowAndHigh[1]); err != nil {
			return 0, "", err
		}
	default:
		return 0, "", fmt.Errorf("too many hyphens: %s", expr)
	}

	if start < min || end > max {
		return 0, "", fmt.Errorf("range(%d-%d) is out of bound(%d-%d): %s", start, end, min, max, expr)
	}
	if start > end {
		start, end = end, start
	}

	// parse step
	switch len(rangeAndStep) {
	case 1:
	case 2:
		step, err = mustParseUint(rangeAndStep[1])
		if err != nil {
			return 0, "", err
		}
		if step <= 0 {
			return 0, "", fmt.Errorf("step must be positive: %s", expr)
		}
	default:
		return 0, "", fmt.Errorf("too many slashes: %s", expr)
	}

	if start != end && end-start < step {
		return 0, "", fmt.Errorf("step(%d) out ot of range(%d-%d): %s", step, start, end, expr)
	}

	var values uint
	// assemable result

	for value << start; i <= end; i += step {
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
