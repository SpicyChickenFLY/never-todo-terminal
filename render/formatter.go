package render

import "github.com/SpicyChickenFLY/never-todo-cmd/model"

type formatter interface {
	AppendRecord()
	AppendSeparator(sep string)
	Render()
	Reset()
}

type taskRecords struct {
	id          string
	content     string
	tagsContent []string
	tagsColor   []string
	dueStr      string
	loopStr     string
}

type taskFormatter struct {
	contentFieldLen    int
	contentFieldMaxLen int
	tagFieldLen        int
	tagFieldMaxLen     int
}

func newTaskFormatter() *taskFormatter {
	return &taskFormatter{}
}

func (tf *taskFormatter) Render() {
	table := newTable()

	pageLen, err := lenOfTerminal()
	if err != nil {
		return err
	}
	if tf.contentFieldMaxLen+tf.tagFieldMaxLen > pageLen {

	}

	if model.DB.Settings.CompressTask {

	} else if model.DB.Settings.WrapContent {

	}

}

type tagFormatter struct {
}
