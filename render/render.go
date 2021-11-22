package render

import (
	"fmt"

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
	tf := taskFormatter{}
	defaultContentTitle := "Content"
	if contenTitle != "" {
		defaultContentTitle = contenTitle
	}
	t.SetFieldNames([]string{"#", defaultContentTitle, "Tags", "Due", "Loop"})

	for _, task := range tasks {
		record := taskRecord{id: fmt.Sprint(task.ID)}

		if task.Content != "" {
			contentStr := task.Content
			for i := 0; i < task.Important; i++ {
				contentStr += "!"
			}
			record.content = contentStr
		}

		tags, _ := controller.FindTagsByTask(task.ID)
		for _, tag := range tags {
			record.tagsContent = append(record.tagsContent, tag.Content)
			record.tagsColor = append(record.tagsColor, tag.Color)
		}

		dueStr := task.Due.Format("2006/01/02 15:04:05")
		if !task.Due.IsZero() {
			record.dueStr = dueStr
		}
		record.loopStr = task.Loop
		tf.records = append(tf.records, record)
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
