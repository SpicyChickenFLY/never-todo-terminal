package ast

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/render"
)

// Command Type
const (
	CMDSummary = 0 + iota
	CMDUI
	CMDExplain
	CMDStmt
	CMDHelp
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
	cmdType  int
	stmtNode StmtNode
}

// NewRootNode return * RootNode
func NewRootNode(cmdType int, sn StmtNode) *RootNode {
	return &RootNode{cmdType: cmdType, stmtNode: sn}
}

// Execute should start from root
func (rn *RootNode) Execute(cmd string) error {
	switch rn.cmdType {
	case CMDHelp:
		return controller.ShowHelp()
	case CMDUI:
		return controller.StartUI()
	case CMDExplain:
		if rn.stmtNode != nil {
			fmt.Println("==== Execute Plan ====")
			fmt.Println("==== rewrite ====\n", rn.explain())
		}
		return nil
	case CMDStmt:
		rn.stmtNode.execute()
		render.Result(cmd, ErrorList, WarnList)

		return nil
	default:
		return errors.New("目前不支持的命令类型")
	}
}

// Explain should explain from root
func (rn *RootNode) Explain() {
	fmt.Println("==== Execute Plan ====")
	fmt.Println("==== rewrite command ====\n", rn.explain())

}

func (rn *RootNode) explain() string {
	switch rn.cmdType {
	case CMDHelp:
		fmt.Println("show help")
		return "show help"
	case CMDUI:
		fmt.Println("show UI")
		return "show UI"
	case CMDExplain:
		if rn.stmtNode != nil {
			return rn.stmtNode.explain()
		}
		return "show help"
	case CMDStmt:
		if rn.stmtNode != nil {
			return rn.stmtNode.explain()
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
