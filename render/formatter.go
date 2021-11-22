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
	records              []taskRecord
	idMaxLen             int
	contentFieldLen      int
	contentFieldMaxLen   int
	tagFieldLen          int
	tagFieldTotalMaxLen  int
	tagFieldSingleMaxLen int
}

func newTaskFormatter() *taskFormatter {
	return &taskFormatter{}
}

func (tf *taskFormatter) Render() {
	// table := newTable()

	// calculate the max size
	pageLen, err := lenOfTerminal()
	if err != nil {
		// return err
	}
	// TODO: 先判断在标签折行的情况下，任务内容是否可以完整渲染
	if tf.contentFieldMaxLen <= pageLen-tf.idMaxLen-tf.tagFieldSingleMaxLen {
		// ok, could be rendered
	}
	// else we should controll the size
	// calculate fields maxlen by ratio
	if model.DB.Settings.WrapContent {

	} else if model.DB.Settings.CompressTask {

	}

}

type tagFormatter struct {
}
