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
	tln.taskListFilterNode.explain()
}
func (tln *TaskListNode) restore() string {
	return "todo list " + tln.taskListFilterNode.restore()
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
		fmt.Println("list todo tasks")
		tlfn.idGroup.explain()
	} else if tlfn.indefiniteTaskListFilter != nil {
		fmt.Println("list todo tasks")
		tlfn.indefiniteTaskListFilter.explain()
	} else {
		fmt.Println("list all todo tasks")
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
	contentGroup *ContentGroupNode
	assignGroup  *AssignGroupNode
}

// NewIndefiniteTaskListFilterNode return IndefiniteTaskListFilterNode
func NewIndefiniteTaskListFilterNode() *IndefiniteTaskListFilterNode {
	return &IndefiniteTaskListFilterNode{}
}

// SetContentFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetContentFilter(cgn *ContentGroupNode) {
	itlfn.contentGroup = cgn
}

// SetAssignFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetAssignFilter(agn *AssignGroupNode) {
	itlfn.assignGroup = agn
}

// SetAgeFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetAgeFilter(tfn *TimeFilterNode) {
	fmt.Println("set age")
}

// SetDueFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetDueFilter(tfn *TimeFilterNode) {
	fmt.Println("set due")
}

func (itlfn *IndefiniteTaskListFilterNode) execute() error { return nil }
func (itlfn *IndefiniteTaskListFilterNode) explain() {
	if itlfn.contentGroup != nil {
		fmt.Print("\tby content ")
		itlfn.contentGroup.explain()
		fmt.Print("\n")
	}
	if itlfn.assignGroup != nil {
		fmt.Print("\tby assign ")
		itlfn.assignGroup.explain()
	}
}
func (itlfn *IndefiniteTaskListFilterNode) restore() string { return "" }

// ============================
// Task Add·
// ============================

type TaskAddNode struct {
	content string
	option  *TaskAddOptionNode
}

func NewTaskAddNode(c string, opt *TaskAddOptionNode) *TaskAddNode {
	if len(c) > 0 && c[0] == ' ' {
		c = c[1:]
	}
	return &TaskAddNode{c, opt}
}

func (tan *TaskAddNode) execute() error { return nil }
func (tan *TaskAddNode) explain() {
	fmt.Println("add new todo task")
	fmt.Printf("\twith content `%s`\n", tan.content)
	if tan.option != nil {
		tan.option.explain()
	}

}
func (tan *TaskAddNode) restore() string {
	return "todo add " // + tan.taskAddOptionNode.restore()
}

type TaskAddOptionNode struct {
	AssignGroupNode
}

func NewTaskAddOptionNode() *TaskAddOptionNode {
	return &TaskAddOptionNode{}
}

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
func NewTaskUpdateNode(id int, content string, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
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
