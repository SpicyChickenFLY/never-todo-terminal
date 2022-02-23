package ast

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/render"
)

// Command Type
const (
	CMDSummary = 0 + iota
	CMDUI
	CMDExplain
	CMDStmt
	CMDHelp
	CMDNotSupport
)

// export for parser
var (
	Result    *RootNode
	ErrorList = []error{}
	WarnList  = []string{}
	VarMap    = map[string]string{}
)

// RootNode is the root of ast
type RootNode struct {
	CmdType int
	Stmt    StmtNode
}

// NewRootNode return * RootNode
func NewRootNode(cmdType int, sn StmtNode) *RootNode {
	return &RootNode{CmdType: cmdType, Stmt: sn}
}

// // SetCmdType is setter of cmd type
// func (rn *RootNode) SetCmdType(cmdType int) {
// 	rn.cmdType = cmdType
// }

// // SetStmtNode is setter of stmt node
// func (rn *RootNode) SetStmtNode(stmt StmtNode) {
// 	rn.stmtNode = stmt
// }

// Execute should start from root
func (rn *RootNode) Execute(cmd string) {
	switch rn.CmdType {
	case CMDNotSupport:
		// ErrorList = append(ErrorList, errors.New("Command not support"))
		render.Result(cmd, ErrorList, WarnList)
	case CMDSummary:
		if err := model.Init(""); err != nil {
			ErrorList = append(ErrorList, err)
			render.Result(cmd, ErrorList, WarnList)
			return
		}
		if err := model.Begin(); err != nil {
			ErrorList = append(ErrorList, err)
			render.Result(cmd, ErrorList, WarnList)
			return
		}
		render.Summary()
	case CMDHelp:
		rn.Stmt.execute()
	case CMDUI:
		controller.StartUI()
	case CMDExplain:
		if rn.Stmt != nil {
			rn.Explain()
		}
		render.Result(cmd, ErrorList, WarnList)
	case CMDStmt:
		if rn.Stmt != nil {
			if err := model.Init(""); err != nil {
				ErrorList = append(ErrorList, err)
				render.Result(cmd, ErrorList, WarnList)
				return
			}
			if err := model.Begin(); err != nil {
				ErrorList = append(ErrorList, err)
				render.Result(cmd, ErrorList, WarnList)
				return
			}
			rn.Stmt.execute()
			if len(ErrorList) > 0 {
				if err := model.RollBack(); err != nil {
					ErrorList = append(ErrorList, err)
				}
			} else if err := model.Commit(); err != nil {
				ErrorList = append(ErrorList, err)
			}
		}
		render.Result(cmd, ErrorList, WarnList)
	}
}

// Explain should explain from root
func (rn *RootNode) Explain() {
	fmt.Println("==== Execute Plan ====")
	fmt.Println("==== rewrite command ====\n", rn.explain())
}

func (rn *RootNode) explain() string {
	switch rn.CmdType {
	case CMDHelp:
		fmt.Println("show help")
		return "show help"
	case CMDUI:
		fmt.Println("show UI")
		return "show UI"
	case CMDExplain:
		if rn.Stmt != nil {
			return rn.Stmt.explain()
		}
		return "show help"
	case CMDStmt:
		if rn.Stmt != nil {
			return rn.Stmt.explain()
		}
	}
	return ""
}

// Node are all ast nodes
type Node interface {
	explain() string
}

// StmtNode contain a complex statement
type StmtNode interface {
	execute()
	Node
}

// SetVarMap set global var map
func SetVarMap(varMap map[string]string) {
	VarMap = varMap
}

// SearchVarMap convert variable to origin string
func SearchVarMap(str string) string {
	result, ok := VarMap[str]
	if ok {
		str = result
	}
	return str
}
