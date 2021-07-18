package ast

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
)

// ============================
// Task List
// ============================

// TaskListNode include task list filter
type TaskListNode struct {
	taskListFilterNode *TaskListFilterNode
}

// NewTaskListNode return TaskListNode
func NewTaskListNode(tlfn *TaskListFilterNode) *TaskListNode {
	return &TaskListNode{tlfn}
}

func (tln *TaskListNode) execute() error { return nil }
func (tln *TaskListNode) explain() {
	fmt.Println("list task ")
	tln.taskListFilterNode.explain()
	fmt.Println()
}
func (tln *TaskListNode) restore() string {
	return "task list " + tln.taskListFilterNode.restore()
}

// TaskListFilterNode include idGroup OR indefiniteTaskListFilter
type TaskListFilterNode struct {
	idGroup                  *IDGroupNode
	indefiniteTaskListFilter *IndefiniteTaskListFilterNode
}

// NewTaskListFilterNode return TaskListFilterNode
func NewTaskListFilterNode(
	idGroup *IDGroupNode,
	itlf *IndefiniteTaskListFilterNode) *TaskListFilterNode {
	return &TaskListFilterNode{idGroup, itlf}
}

func (tlfn *TaskListFilterNode) execute() error { return nil }
func (tlfn *TaskListFilterNode) explain() {
	if tlfn.idGroup != nil {
		tlfn.idGroup.explain()
	}
	if tlfn.indefiniteTaskListFilter != nil {
		tlfn.indefiniteTaskListFilter.explain()
	}
}
func (tlfn *TaskListFilterNode) restore() string {
	result := ""
	if tlfn.idGroup != nil {
		result += tlfn.idGroup.restore()
	}
	if tlfn.indefiniteTaskListFilter != nil {
		result += tlfn.indefiniteTaskListFilter.restore()
	}
	return result
}

// IndefiniteTaskListFilterNode include content assignGroup age due
type IndefiniteTaskListFilterNode struct {
	content     *ContentNode
	assignGroup *AssignGroupNode
}

// NewIndefiniteTaskListFilterNode return IndefiniteTaskListFilterNode
func NewIndefiniteTaskListFilterNode() *IndefiniteTaskListFilterNode {
	return &IndefiniteTaskListFilterNode{}
}

// SetContentFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetContentFilter(content *ContentNode) {
	itlfn.content = content
}

// SetAssignFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetAssignFilter(agn *AssignGroupNode) {
	fmt.Println("SetAssignFilter")
	itlfn.assignGroup = agn
}

// SetAgeFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetAgeFilter(str string) {}

// SetDueFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetDueFilter(str string) {}

func (itlfn *IndefiniteTaskListFilterNode) execute() error { return nil }

func (itlfn *IndefiniteTaskListFilterNode) explain() {
	if itlfn.content != nil {
		itlfn.content.explain()
	}
	if itlfn.assignGroup != nil {
		itlfn.assignGroup.explain()
	}
}
func (itlfn *IndefiniteTaskListFilterNode) restore() string { return "" }

// ============================
// Task Done
// ============================

// TaskDoneNode is node for complete task
type TaskDoneNode struct {
	idGroup *IDGroupNode
}

// NewTaskDoneNode return TaskDoneNode
func NewTaskDoneNode(ign *IDGroupNode) *TaskDoneNode {
	return &TaskDoneNode{ign}
}

// Restore to statement
func (tdn *TaskDoneNode) restore() string {
	result := "task done "
	result += tdn.idGroup.Restore()
	return result
}

// Explain to statement
func (tdn *TaskDoneNode) explain() {
	fmt.Println("complete task ")
	tdn.idGroup.explain()
}

// Execute complete task logic
func (tdn *TaskDoneNode) execute() error {
	return controller.CompleteTask(tdn.idGroup.idGroup)
}

// ============================
// Task Delete
// ============================

// TaskDeleteNode is node for delete task
type TaskDeleteNode struct {
	idGroup IDGroupNode
}

// NewTaskDeleteNode return TaskDeleteNode
func NewTaskDeleteNode(ign *IDGroupNode) *TaskDeleteNode {
	return &TaskDeleteNode{*ign}
}

func (tdn *TaskDeleteNode) restore() string {
	result := "task del "
	result += tdn.idGroup.Restore()
	return result
}

func (tdn *TaskDeleteNode) explain() {
	fmt.Println("delete task ")
	tdn.idGroup.explain()
}

func (tdn *TaskDeleteNode) execute() error {
	return controller.DeleteTask(tdn.idGroup.idGroup)
}

// ============================
// Task Update
// ============================

// TaskUpdateNode is node for task update
type TaskUpdateNode struct {
	id      int
	content string
	Option  TaskUpdateOptionNode
}

// NewTaskUpdateNode return TaskUpdateNode
func NewTaskUpdateNode(id int, content *ContentNode, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
	return &TaskUpdateNode{}
}

// Restore to statement
func (tun *TaskUpdateNode) restore() string {
	result := fmt.Sprintf("task set %d %s", tun.id, tun.content)
	result += tun.Option.restore()
	return result
}

// Explain to statement
func (tun *TaskUpdateNode) explain() {
	fmt.Printf("Update task:%d to \"%s\"", tun.id, tun.content)
	if len(tun.Option.assignGroup) > 0 {
		fmt.Printf(",\nthen assign tags:%v for it", tun.Option.assignGroup)
	}
	if len(tun.Option.unassignGroup) > 0 {
		fmt.Printf(",\nthen unassign tags:%v for it", tun.Option.unassignGroup)
	}
}

// Execute complete task logic
func (tun *TaskUpdateNode) execute() error {
	return controller.UpdateTask(tun.id, tun.content)
}

// ============================
// Task Update Filter
// ============================

// TaskUpdateOptionNode is node for task update option
type TaskUpdateOptionNode struct {
	AssignGroupNode
}

// NewTaskUpdateOptionNode return *TaskUpdateOptionNode
func NewTaskUpdateOptionNode() *TaskUpdateOptionNode {
	return &TaskUpdateOptionNode{}
}

func (tuon TaskUpdateOptionNode) restore() string {
	return ""
}

// AssignTag for task
func (tuon TaskUpdateOptionNode) AssignTag(agn *AssignGroupNode) {
}
