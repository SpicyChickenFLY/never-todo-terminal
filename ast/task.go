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
func (itlfn *IndefiniteTaskListFilterNode) explain() {
	if itlfn.contentGroup != nil {
		fmt.Print("\tby content ")
		itlfn.contentGroup.explain()
		fmt.Print("\n")
	}
	if itlfn.assignGroup != nil {
		fmt.Print("\tby assign ")
		itlfn.assignGroup.explain()
		fmt.Print("\n")
	}
	if itlfn.age != nil {
		fmt.Print("\twhich created ")
		itlfn.age.explain()
		fmt.Print("\n")
	}
	if itlfn.due != nil {
		fmt.Print("\twhich created ")
		itlfn.due.explain()
		fmt.Print("\n")
	}
}
func (itlfn *IndefiniteTaskListFilterNode) restore() string { return "" }

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
func (taon *TaskAddOptionNode) explain() {
	fmt.Println("\tset importance: ", taon.importance)
	if taon.assignGroupNode != nil {
		fmt.Print("\t")
		taon.assignGroupNode.explain()
		fmt.Print("\n")
	}
	if taon.due != nil {
		fmt.Print("\twhich will due ")
		taon.due.explain()
		fmt.Print("\n")
	}
}
func (taon *TaskAddOptionNode) restore() string {
	return "todo add " // + tan.taskAddOptionNode.restore()
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
	option  TaskUpdateOptionNode
}

// NewTaskUpdateNode return TaskUpdateNode
func NewTaskUpdateNode(id int, content string, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
	return &TaskUpdateNode{}
}

// Restore to statement
func (tun *TaskUpdateNode) restore() string {
	result := fmt.Sprintf("task set %d %s", tun.id, tun.content)
	result += tun.option.restore()
	return result
}

// Explain to statement
func (tun *TaskUpdateNode) explain() {
	fmt.Printf("Update task:%d to \"%s\"", tun.id, tun.content)
	tun.option.explain()
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
func (tuon *TaskUpdateOptionNode) explain() {
	if len(tuon.assignGroupNode.assignTags) > 0 {
		fmt.Printf("\tassign tags:%v for it", tuon.assignGroupNode.assignTags)
	}
	if len(tuon.assignGroupNode.unassignTags) > 0 {
		fmt.Printf("\tunassign tags:%v for it", tuon.assignGroupNode.unassignTags)
	}

	fmt.Println("\tset importance: ", tuon.importance)
	if tuon.assignGroupNode != nil {
		fmt.Print("\t")
		tuon.assignGroupNode.explain()
		fmt.Print("\n")
	}
	if tuon.due != nil {
		fmt.Print("\twhich will due ")
		tuon.due.explain()
		fmt.Print("\n")
	}
}
func (tuon *TaskUpdateOptionNode) restore() string {
	return "todo set " // + tan.taskAddOptionNode.restore()
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
