package model

import (
	"time"
)

type FullTask struct {
	Task Task
	Tags []Tag
}

type Task struct {
	ID        int       `json:"id" mapstructure:"id"`
	Content   string    `json:"content" mapstructure:"content"`
	Deleted   bool      `json:"deleted" mapstructure:"deleted"`
	Completed bool      `json:"completed" mapstructure:"completed"`
	Important int       `json:"important" mapstructure:"important"`
	Due       time.Time `json:"due,omitempty" mapstructure:"due"`
	Loop      string    `json:"loop,omitempty" mapstructure:"loop"`
}

type Tag struct {
	ID      int    `json:"id" mapstructure:"id"`
	Content string `json:"content" mapstructure:"content"`
	Deleted bool   `json:"deleted" mapstructure:"deleted"`
	Color   string `json:"color" mapstructure:"color"`
}

type TaskTag struct {
	TaskID int `json:"task_id" mapstructure:"task_id"`
	TagID  int `json:"tag_id" mapstructure:"tag_id"`
}

type TimeGroup struct {
	Level int
	Start time.Time
	End   time.Time
}

type Loop struct {
	Year   int
	Month  int
	Week   int
	Hour   int
	Minute int
}

type Log struct {
	Target string                 `json:"target" mapstructure:"target"`
	Type   string                 `json:"type" mapstructure:"type"`
	Data   map[string]interface{} `json:"data" mapstructure:"data"`
}

type Data struct {
	Tasks          []Task    `json:"tasks" mapstructure:"tasks"`
	Tags           []Tag     `json:"tags" mapstructure:"tags"`
	TaskTags       []TaskTag `json:"task_tags" mapstructure:"task_tags"`
	TaskAutoIncVal int       `json:"taskAutoIncVal" mapstructure:"taskAutoIncVal"`
	TagAutoIncVal  int       `json:"tagAutoIncVal" mapstructure:"tagAutoIncVal"`
}

type Model struct {
	Data Data  `json:"data", mapstructure:"data"`
	Log  []Log `json:"log" mapstructure:"log"`
}
