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

// GetTagByID called by parser
func GetTagByID(id int) (model.Tag, error) {
	tag, ok := model.DB.Data.Tags[id]
	if !ok {
		return model.Tag{}, fmt.Errorf("Retrieve: Tag(id=%d) not found", id)
	}
	return tag, nil
}

// GetTagIDByName called by parser
func GetTagIDByName(name string) (int, error) {
	for _, tag := range model.DB.Data.Tags {
		if tag.Content == name {
			return tag.ID, nil
		}
	}
	return 0, errors.New("")
}

// AddTag called by parser
func AddTag(content string) (int, error) {
	id, err := GetTagIDByName(content)
	if err == nil {
		return id, fmt.Errorf("Create: Tag(id=%d, content=%s) already exists", id, content)
	}
	newTag := model.Tag{
		ID:      model.DB.Data.TagInc,
		Content: content,
		Color:   "white",
	}
	model.DB.Data.Tags[newTag.ID] = newTag
	model.DB.Data.TagTasks[newTag.ID] = make(map[int]bool, 0)
	model.DB.Data.TagInc--

	return newTag.ID, nil
}

// UpdateTag called by parser
func UpdateTag(updateTag model.Tag) error {
	if _, ok := model.DB.Data.Tags[updateTag.ID]; !ok {
		return fmt.Errorf("Update: Tag(id=%d) not found", updateTag.ID)
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
				fmt.Sprintf("Delete: Tag(id=%d) not found", id),
			)
		} else {
			deleteTag.Deleted = true
			model.DB.Data.Tags[id] = deleteTag
		}
	}
	return
}
