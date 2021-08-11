package ast

import "fmt"

// TagListNode include tag list filter
type TagListNode struct {
	tagListFilterNode *TagListFilterNode
}

// NewTagListNode return *TagListNode
func NewTagListNode(tlfn *TagListFilterNode) *TagListNode {
	return &TagListNode{tlfn}
}

func (tln *TagListNode) execute() {}
func (tln *TagListNode) explain() string {
	fmt.Println("list tag ")
	// tln.taskListFilterNode.explain()
	return "tag list " // + tln.taskListFilterNode.restore()
}

// TagListFilterNode include content
type TagListFilterNode struct {
}

// NewTagListFilterNode return *TagListFilterNode
func NewTagListFilterNode(ign *IDGroupNode, content string) *TagListFilterNode {
	return &TagListFilterNode{}
}
