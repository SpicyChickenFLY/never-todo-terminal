package ast

import (
	"fmt"
	"sort"

	"github.com/SpicyChickenFLY/never-todo-cmd/data"
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
	fmt.Println("with following ID: ", ign.idGroup)
}

func (ign *IDGroupNode) sortID() {
	sort.Ints(ign.idGroup)
}

func (ign *IDGroupNode) removeRepeatedIDs() {
	if len(ign.idGroup) == 0 {
		return
	}
	temp := []int{ign.idGroup[0]}
	ign.sortID()
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
	operands []*ContentGroupNode
	operator int
}

// NewContentGroupNode return ContentGroupNode
func NewContentGroupNode(
	content string, operator int, operands []*ContentGroupNode) *ContentGroupNode {
	return &ContentGroupNode{content, operands, operator}
}

func (cgn *ContentGroupNode) filter(tags []data.Tag) []data.Tag {
	switch cgn.operator {
	case OPNone:
		return tags

	default:
		return tags
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
	fmt.Println("assign following tags: ", agn.assignGroup)
	fmt.Println("unassign following tags: ", agn.unassignGroup)
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
