package ast

import (
	"errors"
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/render"
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

func (tln *TaskListNode) execute() { tln.taskListFilterNode.execute() }
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

func (tlfn *TaskListFilterNode) execute() {
	if tlfn.idGroup != nil {
		tasks, warnList := controller.FindTasksByIDGroup(tlfn.idGroup.idGroup)
		render.Tasks(tasks)
		warnList = append(warnList, warnList...)
		return
	} else if tlfn.indefiniteTaskListFilter != nil {
		tlfn.indefiniteTaskListFilter.execute()
	} else {
		tasks := controller.ListTasks()
		WarnList = append(WarnList, render.Tasks(tasks)...)
		return
	}
}

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
	importance   bool
	contentGroup *ContentGroupNode
	assignGroup  *AssignGroupNode
	age          *TimeFilterNode
	due          *TimeFilterNode
}

// NewIndefiniteTaskListFilterNode return IndefiniteTaskListFilterNode
func NewIndefiniteTaskListFilterNode() *IndefiniteTaskListFilterNode {
	return &IndefiniteTaskListFilterNode{}
}

func (itlfn *IndefiniteTaskListFilterNode) execute() {
	tasks := controller.ListTasks()
	tasks = itlfn.contentGroup.filter(tasks)
	tasks = itlfn.assignGroup.filter(tasks)
	tasks = itlfn.age.filter(tasks)
	tasks = itlfn.due.filter(tasks)
	render.Tasks(tasks)
}
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

// SetImportance for TaskUpdateOptionNode
func (itlfn *IndefiniteTaskListFilterNode) SetImportance(importance int) *IndefiniteTaskListFilterNode {
	itlfn.importance = (importance > 0)
	return itlfn
}

// SetContentGroup alway keep the first one
func (itlfn *IndefiniteTaskListFilterNode) SetContentGroup(cgn *ContentGroupNode) *IndefiniteTaskListFilterNode {
	if itlfn.contentGroup != nil {
		WarnList = append(WarnList, "Only one content will be accepted")
	}
	itlfn.contentGroup = cgn
	return itlfn
}

// MergeAssignGroup merge assign group
func (itlfn *IndefiniteTaskListFilterNode) MergeAssignGroup(agn *AssignGroupNode) *IndefiniteTaskListFilterNode {
	if itlfn.assignGroup != nil {
		itlfn.assignGroup.assignTags = append(itlfn.assignGroup.assignTags, agn.assignTags...)
		itlfn.assignGroup.unassignTags = append(itlfn.assignGroup.unassignTags, agn.unassignTags...)
	} else {
		itlfn.assignGroup = agn
	}
	return itlfn
}

// SetAge func
func (itlfn *IndefiniteTaskListFilterNode) SetAge(tfn *TimeFilterNode) *IndefiniteTaskListFilterNode {
	if itlfn.age != nil {
		WarnList = append(WarnList, "Only one age filter will be accepted")
	} else {
		itlfn.age = tfn
	}
	return itlfn
}

// SetDue func
func (itlfn *IndefiniteTaskListFilterNode) SetDue(tfn *TimeFilterNode) *IndefiniteTaskListFilterNode {
	if itlfn.due != nil {
		WarnList = append(WarnList, "Only one due filter will be accepted")
	} else {
		itlfn.due = tfn
	}
	return itlfn
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

func (tan *TaskAddNode) execute() {
	importance, assignTags, due, loop := tan.option.apply()
	taskID, err := controller.AddTask(
		tan.content,
		importance, assignTags, due, loop,
	)
	if err != nil {
		ErrorList = append(ErrorList, err)
		return
	}
	task, ok := controller.FindTaskByID(taskID)
	if !ok {
		ErrorList = append(ErrorList, errors.New("Added task is not found"))
		return
	}
	render.Tasks([]model.Task{task})

}
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

func (taon *TaskAddOptionNode) apply() (
	importance bool, assignTags []string, due *time.Time, loop string) {
	importance, assignTags, due, loop = taon.importance, []string{}, nil, ""
	if taon.assignGroupNode != nil {
		assignTags = taon.assignGroupNode.assignTags
	}
	if taon.due != nil {
		due = taon.due.startTime.time
	}
	return
}

// SetImportance for TaskAddOptionNode
func (taon *TaskAddOptionNode) SetImportance(importance int) *TaskAddOptionNode {
	taon.importance = (importance > 0)
	return taon
}

// SetAssignGroup for TaskAddOptionNode
func (taon *TaskAddOptionNode) SetAssignGroup(assignGourp *AssignGroupNode) *TaskAddOptionNode {
	taon.assignGroupNode = assignGourp
	return taon
}

// SetDue for TaskAddOptionNode
func (taon *TaskAddOptionNode) SetDue(due *TimeFilterNode) *TaskAddOptionNode {
	taon.due = due
	return taon
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
func (tdn *TaskDoneNode) execute() {
	controller.CompleteTask(tdn.idGroup.idGroup)
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

func (tdn *TaskDeleteNode) execute() {
	controller.CompleteTask(tdn.idGroup.idGroup)
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
func (tun *TaskUpdateNode) execute() {
	controller.UpdateTask(
		tun.id,
		tun.content,
		tun.option.importance,
		tun.option.assignGroupNode.assignTags,
		tun.option.assignGroupNode.unassignTags,
		tun.option.due.time,
	)
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

// SetImportance for TaskUpdateOptionNode
func (tuon *TaskUpdateOptionNode) SetImportance(importance int) *TaskUpdateOptionNode {
	tuon.importance = (importance > 0)
	return tuon
}

// SetAssignGroup for TaskUpdateOptionNode
func (tuon *TaskUpdateOptionNode) SetAssignGroup(assignGourp *AssignGroupNode) *TaskUpdateOptionNode {
	tuon.assignGroupNode = assignGourp
	return tuon
}

// SetDue for TaskUpdateOptionNode
func (tuon *TaskUpdateOptionNode) SetDue(due *TimeNode) *TaskUpdateOptionNode {
	tuon.due = due
	return tuon
}
