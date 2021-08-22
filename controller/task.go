package controller

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ListTasks with filter provided by params
func ListTasks() []model.Task {
	return model.DB.Data.Tasks
}

// ListAllTasks with filter provided by params
func ListAllTasks() (todo, done, deleted []model.Task) {
	return ListTodoTasks(), ListDoneTasks(), ListDeletedTasks()
}

// ListTodoTasks with filter provided by params
func ListTodoTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.ProjectID == model.ProjectTodo {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDoneTasks with filter provided by params
func ListDoneTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if !task.Deleted && task.Completed {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDeletedTasks with filter provided by params
func ListDeletedTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Deleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// FindTaskByID called by parser
func FindTaskByID(id int) (model.Task, bool) {
	for _, task := range model.DB.Data.Tasks {
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
			warnList = append(warnList, fmt.Sprintf("Task(id:%d) not found", id))
		}
	}
	return tasks, warnList
}

// DeleteTask called by parser
func DeleteTask(ids []int) {
	// delete task
	for _, id := range ids {
		for i := range model.DB.Data.Tasks {
			if model.DB.Data.Tasks[i].ID == id {
				model.DB.Data.Tasks[i].Deleted = true
			}
		}
	}
}

// CompleteTask called by parse
func CompleteTask(ids []int) {
	for _, id := range ids {
		for _, task := range model.DB.Data.Tasks {
			if task.ID == id {
				task.Completed = true
			}
		}
	}
}

// AddTask called by parser
func AddTask(content string) (taskID int) {
	newTask := model.Task{
		ID:      model.DB.Data.TaskInc,
		Content: content,
	}
	model.DB.Data.Tasks = append(model.DB.Data.Tasks, newTask)
	model.DB.Data.TaskInc--
	return newTask.ID
}

// UpdateTask called by parser
func UpdateTask(updateTask model.Task) error {
	for i := range model.DB.Data.Tasks {
		if model.DB.Data.Tasks[i].ID == updateTask.ID {
			model.DB.Data.Tasks[i] = updateTask
			return nil
		}
	}
	return errors.New("not found the task to be updated")
}
