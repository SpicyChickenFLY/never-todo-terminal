package logic

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/constant"
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

// ShowSummary like the total amount of todo/tag/schedule
func ShowSummary() error {
	// 展示logo，用法，当前待办数和标签数
	fmt.Println(constant.ColorfulLogo)
	fmt.Println(constant.Descirption)
	fmt.Println(constant.Separator)
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

func AddTasks(assign_tags []data.Tag, dueTime time.Time, loop string) error {
	if err := db.Read(model); err != nil {
		return err
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

func SetTask() error {
	if err := db.Read(model); err != nil {
		return err
	}
	return db.Write(model)
}

func ShowTags() {}
func AddTags()  {}
func DelTags()  {}
func SetTags()  {}

func AddLinks() {}
func DelLinks() {}
func SetLinks() {}
