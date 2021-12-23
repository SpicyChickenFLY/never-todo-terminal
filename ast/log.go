package ast

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// TODO: add log for task/tag

// ============================
// Log List
// ============================

// LogListNode include log list filter
type LogListNode struct {
	logListFilterNode *LogListFilterNode
}

// NewLogListNode return LogListNode
func NewLogListNode(tlfn *LogListFilterNode) *LogListNode {
	return &LogListNode{tlfn}
}

func (tln *LogListNode) execute() {
	// logs := controller.ListLogs()
	// render.Logs(tln.logListFilterNode.filter(logs), "Logs")
}

func (tln *LogListNode) explain() string {
	return "todo list " + tln.logListFilterNode.explain()
}

// ============================
// Log List Filter
// ============================

// LogListFilterNode include idGroup OR indefiniteLogListFilter
type LogListFilterNode struct {
	idGroup *IDGroupNode
}

// NewLogListFilterNode return LogListFilterNode
func NewLogListFilterNode(idGroup *IDGroupNode) *LogListFilterNode {
	return &LogListFilterNode{idGroup}
}

func (tlfn *LogListFilterNode) filter(logs []model.Log) (result []model.Log) {
	if tlfn.idGroup != nil {
		for _, id := range tlfn.idGroup.ids {
			for _, log := range logs {
				if id == log.ID {
					result = append(result, log)
				}
			}
		}
		return result
	}
	return logs
}

func (tlfn *LogListFilterNode) explain() string {
	if tlfn.idGroup == nil {
		fmt.Println("list all todo logs")
		return ""
	}
	result := ""
	fmt.Println("list todo logs")
	if tlfn.idGroup != nil {
		result += tlfn.idGroup.explain()
	}
	return result
}
