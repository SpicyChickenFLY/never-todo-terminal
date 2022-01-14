package cron

import (
	"fmt"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
)

type fieldConf struct {
	min, max uint
	units    []string
}

var fieldConfs = []*fieldConf{
	{0, 60, []string{"second"}},
	{0, 60, []string{"minute"}},
	{0, 23, []string{"hour"}},
	{1, 31, []string{"day"}},
	{1, 12, monthName},
	{1, 7, dayName},
	{1, 9999, []string{"year"}},
}

type field struct {
	expr string
	conf *fieldConf
}

func newField(cronExpr string, conf *fieldConf) field {
	return field{cronExpr, conf}
}

func (f *field) getValues(min, max uint) (values uint, err error) {
	parts := strings.Split(f.expr, ",")
	for _, partExpr := range parts {
		var start, end, step uint = min, max, 1
		rangeAndStep := strings.Split(partExpr, "/")
		lowAndHigh := strings.Split(rangeAndStep[0], "-")
		switch len(lowAndHigh) {
		}

		// assemable result
		for i := start; i <= end; i += step {
			values |= 1 << i // set values
		}
	}
}

// explainField return field and error
func (f *field) explainField(needAbbrev bool) (string, error) {
	parts := strings.Split(f.expr, ",")
	explanations := make([]string, 0)
	for _, partExpr := range parts {
		explaination, err := explainRangeAndStep(partExpr, f.conf, needAbbrev)
		if err != nil {
			return "", err
		}
		explanations = append(explanations, explaination)
	}
	return strings.Join(explanations, " and "), nil
}

// explainRangeAndStep returns the bits indicated by the given expression:
//   number | number "-" number [ "/" number ]
// or error parsing range.
func explainRangeAndStep(expr string, fc *fieldConf, needAbbrev bool) (expl string, err error) {
	var start, end, step uint = fc.min, fc.max, 1
	rangeAndStep := strings.Split(expr, "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")

	// parse range
	switch len(lowAndHigh) {
	case 1:
		if lowAndHigh[0] != "*" {
			if start, err = utils.MustParseUint(lowAndHigh[0]); err != nil {
				return "", err
			}
			end = start
		}
	case 2:
		if start, err = utils.MustParseUint(lowAndHigh[0]); err != nil {
			return "", err
		}
		if end, err = utils.MustParseUint(lowAndHigh[1]); err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("too many hyphens: %s", expr)
	}

	if start < fc.min || end > fc.max {
		return "", fmt.Errorf("range(%d-%d) is out of bound(%d-%d): %s", start, end, fc.min, fc.max, expr)
	}
	if start > end {
		start, end = end, start
	}

	// parse step
	switch len(rangeAndStep) {
	case 1:
	case 2:
		step, err = utils.MustParseUint(rangeAndStep[1])
		if err != nil {
			return "", err
		}
		if step <= 0 {
			return "", fmt.Errorf("step must be positive: %s", expr)
		}
	default:
		return "", fmt.Errorf("too many slashes: %s", expr)
	}

	if len(lowAndHigh) == 1 {
		if lowAndHigh[0] != "*" {
			if len(fc.units) == 1 {
				expl = fmt.Sprintf("every %s %s", utils.GetOrdinalFormat(start), processUnit(fc.units[0], needAbbrev))
			} else {
				expl = fmt.Sprintf("every %s", processUnit(fc.units[start], needAbbrev))
			}
		} else {
			if step == 1 {
				expl = fmt.Sprintf("every %s", processUnit(fc.units[0], needAbbrev))
			} else {
				expl = fmt.Sprintf("every %d %s", step, processUnit(fc.units[0], needAbbrev))
			}
		}
	} else {
		if len(fc.units) == 1 {
			if step > 1 {
				expl = fmt.Sprintf("every %d %s ", step, processUnit(fc.units[0], needAbbrev))
			}
			expl += fmt.Sprintf("from every %s %s to %s %s",
				utils.GetOrdinalFormat(start), fc.units[0], utils.GetOrdinalFormat(end), processUnit(fc.units[0], needAbbrev))
		} else {
			explanations := make([]string, 0)
			for i := start; i <= end; i += step {
				explanations = append(explanations, processUnit(fc.units[i], needAbbrev))
			}
			expl = "every " + strings.Join(explanations, ", ")
		}
	}

	return expl, nil
}

func processUnit(unit string, needAbbrev bool) string {
	if needAbbrev {
		return unit[0:3]
	}
	return unit
}
