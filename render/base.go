package render

import "fmt"

type Table struct {
	pageLen       int
	fieldNames    []string
	fieldLenLimit []int
	fieldMaxLen   []int

	rows []Row
}

const (
	rowTypeRecord = iota
	rowTypeLine
)

type Row struct {
	rowType     int
	fieldValues []string
}

type Record []interface{}

func (t *Table) calcRowMaxLen() (sum int) {
	for _, maxLen := range t.fieldMaxLen {
		sum += maxLen
	}
	return
}

func (t *Table) SetFieldNames(fieldNames []string) {
	if len(fieldNames) != len(t.fieldNames) {
		// TODO: fullfill or throw exception //
		return
	}
	for i, fieldName := range fieldNames {
		t.fieldNames[i] = fieldName
		if len(fieldName) > t.fieldMaxLen[i] {
			t.fieldMaxLen[i] = len(fieldName)
		}
	}
}
func (t *Table) SetFieldLenLimit(idx, length int) {}

func (t *Table) AppendRecord(record Record) {
	if len(record) != len(t.fieldNames) {
		// TODO: fullfill or throw exception //
		return
	}
	fieldValues := make([]interface{}, 0)
	for i, field := range record {
		fieldContent := fmt.Sprint(field)
		fieldValues = append(fieldValues, fieldContent)
		if len(fieldContent) >= t.fieldMaxLen[i] {
			t.fieldMaxLen[i] = len(fieldContent)
		}
	}
	row := Row{rowTypeRecord, fieldValues}
}

func (t *Table) AppendDoubleLine() {}
func (t *Table) AppendSolidLine()  {}
func (t *Table) AppendEmptyLine()  {}

func (t *Table) Render() {
	if t.pageLen < t.calcRowMaxLen() {
		// TODO:  do something  //
	}
	for _, row := range t.rows {
		switch row.rowType {
		case rowTypeRecord:
			// TODO:  <22-10-21, yourname> //
		}
	}
}
