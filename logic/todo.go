package logic

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/data"
)

var (
	model *data.Model
	db    *data.DB
)

func init() {
	db = data.NewDB()
	model = &data.Model{}
}

func ShowSummary() error {
	if err := db.Read(model); err != nil {
		return err
	}
	var todoTotal, doneTotal, tagTotal int
	for _, task := range model.Data.Tasks {
		if !task.Deleted {
			if !task.Completed {
				todoTotal++
			} else {
				doneTotal++
			}
		}
	}
	for _, tag := range model.Data.Tags {
		if !tag.Deleted {
			tagTotal++
		}
	}
	fmt.Printf("todo: %d, done: %d, tag: %d\n", todoTotal, doneTotal, tagTotal)

	return nil
}

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
	return db.Write(model)
}

func AddTasks(params []string) error {
	if err := db.Read(model); err != nil {
		return err
	}
	for _, param := range params {
		fmt.Println(param)
	}
	// add new task
	return db.Write(model)
}
func DelTasks(params []string) error {
	if err := db.Read(model); err != nil {
		return err
	}
	for _, param := range params {
		fmt.Println(param)
	}
	// delete task
	return db.Write(model)
}
func SetTask() {}

func ShowTags() {}
func AddTags()  {}
func DelTags()  {}
func SetTags()  {}

func AddLinks() {}
func DelLinks() {}
func SetLinks() {}
