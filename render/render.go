package render

import (
	"fmt"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils/colorful"
)

var t *table

func init() {
	t = newTable()
	t.pageLen = 200
	t.Reset()
}

// Tasks in table
func Tasks(tasks []model.Task, contenTitle string) (warnList []string) {
	defaultContentTitle := "Content"
	if contenTitle != "" {
		defaultContentTitle = contenTitle
	}
	t.SetFieldNames([]string{"#", defaultContentTitle, "Tags", "Due", "Loop"})

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
		// project, ok := controller.GetProjectByID(task.ProjectID)
		// if !ok {
		// 	project = model.Project{}
		// }
		record := record{task.ID, contentStr, strings.Join(tagsStr, ","), dueStr, task.Loop}
		t.AppendRecord(record)
	}
	t.Render()
	t.Reset()
	return
}

// Tags in table
func Tags(tags []model.Tag) {
	t.SetFieldNames([]string{"#", "Content", "Color"})
	for _, tag := range tags {
		color := colorful.RenderStr(tag.Color, "default", "", tag.Color)
		record := record{tag.ID, tag.Content, color}
		t.AppendRecord(record)
	}
	t.Render()
	t.Reset()
}

// Result of execution
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
		fmt.Printf("%s never %s\n",
			colorful.RenderStr("[ FAILED ]: ", "default", "", "red"),
			command,
		)
	} else {
		fmt.Printf("%s never %s\n",
			colorful.RenderStr("[   OK   ]: ", "default", "", "green"),
			command,
		)
	}
	t.Reset()
}
