package controller

import (
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// FindTagsByTask return tags by task id
func FindTagsByTask(taskID int) (tags []model.Tag, warnList []string) {
	if tagMap, ok := model.DB.Data.TaskTags[taskID]; ok {
		for tagID := range tagMap {
			if tag, err := GetTagByID(tagID); err == nil {
				tags = append(tags, tag)
			} else {
				warnList = append(warnList, err.Error())
			}
		}
	}
	return
}

// FindTasksByTag return tasks by tag id
func FindTasksByTag(tagID int) (tasks []model.Task, warnList []string) {
	if taskMap, ok := model.DB.Data.TagTasks[tagID]; ok {
		for taskID := range taskMap {
			if task, err := GetTaskByID(taskID); err == nil {
				tasks = append(tasks, task)
			} else {
				warnList = append(warnList, err.Error())
			}
		}
	}
	return
}

// AddTaskTags called by parser
func AddTaskTags(taskID int, assignTags []string) (err error) {
	for _, assignTag := range assignTags {
		var tagID int
		tagID, err := GetTagIDByName(assignTag)
		if err != nil {
			tagID, err = AddTag(assignTag)
			if err != nil {
				return err
			}
		}
		if !checkTaskTagExist(taskID, tagID) {
			if _, ok := model.DB.Data.TaskTags[taskID]; !ok {
				model.DB.Data.TaskTags[taskID] = make(map[int]bool)
			}
			model.DB.Data.TaskTags[taskID][tagID] = true
			if _, ok := model.DB.Data.TagTasks[tagID]; !ok {
				model.DB.Data.TagTasks[tagID] = make(map[int]bool)
			}
			model.DB.Data.TagTasks[tagID][taskID] = true
		}
	}
	return nil
}

// DeleteTaskTags called by parse
func DeleteTaskTags(taskID int, unassignTags []string) error {
	for _, unassignTag := range unassignTags {
		tagID, err := GetTagIDByName(unassignTag)
		if err != nil {
			return err
		}
		if _, ok := model.DB.Data.TaskTags[taskID]; ok {
			delete(model.DB.Data.TaskTags[taskID], tagID)
		}
		if _, ok := model.DB.Data.TagTasks[tagID]; ok {
			delete(model.DB.Data.TagTasks[tagID], taskID)
		}
	}
	return nil
}

func checkTaskTagExist(taskID, tagID int) bool {
	if tagMap, ok := model.DB.Data.TaskTags[taskID]; ok {
		if _, ok := tagMap[tagID]; ok {
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
