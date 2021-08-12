package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/jedib0t/go-pretty/v6/table"
)

var t table.Writer

func init() {
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
}

func Tasks(tasks []model.Task) (warnList []string) {
	t.AppendHeader(table.Row{"#", "Content", "Tags", "Due", "Loop"})
	for _, task := range tasks {
		contentStr := task.Content
		if task.Important {
			contentStr = Str(contentStr, "default", "yellow", "white")
		}
		dueStr := task.Due.Format("2006/01/02 15:04:05")
		if task.Due.IsZero() {
			dueStr = ""
		}
		tags, subWarnList := controller.FindTagsByTask(task.ID)
		warnList = append(warnList, subWarnList...)
		tagsStr := []string{}
		for _, tag := range tags {
			tagsStr = append(tagsStr, tag.Content)
		}
		row := table.Row{task.ID, contentStr, strings.Join(tagsStr, ","), dueStr, task.Loop}
		t.AppendRow(row)
	}
	t.AppendFooter(table.Row{"", fmt.Sprint("Found ", len(tasks), " tasks")})
	t.Style().Options.SeparateColumns = false
	t.Render()
	return
}

func Tags(tags []model.Tag) {
	t.AppendHeader(table.Row{"#", "Content", "Color"})
	for _, tag := range tags {
		row := table.Row{tag.ID, tag.Content, tag.Color}
		t.AppendRow(row)
	}
	t.AppendFooter(table.Row{"", fmt.Sprint("Found ", len(tags), " tasks")})
	t.Style().Options.SeparateColumns = false
	t.Render()
}
