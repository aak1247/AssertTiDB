package sql

// 能够随机生成insert、select语句，并根据语句生成预期结果

type SqlOpe int

const (
	Insert SqlOpe = iota
	Delete
	Update
	Select
)

type Column struct {
	head     string
	dataType interface{}
	// types.BasicKind
	// or time.Time
	// or types.Map
	// or types.Array
	primary bool
	index   bool
}

type Table struct {
	Name    string
	Columns []Column
}

type SqlGenerator struct {
	Table Table
}

func (this *SqlGenerator) Init(table Table) {
	this.Table = table
}

// func (this *SqlGenerator) generateSelect(fields []string, query Query) {

// }

type Sql interface {
	Execute() interface{}
}

type SqlParams interface {
}

type SelectParams struct {
	fields string
	query  query
}

func (this *Table) Generate(sqlOpe SqlOpe, params SqlParams) Sql {
	switch sqlOpe {
	case Insert:
	case Delete:
	case Update:
	case Select:

	default:
	}
	return nil
}
