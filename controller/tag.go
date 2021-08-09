package controller

import (
	"errors"

	"github.com/SpicyChickenFLY/never-todo-cmd/data"
)

// ShowTags called by parser
func ShowTags() {}

// AddTag called by parser
func AddTag(content string, color string) (int, error) {
	if err := db.Read(model); err != nil {
		return 0, err
	}
	id, ok := GetTagIDByName(content)
	if ok {
		return id, errors.New("")
	}
	model.Data.Tags = append(model.Data.Tags, data.Tag{Content: content, Color: color})
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
