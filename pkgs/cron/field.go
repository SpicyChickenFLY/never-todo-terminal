package cron

import (
	"fmt"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
)

// 用uint map来代表结果值
// 难点在于特殊标识符L
// 在范围值中使用L可以，取值通过限制最大值即可
// 单个L表示最后一天，限制最大值后会丢失结果
// 所以还是需要设置一个特殊标志位
type field struct {
	values      uint64
	conf        *fieldConf
	explanation string
}

func newField(expr string, conf *fieldConf, needAbbrev bool) (*field, error) {
	f := &field{conf: conf}
	err := f.parseField(expr, needAbbrev)
	return f, err
}

func (f *field) getValues(max uint) uint64 {
	offset := maskLen - 1 - max
	mask := (maskAll<<offset)>>offset | maskOpt
	if f.values&maskOptL != 0 {
		f.values |= 1 << max
	}
	return f.values & mask
}

func (f *field) getExplanation() string {
	return f.explanation
}

// explainField return field and error
func (f *field) parseField(expr string, needAbbrev bool) error {
	parts := strings.Split(expr, ",")
	explanations := make([]string, 0)
	values := uint64(0)
	for _, partExpr := range parts {
		explaination, value, err := parseRangeAndStep(partExpr, f.conf, needAbbrev)
		if err != nil {
			return err
		}
		explanations = append(explanations, explaination)
		values |= value
	}
	f.explanation = strings.Join(explanations, " and ")
	f.values = values
	return nil
}

// parseRangeAndStep returns the bits indicated by the given expression:
//   number | number "-" number [ "/" number ]
// or error parsing range.
func parseRangeAndStep(expr string, fc *fieldConf, needAbbrev bool) (expl string, value uint64, err error) {
	var start, end, step uint = fc.min, fc.max, 1
	rangeAndStep := strings.Split(expr, "/")
	lowAndHigh := strings.Split(rangeAndStep[0], "-")

	// parse step
	switch len(rangeAndStep) {
	case 1:
	case 2:
		if step, err = utils.MustParseUint(rangeAndStep[1]); err != nil {
			return "", 0, err
		}
		if step <= 0 {
			return "", 0, fmt.Errorf("step must be positive: %s", expr)
		}
	default:
		return "", 0, fmt.Errorf("too many slashes: %s", expr)
	}

	// parse range
	switch len(lowAndHigh) {
	case 1:
		if lowAndHigh[0] != "*" && lowAndHigh[0] != "L" {
			if start, err = utils.MustParseUint(lowAndHigh[0]); err != nil {
				return "", 0, err
			}
		}
	case 2:
		if start, err = utils.MustParseUint(lowAndHigh[0]); err != nil {
			return "", 0, err
		}
		if lowAndHigh[1] != "L" {
			if end, err = utils.MustParseUint(lowAndHigh[1]); err != nil {
				return "", 0, err
			}
		}
	default:
		return "", 0, fmt.Errorf("too many hyphens: %s", expr)
	}
	if start < fc.min || end > fc.max {
		return "", 0, fmt.Errorf("range(%d-%d) is out of bound(%d-%d): %s", start, end, fc.min, fc.max, expr)
	}
	if start > end {
		return "", 0, fmt.Errorf("max less than min(%d-%d): %s", start, end, expr)
	}

	if len(lowAndHigh) == 1 {
		switch {
		case lowAndHigh[0] == "L":
			expl = fmt.Sprintf("every last %s", processUnit(fc.units[0], needAbbrev))
			value = maskOptL
			return
		case lowAndHigh[0] == "*" && step == 1:
			expl = fmt.Sprintf("every %s", processUnit(fc.units[0], needAbbrev))
		case lowAndHigh[0] == "*" && step != 1:
			expl = fmt.Sprintf("every %d %s", step, processUnit(fc.units[0], needAbbrev))
		default:
			expl = fmt.Sprintf("every %s", valueToExpr(start, fc.units, needAbbrev))
			value = 1 << start
			return
		}
	} else { // range
		switch {
		case len(fc.units) != 1:
			explanations := make([]string, 0)
			for i := start; i <= end; i += step {
				explanations = append(explanations, processUnit(fc.units[i], needAbbrev))
			}
			expl = "every " + strings.Join(explanations, ", ")
		case step > 1:
			expl = fmt.Sprintf("every %d %s ", step, processUnit(fc.units[0], needAbbrev))
		case lowAndHigh[1] == "L":
			expl += fmt.Sprintf("from every %s to the last %s",
				valueToExpr(start, []string{fc.units[0]}, needAbbrev), processUnit(fc.units[0], needAbbrev))
		default:
			expl += fmt.Sprintf("from every %s to %s",
				valueToExpr(start, []string{fc.units[0]}, needAbbrev),
				valueToExpr(end, []string{fc.units[0]}, needAbbrev))
		}
	}
	for i := start; i <= end; i += step {
		value |= 1 << i
	}
	return
}

func assembleExplanation() string {
	return ""
}

func valueToExpr(value uint, units []string, needAbbrev bool) string {
	if len(units) == 1 {
		return fmt.Sprintf("%s %s", utils.GetOrdinalFormat(value), processUnit(units[0], needAbbrev))
	}
	return processUnit(units[value], needAbbrev)
}

func processUnit(unit string, needAbbrev bool) string {
	if needAbbrev {
		return unit[0:3]
	}
	return unit
}
