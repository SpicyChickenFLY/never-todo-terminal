package ast

import "fmt"

// help type
const (
	HelpRoot = iota

	HelpTaskList
	HelpTaskAdd
	HelpTaskDelete
	HelpTaskUpdate
	HelpTaskDone
	HelpTaskTodo

	HelpTagList
	HelpTagAdd
	HelpTagDelete
	HelpTagUpdate
)

// help strings
var helpMap = map[int]string{
	HelpRoot: `Usage: nt [Commands]
Commands:
    list - list tasks
    add - add new tasks
    set - update tasks
    done - complete tasks
    todo - uncomplete tasks
    del - delete tasks
    tag - list/add/delete/update tags
    explain - explain how command precessed`,

	HelpTaskList: `Usage(filter mode): nt list {[todo]|done|all}
    [ <content> [and|or <content>] ]
    [ +<tag>|-<tag> [+<tag>|-<tag>] ]
Usage(locate mode): nt list {[todo]|done|all} <id>[-<id>] [<id>]`,
	HelpTaskAdd: `Usage: nt add <content>
    [ +<tag> [+<tag>] ]
    [ due:<due> ]
    [ loop: y|m|w[-SMTWTFS]|d ]`,
	HelpTaskUpdate: `Usage: nt [set] <id> [<content>] 
    [ +<tag>|-<tag> [+<tag>|-<tag>] ]
    [ due:<due> ]
    [ loop: y|m|w[-SMTWTFS]|d ]`,
	HelpTaskDelete: `Usage: nt del <id>[-<id>] [<id>]`,
	HelpTaskTodo:   `Usage: nt todo <id>[-<id>] [<id>]`,
	HelpTaskDone:   `Usage: nt done <id>[-<id>] [<id>]`,

	HelpTagList: `Usage(filter mode): never tag
    [ <content> [and|or <content>] ]
Usage(locate mode): nt tag <id>[-<id>] [<id>]`,
	HelpTagAdd:    `nt tag add [<content>] [color:<color>]`,
	HelpTagUpdate: `nt tag [set] <id> [<content>] [color:<color>]`,
	HelpTagDelete: `nt tag del <id>[-<id>] [<id>]`,
}

// HelpNode show help for specified cmd
type HelpNode struct {
	helpType int
}

// NewHelpNode returns *HelpNode
func NewHelpNode(helpType int) *HelpNode {
	return &HelpNode{helpType}
}

func (hn *HelpNode) execute() {
	fmt.Println(helpMap[hn.helpType])
}

func (hn *HelpNode) explain() string {
	fmt.Println("Show help for command")
	return "nt -h"
}
