package main

import (
	"context"
	"fmt"
	"go/types"
	"time"

	"github.com/aak1247/AssertTiDB/executor/db"
	"github.com/aak1247/AssertTiDB/generator"
	s "github.com/aak1247/AssertTiDB/generator/sql"
)

func main() {
	// testSQLGenerator()
	insertTest()
}

func insertTest() {
	mysql := db.GetDb()
	defer db.Close()
	mysql.Conn(context.Background())

	var table = s.Table{
		Name: "users",
		Columns: []s.Column{
			{Name: "id", DataType: types.Int64},
			{Name: "name", DataType: types.String},
			{Name: "eng", DataType: types.Int},
			{Name: "math", DataType: types.Int},
			{Name: "ch", DataType: types.Int},
		},
	}

	var stringGenerator generator.StringGenerator
	stringGenerator.Init(20, 21)
	name := stringGenerator.Generate()

	s1 := s.SqlGenerator{Table: table}
	w := s.Where(s.Eq("name", name))
	selectStatement := s1.GenerateSelect(s.GetFieldsFromColumns(table.Columns), s.Query(w))

	insertStatement := s1.GenerateInsert(
		[]s.ColumnData{
			{Name: "name", Data: name},
		},
	)
	res, e := mysql.Exec(insertStatement.Value)
	if e != nil {
		panic(e)
	}
	fmt.Println(res)
	rows, e := mysql.Query(selectStatement.Value)
	res2 := db.Scan(*rows, table.Columns)
	for _, m := range res2 {
		for k, v := range m {
			fmt.Println(k, (v.Data))
		}
	}
	fmt.Println(res2)
}

func testSQLGenerator() {
	// var w = s.Where(s.Eq("name", 1)).And(s.Gt("name", 2)).Or(s.Lt("name", 100)).And(s.Where(s.Eq("name", 666)).And(s.In("name", []interface{}{"6", int64(7), 8})))
	var w = s.Where(s.Gt("eng", 60)).And(s.Gt("ch", 9).Or(s.Gt("math", 9)))
	var str = w.String()
	fmt.Println(str)
	var table = s.Table{
		Name: "users",
		Columns: []s.Column{
			{Name: "id", DataType: types.Int64},
			{Name: "name", DataType: types.String},
			{Name: "eng", DataType: types.Int},
			{Name: "math", DataType: types.Int},
			{Name: "ch", DataType: types.Int},
		},
	}

	s1 := s.SqlGenerator{Table: table}
	selectStatement := s1.GenerateSelect(s.GetFieldsFromColumns(table.Columns), s.Query(w))
	fmt.Println(selectStatement.Value)
	var stringGenerator generator.StringGenerator
	stringGenerator.Init(20, 21)
	insertStatement := s1.GenerateInsert(
		[]s.ColumnData{
			{Name: "name", Data: stringGenerator.Generate()},
		},
	)
	fmt.Println(insertStatement.Value)
}

func testGenerator() {
	fmt.Print("??")
	var faced generator.GeneratorFaced
	var integerGenerator generator.IntegerGenerator
	var floatGenerator generator.FloatGenerator
	var float64Generator generator.FloatGenerator
	var stringGenerator generator.StringGenerator
	var timeGenerator generator.TimeGenerator
	var enumGenerator generator.EnumGenerator
	integerGenerator.Init(6, 100, types.Int64)
	faced.Inject(types.Int64, &integerGenerator)
	floatGenerator.Init(6, 100, types.Float32)
	faced.Inject(types.Float32, &floatGenerator)
	float64Generator.Init(6, 10, types.Float64)
	faced.Inject(types.Float64, &float64Generator)
	stringGenerator.Init(20, 21)
	faced.Inject(types.String, &stringGenerator)
	timeGenerator.Init(time.Now().AddDate(0, 0, -30), time.Now())
	faced.Inject(generator.Datetime, &timeGenerator)
	enums := []interface{}{"hello", "bye", "nice to meet you"}
	enumGenerator.Init(enums)
	faced.Inject(generator.Enum, &enumGenerator)
	// fmt.Print(t.Generate())
	// fmt.Println(faced.Length())
	for range make([]int, 100) {
		fmt.Println(faced.Generate(types.Int64))
		fmt.Println(faced.Generate(types.Float32))
		fmt.Println(faced.Generate(types.Float64))
		str := faced.Generate(types.String).(string)
		fmt.Println(str, len(str))
		fmt.Println(faced.Generate(generator.Datetime).(time.Time).String())
		fmt.Println(faced.Generate(generator.Enum))
	}

}
