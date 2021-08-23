package controller

import (
	"errors"
	"fmt"

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

// FindTagByID called by parser
func FindTagByID(id int) (model.Tag, bool) {
	tag, ok := model.DB.Data.Tags[id]
	return tag, ok
}

// GetTagIDByName called by parser
func GetTagIDByName(string) (int, bool) {
	return 0, false
}

// AddTag called by parser
func AddTag(content string) (int, error) {
	id, ok := GetTagIDByName(content)
	if ok {
		return id, errors.New("tag already exists")
	}
	newTag := model.Tag{
		ID:      model.DB.Data.TagInc,
		Content: content,
		Color:   "white",
	}
	model.DB.Data.Tags[model.DB.Data.TagInc] = newTag
	model.DB.Data.TagInc--

	return newTag.ID, nil
}

// UpdateTag called by parser
func UpdateTag(updateTag model.Tag) error {
	if _, ok := model.DB.Data.Tags[updateTag.ID]; !ok {
		return errors.New(fmt.Sprint("tag(id=%d) not found", updateTag.ID))
	}
	model.DB.Data.Tags[updateTag.ID] = updateTag
	return nil
}

// DeleteTags called by parser
func DeleteTags(ids []int) (warnList []string) {
	// delete tag
	for _, id := range ids {
		if deleteTag, ok := model.DB.Data.Tags[id]; !ok {
			warnList = append(warnList,
				fmt.Sprint("Task(id=%d) not found", id),
			)
		} else {
			deleteTag.Deleted = true
			model.DB.Data.Tags[id] = deleteTag
		}
	}
	return
}
