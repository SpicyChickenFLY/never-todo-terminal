package controller

import "github.com/SpicyChickenFLY/never-todo-cmd/data"

var (
	model *data.Model
	db    *data.DB
)

func init() {
	db = data.NewDB()
	model = &data.Model{}
}
