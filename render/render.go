package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils/colorful"
	"github.com/jedib0t/go-pretty/v6/table"
)

var t table.Writer

func init() {
	Init()
}

func Init() {
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
}

func Tasks(tasks []model.Task, contenTitle string) (warnList []string) {
	defaultContentTitle := "Content"
	if contenTitle != "" {
		defaultContentTitle = contenTitle
	}
	t.AppendHeader(table.Row{"#", defaultContentTitle, "Tags", "Due", "Loop"})
	for _, task := range tasks {
		contentStr := task.Content
		if task.Important > 0 {
			contentStr = colorful.RenderStr(contentStr, "highlight", "", "")
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
	Init()
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

func Result(command string, errorList []error, warnList []string) {
	for _, err := range errorList {
		fmt.Printf("%s %s",
			colorful.RenderStr("[  INFO  ]: ", "default", "", "yellow"),
			err.Error(),
		)
	}
	if len(errorList) > 0 {
		for _, err := range errorList {
			fmt.Printf("%s %s",
				colorful.RenderStr("[ ERROR  ]: ", "default", "", "red"),
				err.Error(),
			)
		}
		fmt.Printf("%s %s",
			colorful.RenderStr("[ FAILED ]: ", "default", "", "green"),
			command,
		)
	} else {
		fmt.Printf("%s never %s",
			colorful.RenderStr("[   OK   ]: ", "default", "", "green"),
			command,
		)
	}
}
