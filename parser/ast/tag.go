package ast

type TagListFilterNode struct {
}

func NewTagListFilterNode(ign *IDGroupNode, cgn *ContentGroupNode) *TagListFilterNode {
	return &TagListFilterNode{}
}
