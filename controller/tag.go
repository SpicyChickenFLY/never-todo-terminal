package controller

import (
	"errors"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ListTags with filter provided by params
func ListTags() (tags []model.Tag) {
	for _, tag := range model.DB.Data.Tags {
		if !tag.Deleted {
			tags = append(tags, tag)
		}
	}
	return tags
}

// AddTag called by parser
func AddTag(content string) (int, error) {
	id, ok := GetTagIDByName(content)
	if ok {
		return id, errors.New("Tag already exists")
	}
	newTag := model.Tag{
		ID:      model.DB.Data.TagAutoIncVal,
		Content: content,
	}
	model.DB.Data.Tags = append(model.DB.Data.Tags, newTag)
	model.DB.Data.TagAutoIncVal--

	return newTag.ID, nil
}

// DelTag called by parser
func DelTag(id int) error {
	return nil
}

// SetTag called by parser
func SetTag(tag model.Tag) {

}

// FindTagByID called by parser
func FindTagByID(id int) (model.Tag, bool) {
	for _, tag := range model.DB.Data.Tags {
		if tag.ID == id {
			return tag, true
		}
	}
	return model.Tag{}, false
}

// GetTagIDByName called by parser
func GetTagIDByName(string) (int, bool) {
	return 0, false
}
