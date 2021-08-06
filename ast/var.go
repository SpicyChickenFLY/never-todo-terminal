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
	for _, id := range idNode.idGroup {
		ign.idGroup = append(ign.idGroup, id)
	}
	ign.removeRepeatedIDs()
}

// Restore to statement
func (ign *IDGroupNode) Restore() string { return ign.restore() }

func (ign *IDGroupNode) restore() string {
	result := ""
	for _, id := range ign.idGroup {
		result += fmt.Sprint(" ", id)
	}
	return result[1:]
}

// Explain which id will be used
func (ign *IDGroupNode) explain() {
	fmt.Printf("\twith ID: %v\n", ign.idGroup)
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
	// OPXOR xor
	OPXOR
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

func (cgn *ContentGroupNode) filter() {
	switch cgn.operator {
	case OPNone:
		return
	default:
		return
	}
}

func (cgn *ContentGroupNode) explain() {
	switch cgn.operator {
	case OPNone:
		fmt.Printf("include `%s`", cgn.content)
	case OPNOT:
		fmt.Print("while not (")
		cgn.operands[0].explain()
		fmt.Print(")")
	case OPAND:
		fmt.Print("(")
		cgn.operands[0].explain()
		fmt.Print(" and ")
		cgn.operands[1].explain()
		fmt.Print(")")
	case OPOR:
		fmt.Print("(")
		cgn.operands[0].explain()
		fmt.Print(" or ")
		cgn.operands[1].explain()
		fmt.Print(")")
	default:
		return
	}
}

func (cgn *ContentGroupNode) restore() string {
	switch cgn.operator {
	case OPNone:
		return cgn.content
	case OPAND:
		return fmt.Sprint("%s and %s")
	default:
		return ""
	}
}

// ============================
// Assign Group
// ============================

// AssignGroupNode is node include id
type AssignGroupNode struct {
	assignGroup   []string
	unassignGroup []string
}

// NewAssignGroupNode return IDGroup
func NewAssignGroupNode() *AssignGroupNode {
	return &AssignGroupNode{}
}

func (agn *AssignGroupNode) restore() string {
	result := ""
	for _, assignTag := range agn.assignGroup {
		result += fmt.Sprint(" +", assignTag)
	}
	for _, unassignTag := range agn.unassignGroup {
		result += fmt.Sprint(" -", unassignTag)
	}
	return result
}

func (agn *AssignGroupNode) explain() {
	fmt.Print("assign tags: ", agn.assignGroup)
	fmt.Print(" unassign tags: ", agn.unassignGroup)
}

// AssignTag for task
func (agn *AssignGroupNode) AssignTag(tag string) {
	agn.assignGroup = append(agn.assignGroup, tag)
	agn.removeRepeatedTags()
}

// UnassignTag for task
func (agn *AssignGroupNode) UnassignTag(tag string) {
	agn.unassignGroup = append(agn.unassignGroup, tag)
	agn.removeRepeatedTags()
}

func (agn *AssignGroupNode) sortTags() {
	sort.Strings(agn.assignGroup)
	sort.Strings(agn.unassignGroup)
}

func (agn *AssignGroupNode) removeRepeatedTags() {
	if len(agn.assignGroup) > 0 {
		temp := []string{agn.assignGroup[0]}
		agn.sortTags()
		for i := 1; i < len(agn.assignGroup); i++ {
			if agn.assignGroup[i] != agn.assignGroup[i-1] {
				temp = append(temp, agn.assignGroup[i])
			}
		}
		agn.assignGroup = temp
	}
	if len(agn.unassignGroup) > 0 {
		temp := []string{agn.unassignGroup[0]}
		agn.sortTags()
		for i := 1; i < len(agn.unassignGroup); i++ {
			if agn.unassignGroup[i] != agn.unassignGroup[i-1] {
				temp = append(temp, agn.unassignGroup[i])
			}
		}
		agn.unassignGroup = temp
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
func (tfn *TimeFilterNode) explain() {
	if tfn.startTime != nil && tfn.endTime != nil {
		fmt.Printf("from %s to %s",
			tfn.startTime.restore(),
			tfn.endTime.restore(),
		)
	} else if tfn.startTime != nil {
		fmt.Printf("after %s", tfn.startTime.restore())
	} else {
		fmt.Printf("before %s", tfn.endTime.restore())
	}
}
func (tfn *TimeFilterNode) restore() string {
	return "todo add " // + tan.taskAddOptionNode.restore()
}

type TimeNode struct {
	time *time.Time
}

func NewTimeNode(str, format string) *TimeNode {
	loc, err := time.LoadLocation("Asia/Shanghai")
	time, err := time.ParseInLocation(format, str, loc)
	//TODO: handle this
	if err != nil {
		fmt.Println(err.Error())
	}
	return &TimeNode{&time}
}

func (tn *TimeNode) execute() error { return nil }
func (tn *TimeNode) explain() {
	fmt.Print(tn.time.Format("2006/01/02 15:04:05"))
}
func (tn *TimeNode) restore() string {
	return tn.time.Format("2006/01/02 15:04:05")
}
