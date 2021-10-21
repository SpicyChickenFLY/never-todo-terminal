package render

type Table struct {
	fieldNames     []string
	fieldMaxLength []int

	Rows []interface{}
}

const (
	rowTypeRecord = iota
	rowTypeLine
)

type Row struct {
	rowType int
}

type Record struct {
	fieldValue []string
}

func (t *Table) SetFieldName(idx int, name string) {}
func (t *Table) SetFieldMaxLength(idx, length int) {}

func (t *Table) AppendRecord() {}

func (t *Table) AppendDoubleLine() {}
func (t *Table) AppendSolidLine()  {}
func (t *Table) AppendEmptyLine()  {}
