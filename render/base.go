package render

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

type table struct {
	pageLen         int
	fieldNames      []string
	fieldLenLimit   []int
	fieldMaxLen     []int
	fieldEmptyFLags []bool

	rows []row
}

func newTable() *table {
	return &table{}
}

const (
	rowTypeRecord = iota
	rowTypeLine
)

type row struct {
	rowType     int
	fieldValues []string
}

type record []interface{}

func (t *table) calcFieldMaxLen() (sum int) {
	for _, maxLen := range t.fieldMaxLen {
		sum += maxLen
	}
	return
}

func (t *table) SetFieldNames(fieldNames []string) {
	t.fieldNames = make([]string, len(fieldNames))
	t.fieldMaxLen = make([]int, len(fieldNames))
	t.fieldEmptyFLags = make([]bool, len(fieldNames))
	for i, fieldName := range fieldNames {
		t.fieldNames[i] = fieldName
		t.fieldMaxLen[i] = utils.LenOnScreen(fieldName)
		t.fieldEmptyFLags[i] = true
	}
}
func (t *table) SetFieldLenLimit(idx, length int) {}

func (t *table) AppendRecord(record record) {
	if len(record) != len(t.fieldNames) {
		// TODO: fullfill or throw exception //
		return
	}
	fieldValues := make([]string, len(record))
	for i, field := range record {
		fieldContent := ""
		if field != nil {
			fieldContent = fmt.Sprint(field)
			fieldValues[i] = fieldContent
			t.fieldEmptyFLags[i] = false
			if utils.LenOnScreen(fieldContent) >= t.fieldMaxLen[i] {
				t.fieldMaxLen[i] = utils.LenOnScreen(fieldContent)
			}
		}
	}
	row := row{rowTypeRecord, fieldValues}
	t.rows = append(t.rows, row)
}

func (t *table) AppendDoubleLine() {}
func (t *table) AppendSolidLine()  {}
func (t *table) AppendEmptyLine()  {}

func (t *table) Render() {
	if t.pageLen < t.calcFieldMaxLen() {
		// TODO:  do something  //
		if model.DB.Settings.WrapContent {
		}
		fmt.Println("small page")
	}
	// render header
	for fieldIdx, fieldName := range t.fieldNames {
		if t.fieldEmptyFLags[fieldIdx] {
			continue
		}
		fmt.Print(fieldName)
		for i := 0; i <= t.fieldMaxLen[fieldIdx]-utils.LenOnScreen(fieldName); i++ {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for _, row := range t.rows {
		switch row.rowType {
		case rowTypeRecord:
			// TODO:  list all fields value of this row //
			for fieldIdx, field := range row.fieldValues {
				// 中文为代表的宽字符换行后会多占一个空格
				// 首先得判断是否需要换行，是否可以省略
				fmt.Print(field)
				for i := 0; i <= t.fieldMaxLen[fieldIdx]-utils.LenOnScreen(field); i++ {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}

func (t *table) Reset() {
}
