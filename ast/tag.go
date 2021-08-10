package ast

import "fmt"

type TagListNode struct {
	tagListFilterNode *TagListFilterNode
}

func NewTagListNode(tlfn *TagListFilterNode) *TagListNode {
	return &TagListNode{tlfn}
}

func (tln *TagListNode) execute() error { return nil }
func (tln *TagListNode) explain() string {
	fmt.Println("list tag ")
	// tln.taskListFilterNode.explain()
	return "tag list " // + tln.taskListFilterNode.restore()
}

type TagListFilterNode struct {
}

func NewTagListFilterNode(ign *IDGroupNode, cgn *ContentGroupNode) *TagListFilterNode {
	return &TagListFilterNode{}
}
