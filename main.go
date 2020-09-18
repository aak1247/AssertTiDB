package main

import (
	"fmt"
	"go/types"
	"time"

	"github.com/aak1247/AssertTiDB/generator"
)

func main() {
	fmt.Print("??")
	var faced generator.GeneratorFaced
	var integerGenerator generator.IntegerGenerator
	var floatGenerator generator.FloatGenerator
	var float64Generator generator.FloatGenerator
	var stringGenerator generator.StringGenerator
	var timeGenerator generator.TimeGenerator
	integerGenerator.Init(6, 100, types.Int64)
	faced.Inject(types.Int64, &integerGenerator)
	floatGenerator.Init(6, 100, types.Float32)
	faced.Inject(types.Float32, &floatGenerator)
	float64Generator.Init(6, 10, types.Float64)
	faced.Inject(types.Float64, &float64Generator)
	stringGenerator.Init(20, 21)
	faced.Inject(types.String, &stringGenerator)
	timeGenerator.Init(time.Now().AddDate(0, 0, -30), time.Now())
	// fmt.Print(t.Generate())
	// fmt.Println(faced.Length())
	for range make([]int, 100) {
		fmt.Println(faced.Generate(types.Int64))
		fmt.Println(faced.Generate(types.Float32))
		fmt.Println(faced.Generate(types.Float64))
		str := faced.Generate(types.String).(string)
		fmt.Println(str, len(str))
		fmt.Printf(timeGenerator.Generate().(time.Time).String())
	}
}
