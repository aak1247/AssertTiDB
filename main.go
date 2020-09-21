package main

import (
	"fmt"
	"go/types"
	"time"

	"github.com/aak1247/AssertTiDB/generator"
	s "github.com/aak1247/AssertTiDB/generator/sql"
)

func main() {

	var t = s.Where(s.Eq("name", 1)).And(s.GreaterThan("name", 2)).Or(s.LessThan("name", 100)).And(s.Where(s.Eq("name", 666)).And(s.In("name", []interface{}{"6", int64(7), 8})))
	var str = t.String()
	fmt.Println(str)
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
