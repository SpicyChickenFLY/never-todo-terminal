package ast

import (
	"fmt"
	"sort"
	"time"
)

// ============================
// ID Group
// ============================

// IDGroupNode is node include id
type IDGroupNode struct {
	idGroup []int
}

// NewIDGroupNode return IDGroup
func NewIDGroupNode(ids ...int) *IDGroupNode {
	switch len(ids) {
	case 1:
		return &IDGroupNode{idGroup: []int{ids[0]}}
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
		return &IDGroupNode{idGroup: temp}
	default:
		return &IDGroupNode{}
	}
}

// MergeIDNode merge with othen IDGroup
func (ign *IDGroupNode) MergeIDNode(idNode *IDGroupNode) {
	ign.idGroup = append(ign.idGroup, idNode.idGroup...)
	ign.removeRepeatedIDs()
}

// Explain which id will be used
func (ign *IDGroupNode) explain() string {
	fmt.Printf("\twith ID: %v\n", ign.idGroup)
	result := ""
	for _, id := range ign.idGroup {
		result += fmt.Sprint(" ", id)
	}
	return result[1:]
}

func (ign *IDGroupNode) sortID() {
	sort.Ints(ign.idGroup)
}

func (ign *IDGroupNode) removeRepeatedIDs() {
	if len(ign.idGroup) == 0 {
		return
	}
	ign.sortID()
	temp := []int{ign.idGroup[0]}
	for i := 1; i < len(ign.idGroup); i++ {
		if ign.idGroup[i] != ign.idGroup[i-1] {
			temp = append(temp, ign.idGroup[i])
		}
	}
	ign.idGroup = temp
}

// ============================
// Content Group
// ============================

const ( // operator type
	// OPNone indicate none command
	OPNone = iota + 0
	// OPNOT not
	OPNOT
	// OPAND and
	OPAND
	// OPOR or
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
	return &ContentGroupNode{content, operator, operands}
}

// func (cgn *ContentGroupNode) filter() {
// 	switch cgn.operator {
// 	case OPNone:
// 		return
// 	default:
// 		return
// 	}
// }

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

// ============================
// Assign Group
// ============================

// AssignGroupNode is node include id
type AssignGroupNode struct {
	assignTags   []string
	unassignTags []string
}

// NewAssignGroupNode return IDGroup
func NewAssignGroupNode() *AssignGroupNode {
	return &AssignGroupNode{}
}

// AssignTag for task
func (agn *AssignGroupNode) AssignTag(tag string) {
	agn.assignTags = append(agn.assignTags, tag)
	agn.removeRepeatedTags()
}

// UnassignTag for task
func (agn *AssignGroupNode) UnassignTag(tag string) {
	agn.unassignTags = append(agn.unassignTags, tag)
	agn.removeRepeatedTags()
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

type TimeFilterNode struct {
	startTime *TimeNode
	endTime   *TimeNode
}

func NewTimeFilterNode(s, e *TimeNode) *TimeFilterNode {
	return &TimeFilterNode{s, e}
}

func (tfn *TimeFilterNode) execute() error { return nil }
func (tfn *TimeFilterNode) explain() string {
	if tfn.startTime != nil && tfn.endTime != nil {
		rs, re := tfn.startTime.explain(), tfn.endTime.explain()
		fmt.Printf("from %s to %s", rs, re)
		return rs + "-" + re
	} else if tfn.startTime != nil {
		rs := tfn.startTime.explain()
		fmt.Printf("after %s", rs)
		return rs
	} else {
		re := tfn.endTime.explain()
		fmt.Printf("before %s", re)
		return "-" + re
	}
}

type TimeNode struct {
	time *time.Time
}

func NewTimeNode(str, format string) *TimeNode {
	loc, err := time.LoadLocation("Asia/Shanghai")
	//TODO: handle this
	if err != nil {
		fmt.Println(err.Error())
	}
	time, err := time.ParseInLocation(format, str, loc)
	//TODO: handle this
	if err != nil {
		fmt.Println(err.Error())
	}
	return &TimeNode{&time}
}

func (tn *TimeNode) execute() error { return nil }
func (tn *TimeNode) explain() string {
	fmt.Print(tn.time.Format("2006/01/02 15:04:05"))
	return tn.time.Format("2006/01/02 15:04:05")
}
