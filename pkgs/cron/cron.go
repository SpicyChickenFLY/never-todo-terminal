package cron

import (
	"fmt"
	"strings"
	"time"
)

// Plan contains loop plan
type Plan struct {
	expr    string
	fields  [fieldCount]*field
	isValid bool
}

// NewPlan return *Plan
func NewPlan(cronStr string) (*Plan, error) {
	p := &Plan{}
	if err := p.parseCronStr(cronStr, false); err != nil {
		return nil, err
	}
	return p, nil
}

// Explain return explanation
func (p *Plan) Explain() string {
	fmt.Printf("this task loop AT %s FOR %s ON %s and %s IN %s\n",
		p.fields[minType].getExplanation(),
		p.fields[hourType].getExplanation(),
		p.fields[domType].getExplanation(),
		p.fields[dowType].getExplanation(),
		p.fields[monthType].getExplanation(),
	)
	return p.expr
}

// GetExpr of plan
func (p *Plan) GetExpr() string {
	return p.expr
}

// Next due time after last schedule in this plan
// this function promise a valid return
func (p *Plan) Next(lastSchedule time.Time) time.Time {
	// TODO: Calc Next Schedule
	return p.searchNextMatchDay(lastSchedule)
}

func (p *Plan) searchNextMatchDay(lastSchedule time.Time) time.Time {
	lastYear := uint(lastSchedule.Year())
	lastMonth := uint(lastSchedule.Month())
	lastDay := uint(lastSchedule.Day())
	for year := lastYear; ; year++ {
		for month := lastMonth; month <= fieldConfs[monthType].max; month++ {
			if 1<<month&p.fields[monthType].getValues(fieldConfs[monthType].max) == 0 {
				continue
			}
			max := daysIn(year, month)
			domVal := p.fields[domType].getValues(uint(max))
			dowVal := p.fields[dowType].getValues(fieldConfs[dowType].max)
			dayVal := mergeDOWToDOM(year, month, dowVal, domVal)
			for day := lastDay; day <= uint(max); day++ {
				if 1<<day&dayVal == 0 {
					continue
				}
				result, ok := p.searchNextMatchTime(day, lastSchedule)
				if ok {
					return result
				}
			}
		}
	}
}

func (p *Plan) searchNextMatchTime(dom uint, lastSchedule time.Time) (time.Time, bool) {
	// TODO: Calc Next Matched day&time in this month
	return time.Time{}, false
}

//parseCronStr return its explaination and calculate next schedule
func (p *Plan) parseCronStr(cronStr string, needAbbrev bool) error {
	p.expr = cronStr
	if len(cronStr) == 0 {
		return fmt.Errorf("Expression is empty: %s", p.expr)
	}

	// Split on whitespace.
	fieldExprs := strings.Fields(cronStr)
	switch len(fieldExprs) {
	case 5:
		fieldExprs = append([]string{"0"}, fieldExprs...) // fill second field
		fieldExprs = append(fieldExprs, "*")              // fill year field
	default:
		return fmt.Errorf("Expression contains too few fields: %s", p.expr)
	}

	var err error
	for i := secType; i <= yearType; i++ {
		p.fields[i], err = newField(fieldExprs[i], fieldConfs[i], needAbbrev)
		if err != nil {
			return err
		}
	}
	// expression like * * 31 2,4,6,9,11 * is very dangerous for plan
	if p.fields[domType].values == 1<<31 &&
		p.fields[monthType].values&monthWith31Days == 0 {
		return fmt.Errorf("Expression like * * 31 2,4,6,9,11 * is invalid: %s", p.expr)
	}
	// also, * * 30 2 * is dangerous
	if p.fields[domType].values == 1<<30 &&
		p.fields[monthType].values == 1<<2 {
		return fmt.Errorf("Expression like * * 30 2 * is invalid: %s", p.expr)
	}
	return nil
}
