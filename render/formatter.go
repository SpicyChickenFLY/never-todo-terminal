package render

import "github.com/SpicyChickenFLY/never-todo-cmd/model"

type formatter interface {
	AppendRecord()
	AppendSeparator(sep string)
	Render()
	Reset()
}

type taskRecord struct {
	id          string
	content     string
	tagsContent []string
	tagsColor   []string
	dueStr      string
	loopStr     string
}

type taskFormatter struct {
	records            []taskRecord
	contentFieldLen    int
	contentFieldMaxLen int
	tagFieldLen        int
	tagFieldMaxLen     int
}

func newTaskFormatter() *taskFormatter {
	return &taskFormatter{}
}

func (tf *taskFormatter) Render() {
	// table := newTable()

	pageLen, err := lenOfTerminal()
	if err != nil {
		// return err
	}
	if tf.contentFieldMaxLen+tf.tagFieldMaxLen > pageLen {

	}

	if model.DB.Settings.WrapContent {

	} else if model.DB.Settings.CompressTask {

	}

}

type tagFormatter struct {
}
