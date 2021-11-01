package render

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

type table struct {
	pageLen       int
	fieldNames    []string
	fieldLenLimit []int
	fieldMaxLen   []int

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

func (t *table) calcrowMaxLen() (sum int) {
	for _, maxLen := range t.fieldMaxLen {
		sum += maxLen
	}
	return
}

func (t *table) SetFieldNames(fieldNames []string) {
	t.fieldNames = make([]string, len(fieldNames))
	t.fieldMaxLen = make([]int, len(fieldNames))
	for i, fieldName := range fieldNames {
		t.fieldNames[i] = fieldName
		t.fieldMaxLen[i] = len(fieldName)
	}
}
func (t *table) SetFieldLenLimit(idx, length int) {}

func (t *table) AppendRecord(record record) {
	// TODO: 中文汉字的length是有问题的
	if len(record) != len(t.fieldNames) {
		// TODO: fullfill or throw exception //
		return
	}
	fieldValues := make([]string, len(record))
	for i, field := range record {
		fieldContent := fmt.Sprint(field)
		fieldValues[i] = fieldContent
		if len(fieldContent) >= t.fieldMaxLen[i] {
			t.fieldMaxLen[i] = utils.LenOnScreen(fieldContent)
		}
	}
	row := row{rowTypeRecord, fieldValues}
	t.rows = append(t.rows, row)
}

func (t *table) AppendDoubleLine() {}
func (t *table) AppendSolidLine()  {}
func (t *table) AppendEmptyLine()  {}

func (t *table) Render() {
	if t.pageLen < t.calcrowMaxLen() {
		// TODO:  do something  //
		fmt.Println("small page")
	}
	for _, row := range t.rows {
		switch row.rowType {
		case rowTypeRecord:
			// TODO:  list all fields value of this row //
			for fieldIdx, field := range row.fieldValues {
				fmt.Print(field)
				for i := 0; i <= t.fieldMaxLen[fieldIdx]-len(field); i++ {
					fmt.Print(" ")
				}
				fmt.Print(t.fieldMaxLen[fieldIdx], len(field))
			}
			fmt.Println()
		}
	}
}

func (t *table) Reset() {
}
