package ast

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/render"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

// TagListNode include tag list filter
type TagListNode struct {
	tagListFilter *TagListFilterNode
}

// NewTagListNode return *TagListNode
func NewTagListNode(tlfn *TagListFilterNode) *TagListNode {
	return &TagListNode{tlfn}
}

func (tln *TagListNode) execute() {
	tags := controller.ListTags()
	render.Tags(tln.tagListFilter.filter(tags))
}
func (tln *TagListNode) explain() string {
	fmt.Println("list tag ")
	// tln.taskListFilterNode.explain()
	return "tag list " // + tln.taskListFilterNode.restore()
}

// TagListFilterNode include content
type TagListFilterNode struct {
	idGroup *IDGroupNode
	content string
}

// NewTagListFilterNode return *TagListFilterNode
func NewTagListFilterNode(ign *IDGroupNode, content string) *TagListFilterNode {
	return &TagListFilterNode{ign, content}
}

func (tlfn *TagListFilterNode) filter(tags []model.Tag) []model.Tag {
	if tlfn.idGroup != nil {
		result := []model.Tag{}
		for _, id := range tlfn.idGroup.ids {
			for _, tag := range tags {
				if tag.ID == id {
					result = append(result, tag)
				}
			}
		}
		return result
	} else if tlfn.content != "" {
		result := []model.Tag{}
		for _, tag := range tags {
			if utils.ContainStr(tag.Content, tlfn.content) {
				result = append(result, tag)
			}
		}
		return result
	} else {
		return tags
	}
}

// TagAddNode include tag list filter
type TagAddNode struct {
	content string
	color   string
}

// NewTagAddNode return *TagAddNode
func NewTagAddNode(content, color string) *TagAddNode {
	return &TagAddNode{content, color}
}

func (tln *TagAddNode) execute() {
	tagID, err := controller.AddTag(tln.content)
	if err != nil {
		ErrorList = append(ErrorList, err)
	}
	tag, ok := controller.FindTagByID(tagID)
	if !ok {
	}
	controller.SetTag(tag)
	render.Tags([]model.Tag{tag})
}
func (tln *TagAddNode) explain() string {

}

// TagUpdatedNode include tag list filter
type TagUpdatedNode struct {
	id      int
	content string
	color   string
}

// NewTagUpdatedNode return *TagUpdatedNode
func NewTagUpdatedNode(id int, content, color string) *TagUpdatedNode {
	return &TagUpdatedNode{id, content, color}
}

func (tln *TagUpdatedNode) execute() {

}
func (tln *TagUpdatedNode) explain() string {

}
