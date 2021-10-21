package ast

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

// Time format
const (
	TimeFormatDate = 0 + iota
	TimeFormatTime
	TimeFormatDateTime
)

// ============================
// ID Group
// ============================

// IDGroupNode is node include id
type IDGroupNode struct {
	ids []int
}

// NewIDGroupNode return IDGroup
func NewIDGroupNode(ids ...int) *IDGroupNode {
	switch len(ids) {
	case 1:
		return &IDGroupNode{ids: []int{ids[0]}}
	case 2:
		temp := []int{}
		if ids[0] < ids[1] {
			for i := ids[0]; i <= ids[1]; i++ {
				temp = append(temp, i)
			}
		} else {
			for i := ids[1]; i <= ids[0]; i++ {
				temp = append(temp, i)
			}
		}
		return &IDGroupNode{ids: temp}
	default:
		return &IDGroupNode{}
	}
}

// MergeIDNode merge with othen IDGroup
func (ign *IDGroupNode) MergeIDNode(ign1 *IDGroupNode) *IDGroupNode {
	ign.ids = append(ign.ids, ign1.ids...)
	ign.removeRepeatedIDs()
	return ign
}

func (ign *IDGroupNode) explain() string {
	fmt.Printf("\twith ID: %v\n", ign.ids)
	result := ""
	for _, id := range ign.ids {
		result += fmt.Sprint(" ", id)
	}
	return result[1:]
}

func (ign *IDGroupNode) sortID() {
	sort.Ints(ign.ids)
}

func (ign *IDGroupNode) removeRepeatedIDs() {
	if len(ign.ids) == 0 {
		return
	}
	ign.sortID()
	temp := []int{ign.ids[0]}
	for i := 1; i < len(ign.ids); i++ {
		if ign.ids[i] != ign.ids[i-1] {
			temp = append(temp, ign.ids[i])
		}
	}
	ign.ids = temp
}

// ============================
// Content Group
// ============================

// operator type
const (
	OPNone = iota + 0
	OPNOT
	OPAND
	OPOR
)

// ContentGroupNode is node include contents
type ContentGroupNode struct {
	content  string
	operator int
	operands []*ContentGroupNode
}

// NewContentGroupNode return ContentGroupNode
func NewContentGroupNode(
	content string, operator int, operands []*ContentGroupNode) *ContentGroupNode {
	fmt.Println("new content" + content)
	return &ContentGroupNode{content, operator, operands}
}

func (cgn *ContentGroupNode) explain() string {
	switch cgn.operator {
	case OPNone:
		fmt.Printf("include `%s`", cgn.content)
		return cgn.content
	case OPNOT:
		result := ""
		fmt.Print("while not (")
		if cgn.operands[0].operator != OPNone {
			result = fmt.Sprintf("!(%s)", cgn.operands[0].explain())
			fmt.Print(")")
			return result
		}
		result = fmt.Sprintf("!%s", cgn.operands[0].explain())
		fmt.Print(")")
		return result
	case OPAND:
		r1, r2 := "", ""
		fmt.Print("(")
		if cgn.operands[0].operator != OPOR {
			r1 = cgn.operands[0].explain()
		} else {
			r1 = fmt.Sprintf("(%s)", cgn.operands[0].explain())
		}
		fmt.Print(" and ")
		if cgn.operands[1].operator != OPOR {
			r2 = cgn.operands[1].explain()
		} else {
			r2 = fmt.Sprintf("(%s)", cgn.operands[1].explain())
		}
		fmt.Print(")")
		return fmt.Sprintf("%s and %s", r1, r2)
	case OPOR:
		fmt.Print("(")
		r1, r2 := "", ""
		if cgn.operands[0].operator != OPAND {
			r1 = cgn.operands[0].explain()
		} else {
			r1 = fmt.Sprintf("(%s)", cgn.operands[0].explain())
		}
		fmt.Print(" or ")
		if cgn.operands[1].operator != OPAND {
			r2 = cgn.operands[1].explain()
		} else {
			r2 = fmt.Sprintf("(%s)", cgn.operands[1].explain())
		}
		fmt.Print(")")
		return fmt.Sprintf("%s or %s", r1, r2)
	default:
		return ""
	}
}

func (cgn *ContentGroupNode) filter(tasks []model.Task) (result, negation []model.Task) {
	switch cgn.operator {
	case OPNone:
		for _, task := range tasks {
			if utils.ContainStr(task.Content, cgn.content) {
				result = append(result, task)
			} else {
				negation = append(negation, task)
			}
		}
		return result, negation
	case OPNOT:
		negation, result = cgn.operands[0].filter(tasks)
	case OPAND:
		leftResult, leftNegation := cgn.operands[0].filter(tasks)
		rightResult, _ := cgn.operands[1].filter(tasks)
		negation = leftNegation
		for _, lt := range leftResult {
			for _, rt := range rightResult {
				if lt.ID == rt.ID {
					result = append(result, lt)
					break
				}
			}
			if len(result) == 0 || result[len(result)-1].ID != lt.ID {
				negation = append(negation, lt)
			}
		}
	case OPOR:
		leftResult, leftNegation := cgn.operands[0].filter(tasks)
		_, rightNegation := cgn.operands[1].filter(tasks)
		result = leftResult
		for _, ln := range leftNegation {
			for _, rn := range rightNegation {
				if ln.ID == rn.ID {
					negation = append(negation, rn)
					break
				}
			}
			if len(negation) == 0 && negation[len(negation)-1].ID != ln.ID {
				result = append(result, ln)
			}
		}
	}
	return
}

// ============================
// Assign Group
// ============================

// AssignGroupNode is node include id
type AssignGroupNode struct {
	assignTags   []string
	unassignTags []string
}

// NewAssignGroupNode return IDGroup
func NewAssignGroupNode(assignTag, unassignTag string) *AssignGroupNode {
	result := &AssignGroupNode{
		assignTags:   []string{},
		unassignTags: []string{},
	}
	if assignTag != "" {
		result.AssignTag(assignTag)
	}
	if unassignTag != "" {
		result.UnassignTag(unassignTag)
	}
	return result
}

// AssignTag for task
func (agn *AssignGroupNode) AssignTag(tag string) *AssignGroupNode {
	agn.assignTags = append(agn.assignTags, tag)
	agn.removeRepeatedTags()
	return agn
}

// UnassignTag for task
func (agn *AssignGroupNode) UnassignTag(tag string) *AssignGroupNode {
	agn.unassignTags = append(agn.unassignTags, tag)
	agn.removeRepeatedTags()
	return agn
}

func (agn *AssignGroupNode) explain() string {
	result := ""
	fmt.Print("assign tags: ", agn.assignTags)
	for _, assignTag := range agn.assignTags {
		result += fmt.Sprint(" +", assignTag)
	}
	fmt.Print(" unassign tags: ", agn.unassignTags)
	for _, unassignTag := range agn.unassignTags {
		result += fmt.Sprint(" -", unassignTag)
	}
	return result
}

func (agn *AssignGroupNode) filter(tasks []model.Task) (result []model.Task) {
	result = []model.Task{}
	filterTagIDs := []int{}
	// fmt.Println("tags: ", agn.assignTags)
	for _, tag := range agn.assignTags {
		tagID, ok := controller.GetTagIDByName(tag)
		if !ok {
			ErrorList = append(ErrorList, errors.New("got inexist tag while finding tasks"))
		} else {
			filterTagIDs = append(filterTagIDs, tagID)
		}

	}
	for _, task := range tasks {
		if controller.CheckTaskByTags(task.ID, filterTagIDs) {
			result = append(result, task)
		}
	}
	return
}

func (agn *AssignGroupNode) sortTags() {
	sort.Strings(agn.assignTags)
	sort.Strings(agn.unassignTags)
}

func (agn *AssignGroupNode) removeRepeatedTags() {
	if len(agn.assignTags) > 0 {
		temp := []string{agn.assignTags[0]}
		agn.sortTags()
		for i := 1; i < len(agn.assignTags); i++ {
			if agn.assignTags[i] != agn.assignTags[i-1] {
				temp = append(temp, agn.assignTags[i])
			}
		}
		agn.assignTags = temp
	}
	if len(agn.unassignTags) > 0 {
		temp := []string{agn.unassignTags[0]}
		agn.sortTags()
		for i := 1; i < len(agn.unassignTags); i++ {
			if agn.unassignTags[i] != agn.unassignTags[i-1] {
				temp = append(temp, agn.unassignTags[i])
			}
		}
		agn.unassignTags = temp
	}
}

// ============================
// Time Filter
// ============================

// TimeFilterNode contains start/end TimeNode
type TimeFilterNode struct {
	startTime *TimeNode
	endTime   *TimeNode
}

// NewTimeFilterNode return *TimeFilterNode
func NewTimeFilterNode(s, e *TimeNode) *TimeFilterNode {
	return &TimeFilterNode{s, e}
}

func (tfn *TimeFilterNode) filter(tasks []model.Task) []model.Task {
	return tasks
}

func (tfn *TimeFilterNode) explain() string {
	if tfn.startTime != nil && tfn.endTime != nil {

		fmt.Printf("from ")
		rs := tfn.startTime.explain()
		fmt.Printf(" to ")
		re := tfn.endTime.explain()
		fmt.Print("\n")
		return rs + "-" + re
	} else if tfn.startTime != nil {
		fmt.Print("after ")
		rs := tfn.startTime.explain()
		fmt.Print("\n")
		return rs
	} else {
		fmt.Print("before")
		re := tfn.endTime.explain()
		fmt.Print("\n")
		return "-" + re
	}
}

// TimeNode contains a single time
type TimeNode struct {
	time *time.Time
}

// NewTimeNode return *TimeNode
func NewTimeNode(str string, dtType int) *TimeNode {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err.Error())
		ErrorList = append(ErrorList, err)
	}
	str = completeDateTime(str, dtType)
	time, err := time.ParseInLocation("2006/01/02 15:04:05", str, loc)
	if err != nil {
		fmt.Println(err.Error())
		ErrorList = append(ErrorList, err)
	}
	return &TimeNode{&time}
}

func (tn *TimeNode) execute() {}

func (tn *TimeNode) explain() string {
	fmt.Print(tn.time.Format("2006/01/02 15:04:05"))
	return tn.time.Format("2006/01/02 15:04:05")
}

func completeDateTime(str string, dtType int) string {
	now := time.Now()
	d, t := now.Format("2006/01/02"), "00:00:00"
	switch dtType {
	case TimeFormatDate:
		d = completeDate(str)
	case TimeFormatTime:
		t = completeTime(str)
	case TimeFormatDateTime:
		dateTimeParts := strings.Split(str, " ")
		d, t = completeDate(dateTimeParts[0]), completeTime(dateTimeParts[1])
	}
	return d + " " + t
}

func completeDate(d string) string {
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

func completeTime(t string) string {
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
