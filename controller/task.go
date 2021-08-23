package controller

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
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
		if task.ProjectID == model.ProjectTodo {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDoneTasks with filter provided by params
func ListDoneTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.ProjectID == model.ProjectDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDeletedTasks with filter provided by params
func ListDeletedTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.ProjectID == model.ProjectDeleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// FindTaskByID called by parser
func FindTaskByID(id int) (model.Task, bool) {
	task, ok := model.DB.Data.Tasks[id]
	return task, ok
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
func DeleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.ProjectID = model.ProjectDeleted
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprint("Task(id=%d) not found", id),
			)
		}
	}
	return
}

// CompleteTask called by parse
func CompleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.ProjectID = model.ProjectDone
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprint("Task(id=%d) not found", id),
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
	model.DB.Data.TaskInc--
	return newTask.ID
}

// UpdateTask called by parser
func UpdateTask(updateTask model.Task) error {
	if _, ok := model.DB.Data.Tasks[updateTask.ID]; !ok {
		return errors.New("not found the task to be updated")
	}
	model.DB.Data.Tasks[updateTask.ID] = updateTask
	return nil
}
