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
	if err := p.parseCronStr(cronStr, true); err != nil {
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
	return p.searchNextMatchDay(lastSchedule.Add(time.Second * time.Duration(1)))
}

func (p *Plan) searchNextMatchDay(lastSchedule time.Time) time.Time {
	lastYear := uint(lastSchedule.Year())
	lastMon := uint(lastSchedule.Month())
	lastDay := uint(lastSchedule.Day())
	year, mon, day := lastYear, lastMon, lastDay
	for ; ; year++ {
		for ; mon <= fcs[monthType].max; mon++ {
			if 1<<mon&p.fields[monthType].getValues(fcs[monthType].max) == 0 {
				day = uint(1)
				continue
			}
			max := daysIn(year, mon)
			domVal := p.fields[domType].getValues(uint(max))
			dowVal := p.fields[dowType].getValues(fcs[dowType].max)
			dayVal := mergeDOWToDOM(year, mon, dowVal, domVal)
			for ; day <= uint(max); day++ {
				if 1<<day&dayVal == 0 {
					continue
				}
				if lastDay == day {
					if h, min, s, ok := p.searchNextMatchTime(lastSchedule); ok {
						return time.Date(
							int(year), time.Month(mon), int(day),
							int(h), int(min), int(s),
							0, time.UTC)
					}
				} else {
					h, min, s := p.findFirstMatchTime()
					return time.Date(
						int(year), time.Month(mon), int(day),
						int(h), int(min), int(s),
						0, time.UTC)
				}
			}
			day = fcs[domType].min
		}
		mon = fcs[monthType].min
		day = fcs[domType].min
	}
}

func (p *Plan) searchNextMatchTime(lastSchedule time.Time) (h, min, s uint, ok bool) {
	lastHour := uint(lastSchedule.Hour())
	lastMin := uint(lastSchedule.Minute())
	lastSec := uint(lastSchedule.Second())
	for h = lastHour; h <= fcs[hourType].max; h++ {
		if 1<<h&p.fields[hourType].getValues(fcs[hourType].max) == 0 {
			continue
		}
		for min = lastMin; min <= fcs[minType].max; min++ {
			if 1<<min&p.fields[minType].getValues(fcs[minType].max) == 0 {
				continue
			}
			for s = lastSec; s <= fcs[secType].max; s++ {
				if 1<<s&p.fields[secType].getValues(fcs[secType].max) == 0 {
					continue
				}
				return
			}
		}
	}
	return 0, 0, 0, false
}

func (p *Plan) findFirstMatchTime() (h, min, s uint) {
	hourVal := p.fields[hourType].getValues(fcs[hourType].max)
	for h = 0; hourVal&1 == 0; hourVal >>= 1 {
		h++
	}
	minVal := p.fields[minType].getValues(fcs[minType].max)
	for min = 0; minVal&1 == 0; minVal >>= 1 {
		min++
	}
	secVal := p.fields[secType].getValues(fcs[secType].max)
	for s = 0; secVal&1 == 0; secVal >>= 1 {
		s++
	}
	return
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
		p.fields[i], err = newField(fieldExprs[i], fcs[i], needAbbrev)
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
