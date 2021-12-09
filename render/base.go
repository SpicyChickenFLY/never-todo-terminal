package render

import (
	"fmt"
	"sort"
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
		record := record{task.ID}

		if task.Content != "" {
			contentStr := task.Content
			for i := 0; i < task.Important; i++ {
				// contentStr = colorful.RenderStr(contentStr, "line", "", "")
				contentStr += "*" // ★
			}
			record = append(record, contentStr)
		} else {
			record = append(record, nil)
		}

		tags, _ := controller.FindTagsByTask(task.ID)
		tagsStr := []string{}
		for _, tag := range tags {
			content := colorful.RenderStr(tag.Content, "default", "", tag.Color)
			tagsStr = append(tagsStr, content)
		}
		sort.SliceStable(tagsStr, func(i, j int) bool { return tagsStr[i] < tagsStr[j] })
		tagStr := strings.Join(tagsStr, ",")
		if tagStr == "" {
			record = append(record, nil)
		} else {
			record = append(record, tagStr)
		}

		dueStr := task.Due.Format("2006/01/02 15:04:05")
		if task.Due.IsZero() {
			record = append(record, nil)
		} else {
			record = append(record, dueStr)
		}
		if task.Loop == "" {
			record = append(record, nil)
		} else {
			record = append(record, task.Loop)
		}
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

// Seperator is a line full of "=" of "-"
func Seperator() {

}

// MainTheme render logos and statistics
func MainTheme() {

}

// Logo render logo in colors
func Logo() {

}
