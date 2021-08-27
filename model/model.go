package model

import (
	"time"
)

// Task is task struct
type Task struct {
	ID        int
	Content   string
	Important int
	ProjectID int
	Due       time.Time
	Loop      string
}

// Tag is tag struct
type Tag struct {
	ID      int
	Content string
	Deleted bool
	Color   string
}

// TaskTag is tasktag struct
type TaskTag struct {
	TaskID int
	TagID  int
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
	Color   string
	Deleted bool
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
	Data Data
	Logs []Log
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
