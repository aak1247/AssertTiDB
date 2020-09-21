package generator

import "go/types"

type EnumGenerator struct {
	enums        []interface{}
	intGenerator IntegerGenerator
}

func (this *EnumGenerator) Init(enums []interface{}) {
	this.enums = enums
	integerGenerator := IntegerGenerator{}
	integerGenerator.Init(0, int64(len(enums)), types.Int64)
	this.intGenerator = integerGenerator
}

func (this *EnumGenerator) Generate() interface{} {
	return this.enums[this.intGenerator.Generate().(int64)]
}
