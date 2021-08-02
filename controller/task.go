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

func AddTasks(assign_tags []data.Tag, dueTime time.Time, loop string) error {
	if err := db.Read(model); err != nil {
		return err
	}
	// add new task
	return db.Write(model)
}

func UpdateTask(id int, content string) error {
	if err := db.Read(model); err != nil {
		return err
	}
	for _, task := range model.Data.Tasks {
		if task.ID == id {
			task.Content = content
			// TODO: update Option
			task.Important = false
		}
	}
	return db.Write(model)
}
