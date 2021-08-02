package ast

type TagListNode struct {
}

func NewTagListNode() TagListNode {
	return TagListNode{}
}

type TagListFilterNode struct {
}

func NewTagListFilterNode(ign *IDGroupNode, cgn *ContentGroupNode) *TagListFilterNode {
	return &TagListFilterNode{}
}
