package ast

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
)

// Command Type
const (
	CMDSummary = 0 + iota
	CMDUI
	CMDGUI
	CMDExplain
	CMDStmt
	CMDHelp
)

// export for parser
var (
	Result    *RootNode
	ErrorList = []error{}
	WarnList  = []string{}
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
func (rn *RootNode) Execute() error {
	fmt.Println(rn.cmdType)
	switch rn.cmdType {
	case CMDSummary:
		return controller.ShowSummary()
	case CMDHelp:
		return controller.ShowHelp()
	case CMDUI:
		return controller.StartUI()
	case CMDGUI:
		return controller.StartGUI()
	case CMDExplain:
		if rn.stmtNode != nil {
			fmt.Println("==== Execute Plan ====")
			fmt.Println("rewrite command: ", rn.explain())
		}
		return nil
	case CMDStmt:
		rn.stmtNode.execute()
		// TODO: handle error list
		return errors.New("语句执行失败")
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
	case CMDSummary:
		fmt.Println("show summary")
		return "show summary"
	case CMDHelp:
		fmt.Println("show help")
		return "show help"
	case CMDUI:
		fmt.Println("show UI")
		return "show UI"
	case CMDGUI:
		fmt.Println("show GUI")
		return "show GUI"
	case CMDExplain:
		fmt.Println("show explaination")
		return "show explaination"
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
