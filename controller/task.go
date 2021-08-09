package controller

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/data"
)

// ShowTasks with filter provided by params
func ShowTasks(params []string) error {
	if err := db.Read(model); err != nil {
		return err
	}
	if len(params) > 0 { // search

	} else { //list all
		for _, task := range model.Data.Tasks {
			if !task.Deleted && !task.Completed {
				fmt.Println(task.Content)
			}
		}
	}
	return nil
}

// FindTaskByID called by parser
func FindTaskByID(id int) (data.Task, bool, error) {
	if err := db.Read(model); err != nil {
		return data.Task{}, false, err
	}
	for _, task := range model.Data.Tasks {
		if task.ID == id {
			return task, true, nil
		}
	}
	return data.Task{}, false, nil
}

// DeleteTask called by parser
func DeleteTask(ids []int) error {
	if err := db.Read(model); err != nil {
		return err
	}
	// delete task
	for _, id := range ids {
		for _, task := range model.Data.Tasks {
			if task.ID == id {
				task.Deleted = true
			}
		}
	}
	return db.Write(model)
}

// CompleteTask called by parse
func CompleteTask(ids []int) error {
	if err := db.Read(model); err != nil {
		return err
	}
	for _, id := range ids {
		for _, task := range model.Data.Tasks {
			if task.ID == id {
				task.Completed = true
			}
		}
	}
	return db.Write(model)
}

// AddTask called by parser
func AddTask(
	content string,
	importance int,
	assignTags []int,
	due time.Time,
	loop string) (int, error) {
	if err := db.Read(model); err != nil {
		return 0, err
	}
	// add new task
	return 0, db.Write(model)
}

// UpdateTask called by parser
func UpdateTask(
	id int,
	content string,
	importance int,
	assignTags, unasssignTags []string,
	due time.Time) error {
	if err := db.Read(model); err != nil {
		return err
	}
	for _, task := range model.Data.Tasks {
		if task.ID == id {
			task.Content = content
			task.Important = (importance > 0)
			AddTaskTags(task.ID, assignTags)
			DeleteTaskTags(task.ID, unasssignTags)
			// TODO: Set due
			// TODO: Set Loop
		}
	}
	return db.Write(model)
}
