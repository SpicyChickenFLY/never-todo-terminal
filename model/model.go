package model

import (
	"time"
)

// Task is task struct
type Task struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Important int    `json:"important"`
	ProjectID int
	Due       time.Time `json:"due,omitempty"`
	Loop      string    `json:"loop,omitempty"`
}

// Tag is tag struct
type Tag struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Deleted bool   `json:"deleted"`
	Color   string `json:"color"`
}

// TaskTag is tasktag struct
type TaskTag struct {
	TaskID int `json:"task_id"`
	TagID  int `json:"tag_id"`
}

// Project default ID
const (
	ProjectTodo = iota
	ProjectDone
	ProjectDeleted
)

// Project is project struct
type Project struct {
	ID      int
	Content string
}

// LogType
const (
	LogTypeCreate = iota
	LogTypeUpdate
	logTypeDelete
)

// Log is log struct
type Log struct {
	Table string
	ID    int
	Type  int
	Data  []interface{}
}

// Data is data struct
type Data struct {
	Tasks      map[int]Task
	Tags       map[int]Tag
	TaskTags   map[int]map[int]bool
	TagTasks   map[int]map[int]bool
	Projects   map[int]Project
	TaskInc    int
	TagInc     int
	ProjectInc int
}

//Model is model struct
type Model struct {
	Data Data  `json:"data"`
	Logs []Log `json:"log"`
}

// TimeGroup is timegroup struct
type TimeGroup struct {
	Level int
	Start time.Time
	End   time.Time
}

// Loop is loop struct
type Loop struct {
	Year   int
	Month  int
	Week   int
	Hour   int
	Minute int
}
