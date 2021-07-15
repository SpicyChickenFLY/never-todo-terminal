package ast

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
)

// ============================
// Task List
// ============================

type TaskListNode struct {
	taskListFilterNode *TaskListFilterNode
}

func NewTaskListNode() {}

type TaskListFilterNode struct {
	definiteTaskListFilterNode   *DefiniteTaskListFilterNode
	indefiniteTaskListFilterNode *IndefiniteTaskListFilterNode
}

func NewTaskListFilterNode() {}

type DefiniteTaskListFilterNode struct {
	idGroup *IDGroupNode
}

func NewDefiniteTaskListFilterNode() {}

type IndefiniteTaskListFilterNode struct {
	contentGroup *ContentGroupNode
	assignGroup  *AssignGroupNode
}

func NewIndefiniteTaskListFilterNode() {}

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
	result += tdn.idGroup.restore()
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
	result += tdn.idGroup.restore()
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
func NewTaskUpdateNode(id int, content string, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
	return &TaskUpdateNode{}
}

// Restore to statement
func (tun *TaskUpdateNode) Restore() string {
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

func (tuon TaskUpdateOptionNode) assignTag(agn *AssignGroupNode) {
}
