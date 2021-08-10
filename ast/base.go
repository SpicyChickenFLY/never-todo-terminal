package ast

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
)

const ( // Command Type
	// CMDSummary 0
	CMDSummary = 0 + iota
	// CMDUI 1
	CMDUI
	// CMDGUI 2
	CMDGUI
	// CMDExplain 3
	CMDExplain
	// CMDStmt 4
	CMDStmt
	// CMDHelp 5
	CMDHelp
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
			fmt.Println("rewrite command: ", rn.stmtNode.restore())
			fmt.Println("==== Execute Plan ====")
			rn.stmtNode.explain()
		}
		return nil
	case CMDStmt:
		return rn.stmtNode.execute()
	default:
		return errors.New("目前不支持的命令类型")
	}
}

// Explain should explain from root
func (rn *RootNode) Explain() {
	rn.explain()
}

func (rn *RootNode) explain() {
	switch rn.cmdType {
	case CMDSummary:
		fmt.Println("show summary")
	case CMDHelp:
		fmt.Println("show help")
	case CMDUI:
		fmt.Println("show UI")
	case CMDGUI:
		fmt.Println("show GUI")
	case CMDExplain:
		fmt.Println("show explaination")

	case CMDStmt:
		if rn.stmtNode != nil {
			rn.stmtNode.explain()
		}
	}
}

func (rn *RootNode) restore() string {
	switch rn.cmdType {
	case CMDSummary:
		return "show summary"
	case CMDHelp:
		return "show help"
	case CMDUI:
		return "show UI"
	case CMDGUI:
		return "show GUI"
	case CMDExplain:
		return "show explaination"

	case CMDStmt:
		if rn.stmtNode != nil {
			rn.stmtNode.restore()
		}
	}
	return ""
}

// Node are all ast nodes
type Node interface {
	explain()
	restore() string
}

// StmtNode contain a complex statement
type StmtNode interface {
	execute() error
	Node
}
