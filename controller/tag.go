package controller

import (
	"errors"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ShowTags called by parser
func ShowTags() {}

// AddTag called by parser
func AddTag(content string, color string) (int, error) {
	id, ok := GetTagIDByName(content)
	if ok {
		return id, errors.New("")
	}
	model.M.Data.Tags = append(model.M.Data.Tags, model.Tag{Content: content, Color: color})
	return 0, nil
}

// DelTag called by parser
func DelTag(id int) error {
	return nil
}

// SetTag called by parser
func SetTag() {}

// FindTagByID called by parser
func FindTagByID() {}

// GetTagIDByName called by parser
func GetTagIDByName(string) (int, bool) {
	return 0, false
}
