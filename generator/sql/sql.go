package sql

import (
	"database/sql"
	"go/types"
	"reflect"
	"time"

	"github.com/aak1247/AssertTiDB/generator"
)

// 能够随机生成insert、select语句，并根据语句生成预期结果

type SqlOpe int

const (
	Insert SqlOpe = iota
	Delete
	Update
	Select
)

type Column struct {
	Name     string
	DataType interface{}
	// types.BasicKind
	// or time.Time
	// or types.Map
	// or types.Array
	Primary bool
	Index   bool
}

// func (this Column) GenerateColumnData() ColumnData {
// 	switch this.DataType {
// 	case types.Int:
// 		return ColumnData{Data: int(0)}
// 	case types.Int16:
// 		return ColumnData{Data: int16(0)}
// 	case types.Int32:
// 		return ColumnData{Data: int32(0)}
// 	case types.Int64:
// 		return ColumnData{Data: int64(0)}
// 	case types.Int8:
// 		return ColumnData{Data: int8(0)}
// 	case types.Uint:
// 		return ColumnData{Data: uint(0)}
// 	case types.Uint16:
// 		return ColumnData{Data: uint16(0)}
// 	case types.Uint32:
// 		return ColumnData{Data: uint32(0)}
// 	case types.Uint64:
// 		return ColumnData{Data: uint64(0)}
// 	case types.Uint8:
// 		return ColumnData{Data: uint8(0)}
// 	case types.Bool:
// 		return ColumnData{Data: false}
// 	case types.String:
// 		return ColumnData{Data: ""}
// 	case generator.Datetime:
// 		return ColumnData{Data: time.Time{}}
// 	default:
// 		// other???
// 		return ColumnData{Data: nil}
// 	}
// }

type nullInt int

func (this *nullInt) Scan(v interface{}) error {
	if v == nil {
		*this = 0
	} else {
		t := sql.NullInt64{}
		t.Scan(v)
		*this = nullInt(t.Int64)
	}
	return nil
}

func (this Column) GenerateColumnRefData() ColumnData {
	switch this.DataType {
	case types.Int:
		t := nullInt(0)
		return ColumnData{Data: &t}
	case types.Int16:
		t := int16(0)
		return ColumnData{Data: &t}
	case types.Int32:
		t := int32(0)
		return ColumnData{Data: &t}
	case types.Int64:
		t := int64(0)
		return ColumnData{Data: &t}
	case types.Int8:
		t := int8(0)
		return ColumnData{Data: &t}
	case types.Uint:
		t := uint(0)
		return ColumnData{Data: &t}
	case types.Uint16:
		t := uint16(0)
		return ColumnData{Data: &t}
	case types.Uint32:
		t := uint32(0)
		return ColumnData{Data: &t}
	case types.Uint64:
		t := uint64(0)
		return ColumnData{Data: &t}
	case types.Uint8:
		t := uint8(0)
		return ColumnData{Data: &t}
	case types.Bool:
		t := false
		return ColumnData{Data: &t}
	case types.String:
		t := ""
		return ColumnData{Data: &t}
	case generator.Datetime:
		t := time.Time{}
		return ColumnData{Data: &t}
	default:
		// other???
		var t interface{} = nil
		return ColumnData{Data: &t}
	}
}

type ColumnData struct {
	Name string
	Data interface{}
}

func (this *ColumnData) DeRef() {
	switch this.Data.(type) {
	case *int:
		t := this.Data.(*int)
		this.Data = *t
	case *int64:
		t := this.Data.(*int64)
		this.Data = *t
	case *nullInt:
		t := this.Data.(*nullInt)
		this.Data = *t
	case *int16:
		t := this.Data.(*int16)
		this.Data = *t
	case *int32:
		t := this.Data.(*int32)
		this.Data = *t
	case *uint:
		t := this.Data.(*uint)
		this.Data = *t
	case *uint8:
		t := this.Data.(*uint8)
		this.Data = *t
	case *uint16:
		t := this.Data.(*uint16)
		this.Data = *t
	case *uint32:
		t := this.Data.(*uint32)
		this.Data = *t
	case *uint64:
		t := this.Data.(*uint64)
		this.Data = *t
	case *string:
		t := this.Data.(*string)
		this.Data = *t
	default:
		this.Data = reflect.Indirect(reflect.ValueOf(this.Data))
	}
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

func (this *SqlGenerator) GenerateSelect(fields []string, query query) Sql {
	return Sql{Value: "SELECT " + encodeFields(fields) + "FROM " + this.Table.Name + " " + query.String() + " ;"}
}

func (this *SqlGenerator) GenerateInsert(data []ColumnData) Sql {
	return Sql{Value: "INSERT INTO " + this.Table.Name + " ( " + StringList(GetFieldsFromColumnData(data)).join(", ") + " " + ") values ( " + encodeData(GetValuesFromColumnData(data)) + " " + ") ;"}
}

func (this *SqlGenerator) GenerateUpdate() Sql {
	// todo
	return Sql{}
}

func (this *SqlGenerator) GenerateDelete() Sql {
	// todo
	return Sql{}
}

type Sql struct {
	Value string
}

type SqlParams interface {
}

type SelectParams struct {
	fields []string
	query  query
}

type InsertParams struct {
	data []ColumnData
}

func (this *SqlGenerator) Generate(sqlOpe SqlOpe, params SqlParams) Sql {
	switch sqlOpe {
	case Insert:
		return this.GenerateInsert(params.(InsertParams).data)
	case Delete:
		return this.GenerateDelete()
	case Update:
		return this.GenerateUpdate()
	case Select:
		return this.GenerateSelect(params.(SelectParams).fields, params.(SelectParams).query)
	default:
		// panic
	}
	return Sql{}
}
