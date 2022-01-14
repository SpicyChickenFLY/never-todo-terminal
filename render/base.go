package render

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/colorful"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
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
		record := make(record, recordLen)
		record[recordID] = task.ID

		if task.Content != "" {
			contentStr := task.Content
			for i := 0; i < task.Important; i++ {
				// contentStr = colorful.RenderStr(contentStr, "line", "", "")
				contentStr += "*" // ★
			}
			record[recordContent] = contentStr
		}

		tags, _ := controller.FindTagsByTask(task.ID)
		tagsStr := []string{}
		for _, tag := range tags {
			content := colorful.RenderStr(tag.Content, "default", "", tag.Color)
			tagsStr = append(tagsStr, content)
		}
		sort.SliceStable(tagsStr, func(i, j int) bool { return tagsStr[i] < tagsStr[j] })
		tagStr := strings.Join(tagsStr, ",")
		if tagStr != "" {
			record[recordTags] = tagStr
		}

		if !task.Due.IsZero() {
			dueStr := task.Due.Format("2006/01/02 15:04:05")
			estTimeStr := utils.EstimateTime(task.Due, time.Now(), false)
			record[recordDue] = dueStr + estTimeStr
		}

		if task.Loop != "" {
			record[recordLoop] = task.Loop
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

// Summary of this software and statistics
func Summary() {
	// 展示logo，用法，当前待办数和标签数
	Logo()
	var todoTotal, doneTotal, tagTotal int
	for _, task := range model.DB.Data.Tasks {
		switch task.Status {
		case model.TaskTodo:
			todoTotal++
		case model.TaskDone:
			doneTotal++
		}
	}
	for _, tag := range model.DB.Data.Tags {
		if !tag.Deleted {
			tagTotal++
		}
	}
	fmt.Printf("todo: %d, done: %d, tag: %d\n", todoTotal, doneTotal, tagTotal)

	// tasks := controller.ListTodoTasks()
	// Tasks(tasks, "")
}

// Logo render logo in colors
func Logo() {
	fmt.Println(ColorfulLogo)
	fmt.Println(Descirption)
	fmt.Println(Separator)
}
