package controller

import (
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

func FindTagsByTask(taskID int) (tags []model.Tag) {
	if tagsID, ok := model.DB.Data.TaskTags[taskID]; ok {
		for _, tagID := range tagsID {
			if tag, ok := FindTagByID(tagID); ok {
				tags = append(tags, tag)
			}
		}
	}
	return
}

// AddTaskTags called by parser
func AddTaskTags(taskID int, assignTags []string) (err error) {
	for _, assignTag := range assignTags {
		var tagID int
		tagID, ok := GetTagIDByName(assignTag)
		if !ok {
			tagID, err = AddTag(assignTag)
			if err != nil {
				return err
			}
		}
		if !checkTaskTagExist(taskID, tagID) {
			model.DB.Data.TaskTags = append(model.DB.Data.TaskTags, model.TaskTag{TaskID: taskID, TagID: tagID})
		}
	}
	return nil
}

// DeleteTaskTags called by parse
func DeleteTaskTags(taskID int, unassignTags []string) error {
	for _, unassignTag := range unassignTags {
		tagID, _ := GetTagIDByName(unassignTag)
		if err := deleteTaskTag(taskID, tagID); err != nil {
			return err
		}
	}
	return nil
}

func checkTaskTagExist(taskID, tagID int) bool {
	for _, taskTag := range model.DB.Data.TaskTags {
		if taskTag.TaskID == taskID && taskTag.TagID == tagID {
			return true
		}
	}
	return false
}

// CheckTaskByTags filter tasks
func CheckTaskByTags(taskID int, tagIDs []int) bool {
	for _, tagID := range tagIDs {
		if !checkTaskTagExist(taskID, tagID) {
			return false
		}
	}
	return true
}

func deleteTaskTag(taskID, tagID int) error {
	index := 0
	for i, taskTag := range model.DB.Data.TaskTags {
		if taskTag.TaskID == taskID && taskTag.TagID == tagID {
			index = i
			break
		}
	}
	if index != 0 {
		// FIXME: 会不会越界
		model.DB.Data.TaskTags = append(model.DB.Data.TaskTags[:index], model.DB.Data.TaskTags[index+1:]...)
	}
	return nil
}
