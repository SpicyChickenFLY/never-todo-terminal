package controller

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ListTasks with filter provided by params
func ListTasks(params []string) error {
	if len(params) > 0 { // search

	} else { //list all
		for _, task := range model.M.Data.Tasks {
			if !task.Deleted && !task.Completed {
				fmt.Println(task.Content)
			}
		}
	}
	return nil
}

// FindTaskByID called by parser
func FindTaskByID(id int) (model.Task, bool, error) {
	for _, task := range model.M.Data.Tasks {
		if task.ID == id {
			return task, true, nil
		}
	}
	return model.Task{}, false, nil
}

// DeleteTask called by parser
func DeleteTask(ids []int) error {
	// delete task
	for _, id := range ids {
		for _, task := range model.M.Data.Tasks {
			if task.ID == id {
				task.Deleted = true
			}
		}
	}
	return nil
}

// CompleteTask called by parse
func CompleteTask(ids []int) error {
	for _, id := range ids {
		for _, task := range model.M.Data.Tasks {
			if task.ID == id {
				task.Completed = true
			}
		}
	}
	return nil
}

// AddTask called by parser
func AddTask(
	content string,
	importance int,
	assignTags []int,
	due time.Time,
	loop string) (int, error) {
	// add new task
	return 0, nil
}

// UpdateTask called by parser
func UpdateTask(
	id int,
	content string,
	importance int,
	assignTags, unasssignTags []string,
	due time.Time) error {
	for _, task := range model.M.Data.Tasks {
		if task.ID == id {
			task.Content = content
			task.Important = (importance > 0)
			AddTaskTags(task.ID, assignTags)
			DeleteTaskTags(task.ID, unasssignTags)
			task.Due = due
			// TODO: Set Loop
			// task.Loop = loop
		}
	}
	return nil
}
