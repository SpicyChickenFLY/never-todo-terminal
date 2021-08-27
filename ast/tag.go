package ast

import (
	"errors"
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

func (tan *TagAddNode) execute() {
	tagID, err := controller.AddTag(tan.content)
	if err != nil {
		ErrorList = append(ErrorList, err)
	}
	tag, ok := controller.GetTagByID(tagID)
	if !ok {
		WarnList = append(WarnList, fmt.Sprintf("task(%d) not found", tagID))
	}
	controller.UpdateTag(tag)
	render.Tags([]model.Tag{tag})
}
func (tan *TagAddNode) explain() string {
	result := "tag add "
	fmt.Println("Add new tag")
	fmt.Printf("\twith content `%s`\n", tan.content)
	result += fmt.Sprintf("`%s` ", tan.content)
	if tan.color != "" {
		fmt.Printf("\twith color `%s`\n", tan.color)
		result += fmt.Sprintf("#%s ", tan.color)
	}
	return result
}

// TagUpdateNode include tag list filter
type TagUpdateNode struct {
	id     int
	option *TagUpdateOptionNode
}

// NewTagUpdateNode return *TagUpdateNode
func NewTagUpdateNode(id int, tuon *TagUpdateOptionNode) *TagUpdateNode {
	return &TagUpdateNode{id, tuon}
}

func (tun *TagUpdateNode) execute() {
	tag, ok := controller.GetTagByID(tun.id)
	if !ok {
		ErrorList = append(ErrorList, errors.New("updated tag is not found"))
		return
	}
	tun.option.apply(tag)

}
func (tun *TagUpdateNode) explain() string {
	result := fmt.Sprintf("tag set %d ", tun.id)
	fmt.Printf("Update tag:%d ", tun.id)
	result += tun.option.explain()
	return result
}

// ============================
// Tag Update Option
// ============================

// TagUpdateOptionNode is node for tag update option
type TagUpdateOptionNode struct {
	content string
	color   string
}

// NewTagUpdateOptionNode return *TagUpdateOptionNode
func NewTagUpdateOptionNode() *TagUpdateOptionNode {
	return &TagUpdateOptionNode{}
}

func (tuon *TagUpdateOptionNode) explain() string {
	result := ""
	if tuon.content != "" {
		fmt.Printf("\tset content:%s", tuon.content)
		result += fmt.Sprintf("`%s` ", tuon.content)
	}
	if tuon.content != "" {
		fmt.Printf("\tset color:%s", tuon.color)
		result += fmt.Sprintf("`%s` ", tuon.color)
	}
	return result
}

func (tuon *TagUpdateOptionNode) apply(tag model.Tag) {
	if tuon.content != "" {
		tag.Content = tuon.content
	}
	if tuon.color != "" {
		tag.Color = tuon.color
	}
	if err := controller.UpdateTag(tag); err != nil {
		ErrorList = append(ErrorList, err)
	}
	render.Tags([]model.Tag{tag})
}

// SetContent for TagUpdateOptionNode
func (tuon *TagUpdateOptionNode) SetContent(content string) *TagUpdateOptionNode {
	tuon.content = content
	return tuon
}

// SetColor for TagUpdateOptionNode
func (tuon *TagUpdateOptionNode) SetColor(color string) *TagUpdateOptionNode {
	tuon.color = color
	return tuon
}

// ============================
// Tag Delete
// ============================

// TagDeleteNode is node for delete tag
type TagDeleteNode struct {
	idGroup IDGroupNode
}

// NewTagDeleteNode return TagDeleteNode
func NewTagDeleteNode(ign *IDGroupNode) *TagDeleteNode {
	return &TagDeleteNode{*ign}
}

func (tdn *TagDeleteNode) explain() string {
	result := "tag del "
	fmt.Println("delete tag ")
	result += tdn.idGroup.explain()
	return result
}

func (tdn *TagDeleteNode) execute() {
	controller.DeleteTags(tdn.idGroup.ids)
}
