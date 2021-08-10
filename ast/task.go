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
func (tln *TaskListNode) explain() string {
	return "todo list " + tln.taskListFilterNode.explain()
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
func (tlfn *TaskListFilterNode) explain() string {
	if tlfn.idGroup == nil && tlfn.indefiniteTaskListFilter == nil {
		fmt.Println("list all todo tasks")
		return ""
	}
	result := ""
	fmt.Println("list todo tasks")
	if tlfn.idGroup != nil {
		result += tlfn.idGroup.explain()
	} else if tlfn.indefiniteTaskListFilter != nil {
		result += tlfn.indefiniteTaskListFilter.explain()
	}
	return result
}

// IndefiniteTaskListFilterNode include content assignGroup age due
type IndefiniteTaskListFilterNode struct {
	contentGroup *ContentGroupNode
	assignGroup  *AssignGroupNode
	age          *TimeFilterNode
	due          *TimeFilterNode
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
	itlfn.age = tfn
}

// SetDueFilter func
func (itlfn *IndefiniteTaskListFilterNode) SetDueFilter(tfn *TimeFilterNode) {
	itlfn.due = tfn
}

func (itlfn *IndefiniteTaskListFilterNode) execute() error { return nil }
func (itlfn *IndefiniteTaskListFilterNode) explain() string {
	result := ""
	if itlfn.contentGroup != nil {
		fmt.Print("\tby content ")
		result += itlfn.contentGroup.explain() + " "
		fmt.Print("\n")
	}
	if itlfn.assignGroup != nil {
		fmt.Print("\tby assign ")
		result += itlfn.assignGroup.explain() + " "
		fmt.Print("\n")
	}
	if itlfn.age != nil {
		fmt.Print("\twhich created ")
		result += itlfn.age.explain() + " "
		fmt.Print("\n")
	}
	if itlfn.due != nil {
		fmt.Print("\twhich created ")
		result += itlfn.due.explain() + " "
		fmt.Print("\n")
	}
	return result
}

// ============================
// Task Add·
// ============================

// TaskAddNode include task content and add option
type TaskAddNode struct {
	content string
	option  *TaskAddOptionNode
}

// NewTaskAddNode return *TaskAddNode
func NewTaskAddNode(c string, opt *TaskAddOptionNode) *TaskAddNode {
	if len(c) > 0 && c[0] == ' ' {
		c = c[1:]
	}
	return &TaskAddNode{c, opt}
}

func (tan *TaskAddNode) execute() error { return nil }
func (tan *TaskAddNode) explain() string {
	result := "todo add "
	fmt.Println("add new todo task")
	fmt.Printf("\twith content `%s`\n", tan.content)
	result += fmt.Sprintf("`%s` ", tan.content)
	if tan.option != nil {
		result += tan.option.explain()
	}
	return result
}

// TaskAddOptionNode include add options
type TaskAddOptionNode struct {
	importance      bool
	assignGroupNode *AssignGroupNode
	due             *TimeFilterNode
}

// NewTaskAddOptionNode return *TaskAddOptionNode
func NewTaskAddOptionNode() *TaskAddOptionNode {
	return &TaskAddOptionNode{}
}

func (taon *TaskAddOptionNode) execute() error { return nil }
func (taon *TaskAddOptionNode) explain() string {
	result := ""
	fmt.Println("\tset importance: ", taon.importance)
	if taon.importance {
		result += "!1 "
	} else {
		result += "!0 "
	}
	if taon.assignGroupNode != nil {
		fmt.Print("\t")
		result += taon.assignGroupNode.explain() + " "
		fmt.Print("\n")
	}
	if taon.due != nil {
		fmt.Print("\twhich will due ")
		result += "due:" + taon.due.explain() + " "
		fmt.Print("\n")
	}
	return result
}

//
func (taon *TaskAddOptionNode) SetImportance(importance int) {
	taon.importance = (importance > 0)
}

func (taon *TaskAddOptionNode) SetAssignGroup(assignGourp *AssignGroupNode) {
	taon.assignGroupNode = assignGourp
}

func (taon *TaskAddOptionNode) SetDue(due *TimeFilterNode) {
	taon.due = due
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

// Explain to statement
func (tdn *TaskDoneNode) explain() string {
	result := "task done "
	fmt.Println("complete task ")
	result += tdn.idGroup.explain()
	return result
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

func (tdn *TaskDeleteNode) explain() string {
	result := "task del "
	fmt.Println("delete task ")
	result += tdn.idGroup.explain()
	return result
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
	option  TaskUpdateOptionNode
}

// NewTaskUpdateNode return TaskUpdateNode
func NewTaskUpdateNode(id int, content string, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
	return &TaskUpdateNode{}
}

func (tun *TaskUpdateNode) explain() string {
	result := fmt.Sprintf("task set %d %s ", tun.id, tun.content)
	fmt.Printf("Update task:%d to \"%s\"", tun.id, tun.content)
	result += tun.option.explain()
	return result
}

// Execute complete task logic
func (tun *TaskUpdateNode) execute() error {
	// return controller.UpdateTask(tun.id, tun.content,tun)
	return nil
}

// ============================
// Task Update Filter
// ============================

// TaskUpdateOptionNode is node for task update option
type TaskUpdateOptionNode struct {
	importance      bool
	assignGroupNode *AssignGroupNode
	due             *TimeNode
}

// NewTaskUpdateOptionNode return *TaskUpdateOptionNode
func NewTaskUpdateOptionNode() *TaskUpdateOptionNode {
	return &TaskUpdateOptionNode{}
}

func (tuon *TaskUpdateOptionNode) execute() error { return nil }
func (tuon *TaskUpdateOptionNode) explain() string {
	result := ""
	if len(tuon.assignGroupNode.assignTags) > 0 {
		fmt.Printf("\tassign tags:%v for it", tuon.assignGroupNode.assignTags)
	}
	if len(tuon.assignGroupNode.unassignTags) > 0 {
		fmt.Printf("\tunassign tags:%v for it", tuon.assignGroupNode.unassignTags)
	}
	result += tuon.assignGroupNode.explain() + " "

	fmt.Println("\tset importance: ", tuon.importance)
	if tuon.importance {
		result += "!1"
	} else {
		result += "!0"
	}
	if tuon.due != nil {
		fmt.Print("\twhich will due ")
		result += tuon.due.explain()
		fmt.Print("\n")
	}
	return result
}

func (tuon *TaskUpdateOptionNode) SetImportance(importance int) {
	tuon.importance = (importance > 0)
}

func (tuon *TaskUpdateOptionNode) SetAssignGroup(assignGourp *AssignGroupNode) {
	tuon.assignGroupNode = assignGourp
}

func (tuon *TaskUpdateOptionNode) SetDue(due *TimeNode) {
	tuon.due = due
}
