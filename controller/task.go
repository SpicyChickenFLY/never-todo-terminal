package controller

import (
	"errors"
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

// ListTodoTasks with filter provided by params
func ListTodoTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.TaskTodo {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDoneTasks with filter provided by params
func ListDoneTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.TaskDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTaskByID called by parser
func GetTaskByID(id int) (model.Task, error) {
	task, ok := model.DB.Data.Tasks[id]
	if !ok {
		return model.Task{}, errors.New("Retrieve: Task(id:%d) not found")
	}
	return task, nil
}

// GetTasksByIDGroup called by parser
func GetTasksByIDGroup(ids []int) (tasks []model.Task, warnList []string) {
	for _, id := range ids {
		task, err := GetTaskByID(id)
		if err == nil {
			tasks = append(tasks, task)
		} else {
			warnList = append(warnList,
				err.Error(),
			)
		}
	}
	return tasks, warnList
}

// DeleteTasks called by parser
func DeleteTasks(ids []int) (warnList []string) {
	for _, id := range ids {
		if _, ok := model.DB.Data.Tasks[id]; ok {
			// task.Status = model.ProjectDeleted
			// model.DB.Data.Tasks[id] = task
			delete(model.DB.Data.Tasks, id)
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Delete: Task(id=%d) not found", id),
			)
		}
	}
	return
}

// CompleteTask called by parse
func CompleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.Status = model.TaskDone
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Complete: Task(id=%d) not found", id),
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
		return fmt.Errorf("Update: Task(id=%d) not found", updateTask.ID)
	}
	model.DB.Data.Tasks[updateTask.ID] = updateTask
	return nil
}

// SortTask with specified metric
func SortTask(tasks []model.Task, metricName string) []model.Task {
	switch metricName {
	case "NAME":
		sort.SliceStable(tasks, func(a, b int) bool {
			return !utils.LessInString(tasks[a].Content, tasks[a].Content)
		})
	case "name":
		sort.SliceStable(tasks, func(a, b int) bool {
			return utils.LessInString(tasks[a].Content, tasks[a].Content)
		})
	case "DUE":
		sort.SliceStable(tasks, func(a, b int) bool {
			return !utils.LessInTime(tasks[a].Due, tasks[b].Due)
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
