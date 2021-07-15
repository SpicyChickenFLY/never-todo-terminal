package ast

import (
	"errors"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
)

const ( // Command Type
	CMDSummary = 0 + iota
	CMDUI
	CMDGUI
	CMDExplain
	CMDStmt
	CMDHelp
)

// Node is node of abstract syntax tree
type Node interface {
	Execute() error
}

// RootNode is the root of ast
type RootNode struct {
	cmdType  int
	stmtNode *StmtNode
}

// NewRootNode return * RootNode
func NewRootNode(cmdType int, sn *StmtNode) *RootNode {
	return &RootNode{cmdType: cmdType, stmtNode: sn}
}

// Execute should start from root
func (rn *RootNode) Execute() error {
	switch rn.getCmdType() {
	case CMDSummary:
		return controller.ShowSummary()
	case CMDHelp:
		return controller.ShowHelp()
	case CMDUI:
		return controller.StartUI()
	case CMDGUI:
		return controller.StartGUI()
	case CMDExplain:
		(*(rn.stmtNode)).explain()
		return nil
	case CMDStmt:
		return (*(rn.stmtNode)).execute()
	default:
		return errors.New("目前不支持的命令类型")
	}
}

func (rn *RootNode) getCmdType() int { return rn.cmdType }

// StmtNode contain a complex statement
type StmtNode interface {
	execute() error
	explain()
}
