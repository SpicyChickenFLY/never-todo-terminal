package controller

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ListTasks with filter provided by params
func ListTasks() (tasks []model.Task) {
	for _, task := range model.M.Data.Tasks {
		if !task.Deleted && !task.Completed {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// FindTaskByID called by parser
func FindTaskByID(id int) (model.Task, bool) {
	for _, task := range model.M.Data.Tasks {
		if task.ID == id {
			return task, true
		}
	}
	return model.Task{}, false
}

// FindTasksByIDGroup called by parser
func FindTasksByIDGroup(ids []int) (tasks []model.Task, warnList []string) {
	for _, id := range ids {
		task, ok := FindTaskByID(id)
		if ok {
			tasks = append(tasks, task)
		} else {
			warnList = append(warnList, fmt.Sprintf("task(id:%d) not found", id))
		}
	}
	return tasks, warnList
}

// DeleteTask called by parser
func DeleteTask(ids []int) {
	// delete task
	for _, id := range ids {
		for _, task := range model.M.Data.Tasks {
			if task.ID == id {
				task.Deleted = true
			}
		}
	}
}

// CompleteTask called by parse
func CompleteTask(ids []int) {
	for _, id := range ids {
		for _, task := range model.M.Data.Tasks {
			if task.ID == id {
				task.Completed = true
			}
		}
	}
}

// AddTask called by parser
func AddTask(
	content string,
	importance bool,
	assignTags []string,
	due *time.Time,
	loop string) (int, error) {
	// TODO: add new task
	return 0, nil
}

// UpdateTask called by parser
func UpdateTask(
	id int,
	content string,
	importance bool,
	assignTags, unasssignTags []string,
	due *time.Time) error {
	for _, task := range model.M.Data.Tasks {
		if task.ID == id {
			task.Content = content
			task.Important = importance
			AddTaskTags(task.ID, assignTags)
			DeleteTaskTags(task.ID, unasssignTags)
			task.Due = *due
		}
	}
	return nil
}
