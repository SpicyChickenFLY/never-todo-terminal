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
	t.SetStyle(table.StyleLight)
}

func Tasks(tasks []model.Task, contenTitle string) (warnList []string) {
	defaultContentTitle := "Content"
	if contenTitle != "" {
		defaultContentTitle = contenTitle
	}
	t.AppendHeader(table.Row{"#", defaultContentTitle, "Project", "Tags", "Due", "Loop"})

	for _, task := range tasks {
		contentStr := task.Content
		for i := 0; i < task.Important; i++ {
			// contentStr = colorful.RenderStr(contentStr, "line", "", "")
			contentStr += "*" // ★
		}
		dueStr := task.Due.Format("2006/01/02 15:04:05")
		if task.Due.IsZero() {
			dueStr = ""
		}
		tags := controller.FindTagsByTask(task.ID)
		tagsStr := []string{}
		for _, tag := range tags {
			content := colorful.RenderStr(tag.Content, "default", "", tag.Color)
			tagsStr = append(tagsStr, content)
		}
		project, ok := controller.GetProjectByID(task.ProjectID)
		if !ok {
			project = model.Project{}
		}
		row := table.Row{task.ID, contentStr, project.Content, strings.Join(tagsStr, ","), dueStr, task.Loop}
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
		color := colorful.RenderStr(tag.Color, "default", "", tag.Color)
		row := table.Row{tag.ID, tag.Content, color}
		t.AppendRow(row)
	}
	t.AppendFooter(table.Row{"", fmt.Sprint("Found ", len(tags), " tasks")})
	t.Style().Options.SeparateColumns = false
	t.Render()
	Init()
}

func Result(command string, errorList []error, warnList []string) {
	for _, warn := range warnList {
		fmt.Printf("%s %s\n",
			colorful.RenderStr("[  INFO  ]: ", "default", "", "yellow"),
			warn,
		)
	}
	if len(errorList) > 0 {
		for _, err := range errorList {
			fmt.Printf("%s %s\n",
				colorful.RenderStr("[ ERROR  ]: ", "default", "", "red"),
				err.Error(),
			)
		}
		fmt.Printf("%s %s\n",
			colorful.RenderStr("[ FAILED ]: ", "default", "", "red"),
			command,
		)
	} else {
		fmt.Printf("%s never %s\n",
			colorful.RenderStr("[   OK   ]: ", "default", "", "green"),
			command,
		)
	}
	Init()
}
