package controller

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ListProjects with filter provided by params
func ListProjects() (projects []model.Project) {
	for _, project := range model.DB.Data.Projects {
		if !project.Deleted {
			projects = append(projects, project)
		}
	}
	return projects
}

// GetProjectByID called by parser
func GetProjectByID(id int) (model.Project, bool) {
	project, ok := model.DB.Data.Projects[id]
	return project, ok
}

// GetProjectIDByName called by parser
func GetProjectIDByName(name string) (int, bool) {
	for _, project := range model.DB.Data.Projects {
		if project.Content == name {
			return project.ID, true
		}
	}
	return 0, false
}

// AddProject called by parser
func AddProject(content string) (int, error) {
	id, ok := GetProjectIDByName(content)
	if ok {
		return id, errors.New("project already exists")
	}
	newProject := model.Project{
		ID:      model.DB.Data.ProjectInc,
		Content: content,
		Color:   "white",
	}
	model.DB.Data.Projects[model.DB.Data.ProjectInc] = newProject
	model.DB.Data.ProjectInc--

	return newProject.ID, nil
}

// UpdateProject called by parser
func UpdateProject(updateProject model.Project) error {
	if _, ok := model.DB.Data.Projects[updateProject.ID]; !ok {
		return fmt.Errorf("project(id=%d) not found", updateProject.ID)
	}
	model.DB.Data.Projects[updateProject.ID] = updateProject
	return nil
}

// DeleteProjects called by parser
func DeleteProjects(ids []int) (warnList []string) {
	// delete project
	for _, id := range ids {
		if deleteProject, ok := model.DB.Data.Projects[id]; !ok {
			warnList = append(warnList,
				fmt.Sprintf("Task(id=%d) not found", id),
			)
		} else {
			deleteProject.Deleted = true
			model.DB.Data.Projects[id] = deleteProject
		}
	}
	return
}
