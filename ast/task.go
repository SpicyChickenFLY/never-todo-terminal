package ast

import (
	"errors"
	"fmt"

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

func (tln *TaskListNode) execute() {
	tasks := controller.ListTasks()
	render.Tasks(tln.taskListFilterNode.filter(tasks), "Tasks")
}

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

func (tlfn *TaskListFilterNode) filter(tasks []model.Task) (result []model.Task) {
	if tlfn.idGroup != nil {
		for _, id := range tlfn.idGroup.ids {
			for _, task := range tasks {
				if id == task.ID {
					result = append(result, task)
				}
			}
		}
		return result
	} else if tlfn.indefiniteTaskListFilter != nil {
		return tlfn.indefiniteTaskListFilter.filter(tasks)
	} else {
		return tasks
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
	project      string
	importance   int
	contentGroup *ContentGroupNode
	assignGroup  *AssignGroupNode
	age          *TimeFilterNode
	due          *TimeFilterNode
}

// NewIndefiniteTaskListFilterNode return IndefiniteTaskListFilterNode
func NewIndefiniteTaskListFilterNode() *IndefiniteTaskListFilterNode {
	return &IndefiniteTaskListFilterNode{}
}

func (itlfn *IndefiniteTaskListFilterNode) filter(tasks []model.Task) []model.Task {
	// tasks := controller.ListTodoTasks()
	tasks, _ = itlfn.contentGroup.filter(tasks)
	tasks = itlfn.assignGroup.filter(tasks)
	tasks = itlfn.age.filter(tasks)
	tasks = itlfn.due.filter(tasks)
	return tasks
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
	itlfn.importance = importance
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

// SetProject func
func (itlfn *IndefiniteTaskListFilterNode) SetProject(project string) *IndefiniteTaskListFilterNode {
	if itlfn.project != "" {
		WarnList = append(WarnList, "Only one project filter will be accepted")
	} else {
		itlfn.project = project
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
	taskID := controller.AddTask(tan.content)
	task, ok := controller.FindTaskByID(taskID)
	if !ok {
		ErrorList = append(ErrorList, errors.New("ddded task is not found"))
		return
	}
	tan.option.apply(task)

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
	importance      int
	assignGroupNode *AssignGroupNode
	due             *TimeFilterNode
}

// NewTaskAddOptionNode return *TaskAddOptionNode
func NewTaskAddOptionNode() *TaskAddOptionNode {
	return &TaskAddOptionNode{importance: -1}
}

func (taon *TaskAddOptionNode) explain() string {
	result := ""
	fmt.Println("\tset importance: ", taon.importance)
	if taon.importance >= 0 {
		result += fmt.Sprintf("!%d ", taon.importance)
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

func (taon *TaskAddOptionNode) apply(task model.Task) {
	if taon.assignGroupNode != nil {
		assignTags := taon.assignGroupNode.assignTags
		if err := controller.AddTaskTags(task.ID, assignTags); err != nil {
			ErrorList = append(ErrorList, err)
		}
	}
	task.Important = taon.importance
	if taon.due != nil {
		task.Due = *taon.due.startTime.time
	}
	if err := controller.UpdateTask(task); err != nil {
		ErrorList = append(ErrorList, err)
	}
	render.Tasks([]model.Task{task}, "Add")
}

// SetImportance for TaskAddOptionNode
func (taon *TaskAddOptionNode) SetImportance(importance int) *TaskAddOptionNode {
	taon.importance = importance
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
// Task Update
// ============================

// TaskUpdateNode is node for task update
type TaskUpdateNode struct {
	id     int
	option TaskUpdateOptionNode
}

// NewTaskUpdateNode return TaskUpdateNode
func NewTaskUpdateNode(id int, tuon *TaskUpdateOptionNode) *TaskUpdateNode {
	return &TaskUpdateNode{id, *tuon}
}

func (tun *TaskUpdateNode) explain() string {
	result := fmt.Sprintf("task set %d ", tun.id)
	fmt.Printf("Update task:%d ", tun.id)
	result += tun.option.explain()
	return result
}

// Execute complete task logic
func (tun *TaskUpdateNode) execute() {
	fmt.Println(tun.id)
	task, ok := controller.FindTaskByID(tun.id)
	if !ok {
		ErrorList = append(ErrorList, errors.New("updated task is not found"))
		return
	}
	tun.option.apply(task)
}

// ============================
// Task Update Option
// ============================

// TaskUpdateOptionNode is node for task update option
type TaskUpdateOptionNode struct {
	content         string
	importance      int
	assignGroupNode *AssignGroupNode
	due             *TimeNode
}

// NewTaskUpdateOptionNode return *TaskUpdateOptionNode
func NewTaskUpdateOptionNode() *TaskUpdateOptionNode {
	return &TaskUpdateOptionNode{importance: -1}
}

func (tuon *TaskUpdateOptionNode) explain() string {
	result := ""
	if tuon.content != "" {
		fmt.Printf("\tset content:%s", tuon.content)
		result += fmt.Sprintf("`%s` ", tuon.content)
	}
	if len(tuon.assignGroupNode.assignTags) > 0 {
		fmt.Printf("\tassign tags:%v for it", tuon.assignGroupNode.assignTags)
	}
	if len(tuon.assignGroupNode.unassignTags) > 0 {
		fmt.Printf("\tunassign tags:%v for it", tuon.assignGroupNode.unassignTags)
	}
	result += tuon.assignGroupNode.explain() + " "

	fmt.Println("\tset importance: ", tuon.importance)
	if tuon.importance >= 0 {
		result += fmt.Sprintf("!%d ", tuon.importance)
	}
	if tuon.due != nil {
		fmt.Print("\twhich will due ")
		result += tuon.due.explain()
		fmt.Print("\n")
	}
	return result
}

func (tuon *TaskUpdateOptionNode) apply(task model.Task) {
	if tuon.assignGroupNode != nil {
		assignTags := tuon.assignGroupNode.assignTags
		if err := controller.AddTaskTags(task.ID, assignTags); err != nil {
			ErrorList = append(ErrorList, err)
		}
		unassignTags := tuon.assignGroupNode.unassignTags
		if err := controller.DeleteTaskTags(task.ID, unassignTags); err != nil {
			ErrorList = append(ErrorList, err)
		}
	}
	if tuon.content != "" {
		task.Content = tuon.content
	}
	task.Important = tuon.importance
	if tuon.due != nil {
		task.Due = *tuon.due.time
	}
	if err := controller.UpdateTask(task); err != nil {
		ErrorList = append(ErrorList, err)
	}
	render.Tasks([]model.Task{task}, "Update")
}

// SetContent for TaskUpdateOptionNode
func (tuon *TaskUpdateOptionNode) SetContent(content string) *TaskUpdateOptionNode {
	tuon.content = content
	return tuon
}

// SetImportance for TaskUpdateOptionNode
func (tuon *TaskUpdateOptionNode) SetImportance(importance int) *TaskUpdateOptionNode {
	tuon.importance = importance
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
	fmt.Println("complete/uncomplete task ")
	result += tdn.idGroup.explain()
	return result
}

// Execute complete task logic
func (tdn *TaskDoneNode) execute() {
	controller.CompleteTask(tdn.idGroup.ids)
}

// ============================
// Task Todo
// ============================

// TaskTodoNode is node for complete task
type TaskTodoNode struct {
	idGroup *IDGroupNode
}

// NewTaskTodoNode return TaskTodoNode
func NewTaskTodoNode(ign *IDGroupNode) *TaskTodoNode {
	return &TaskTodoNode{ign}
}

// Explain to statement
func (tdn *TaskTodoNode) explain() string {
	result := "task Todo "
	fmt.Println("complete/uncomplete task ")
	result += tdn.idGroup.explain()
	return result
}

// Execute complete task logic
func (tdn *TaskTodoNode) execute() {
	controller.CompleteTask(tdn.idGroup.ids)
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
	controller.DeleteTask(tdn.idGroup.ids)
}
