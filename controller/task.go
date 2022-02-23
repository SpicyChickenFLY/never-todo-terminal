package controller

import (
	"fmt"
	"sort"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/cron"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
)

// ListTasks with filter provided by params
func ListTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// ListTodoTasks with filter provided by params
func ListTodoTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.TaskTodo {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListDoneTasks with filter provided by params
func ListDoneTasks() (tasks []model.Task) {
	for _, task := range model.DB.Data.Tasks {
		if task.Status == model.TaskDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTaskByID called by parser
func GetTaskByID(id int) (model.Task, error) {
	task, ok := model.DB.Data.Tasks[id]
	if !ok {
		return model.Task{}, fmt.Errorf("Retrieve: Task(id:%d) not found", id)
	}
	return task, nil
}

// GetTasksByIDGroup called by parser
func GetTasksByIDGroup(ids []int) (tasks []model.Task, warnList []string) {
	for _, id := range ids {
		task, err := GetTaskByID(id)
		if err == nil {
			tasks = append(tasks, task)
		} else {
			warnList = append(warnList,
				err.Error(),
			)
		}
	}
	return tasks, warnList
}

// DeleteTasks called by parser
func DeleteTasks(ids []int) (warnList []string) {
	for _, id := range ids {
		if _, ok := model.DB.Data.Tasks[id]; ok {
			// task.Status = model.ProjectDeleted
			// model.DB.Data.Tasks[id] = task
			delete(model.DB.Data.Tasks, id)
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Delete: Task(id=%d) not found", id),
			)
		}
	}
	return
}

// CompleteTask called by parse
func CompleteTask(ids []int) (warnList []string) {
	for _, id := range ids {
		if task, ok := model.DB.Data.Tasks[id]; ok {
			task.Status = model.TaskDone
			model.DB.Data.Tasks[id] = task
			if task.Loop != "" {
				p, err := cron.NewPlan(task.Loop)
				if err != nil {
					warnList = append(warnList,
						fmt.Sprintf("Complete: Task(id=%d) loop is invalid because: %s", id, err.Error()))
				} else {
					if task.Due.IsZero() {
						task.Due = time.Now()
					}
					task.Due = p.Next(task.Due)
					task.Status = model.TaskTodo
				}
			}
			model.DB.Data.Tasks[id] = task
		} else {
			warnList = append(warnList,
				fmt.Sprintf("Complete: Task(id=%d) not found", id),
			)
		}
	}
	return
}

// AddTask called by parser
func AddTask(content string) (taskID int) {
	newTask := model.Task{
		ID:      model.DB.Data.TaskInc,
		Content: content,
		Status:  model.TaskTodo,
	}
	model.DB.Data.Tasks[newTask.ID] = newTask
	model.DB.Data.TaskTags[newTask.ID] = make(map[int]bool, 0)
	model.DB.Data.TaskInc--
	return newTask.ID
}

// UpdateTask called by parser
func UpdateTask(updateTask model.Task) error {
	if _, ok := model.DB.Data.Tasks[updateTask.ID]; !ok {
		return fmt.Errorf("Update: Task(id=%d) not found", updateTask.ID)
	}
	if updateTask.Loop != "" {
		p, err := cron.NewPlan(updateTask.Loop)
		if err != nil {
			return fmt.Errorf("Complete: Task(id=%d) loop is invalid because: %s", updateTask.ID, err.Error())
		}
		updateTask.Due = p.Next(time.Now())
	}
	model.DB.Data.Tasks[updateTask.ID] = updateTask
	return nil
}

// SortTask with specified metric
func SortTask(tasks []model.Task, metricName string) []model.Task {
	var lessFunc func(a, b int) bool
	switch metricName {
	case "NAME":
		lessFunc = func(a, b int) bool { return tasks[a].Content > tasks[a].Content }
	case "name":
		lessFunc = func(a, b int) bool { return tasks[a].Content <= tasks[a].Content }
	case "DUE":
		lessFunc = func(a, b int) bool { return !utils.LessInTime(tasks[a].Due, tasks[b].Due) }
	case "due":
		lessFunc = func(a, b int) bool { return utils.LessInTime(tasks[a].Due, tasks[b].Due) }
	case "ID":
		lessFunc = func(a, b int) bool { return !utils.LessInID(tasks[a].ID, tasks[b].ID) }
	default:
		lessFunc = func(a, b int) bool { return utils.LessInID(tasks[a].ID, tasks[b].ID) }
	}
	sort.SliceStable(tasks, lessFunc)
	return tasks
}
