package controller

import (
	"fmt"
	"sort"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

// ListTasks with filter provided by params
func ListTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// ListAllTasks with filter provided by params
func ListAllTasks() (todo, done, deleted []model.Task) {
	return ListTodoTasks(), ListDoneTasks(), ListDeletedTasks()
}

// ListTodoTasks with filter provided by params
func ListTodoTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.ProjectTodo {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDoneTasks with filter provided by params
func ListDoneTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.ProjectDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDeletedTasks with filter provided by params
func ListDeletedTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.ProjectDeleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTaskByID called by parser
func GetTaskByID(id int) (model.Task, bool) {
	task, ok := model.DB.Data.Tasks[id]
	return task, ok
}

// GetTasksByIDGroup called by parser
func GetTasksByIDGroup(ids []int) (tasks []model.Task, warnList []string) {
	for _, id := range ids {
		task, ok := GetTaskByID(id)
		if ok {
			tasks = append(tasks, task)
		} else {
			warnList = append(warnList, fmt.Sprintf("Task(id:%d) not found", id))
		}
	}
	return tasks, warnList
}

// DeleteTask called by parser
func DeleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.Status = model.ProjectDeleted
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Task(id=%d) not found", id),
			)
		}
	}
	return
}

// CompleteTask called by parse
func CompleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.Status = model.ProjectDone
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Task(id=%d) not found", id),
			)
		}
	}
	return
}

// AddTask called by parser
func AddTask(content string) (taskID int) {
	newTask := model.Task{
		ID:      model.DB.Data.TaskInc,
		Content: content,
	}
	model.DB.Data.Tasks[newTask.ID] = newTask
	model.DB.Data.TaskTags[newTask.ID] = make(map[int]bool, 0)
	model.DB.Data.TaskInc--
	return newTask.ID
}

// UpdateTask called by parser
func UpdateTask(updateTask model.Task) error {
	if _, ok := model.DB.Data.Tasks[updateTask.ID]; !ok {
		return fmt.Errorf("task(id=%d) not found", updateTask.ID)
	}
	model.DB.Data.Tasks[updateTask.ID] = updateTask
	return nil
}

// TaskSortDue sorted by task id
type TaskSortDue []model.Task

func (tsd TaskSortDue) Len() int           { return len(tsd) }
func (tsd TaskSortDue) Less(i, j int) bool { return utils.LessInAbs(tsd[i].ID, tsd[j].ID) }
func (tsd TaskSortDue) Swap(i, j int)      { tsd[i], tsd[j] = tsd[j], tsd[i] }

// SortTask with specified metric
func SortTask(tasks []model.Task, metricName string) []model.Task {
	switch metricName {
	case "NAME":
		sort.SliceStable(tasks, func(a, b int) bool {
			return true
		})
	case "name":
		sort.SliceStable(tasks, func(a, b int) bool {
			return true
		})
	case "DUE":
		sort.SliceStable(tasks, func(a, b int) bool {
			return utils.LessInTime(tasks[a].Due, tasks[b].Due)
		})
	case "due":
		sort.SliceStable(tasks, func(a, b int) bool {
			return utils.LessInTime(tasks[a].Due, tasks[b].Due)
		})
	case "ID":
		sort.SliceStable(tasks, func(a, b int) bool {
			return !utils.LessInID(tasks[a].ID, tasks[b].ID)
		})
	default:
		sort.SliceStable(tasks, func(a, b int) bool {
			return utils.LessInID(tasks[a].ID, tasks[b].ID)
		})
	}
	return tasks
}
