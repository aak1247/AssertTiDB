package generator

import (
	"go/types"
)

type StringGenerator struct {
	minLen int
	maxLen int
}

// minLen 最小长度
// maxLen 最大长度
// 左闭右开
func (this *StringGenerator) Init(minLen, maxLen int) {
	this.minLen = minLen
	this.maxLen = maxLen
}

func (this *StringGenerator) Generate() interface{} {
	var lenGenerator = IntegerGenerator{}
	lenGenerator.Init(int64(this.minLen), int64(this.maxLen), types.Int64)
	var byteGenerator = IntegerGenerator{}
	byteGenerator.Init(-128, 128, types.Byte)
	len := lenGenerator.Generate()
	var bytes = make([]rune, len.(int64))
	for i := 0; int64(i) < len.(int64); i++ {
		bytes[i] = rune(byteGenerator.Generate().(byte))
	}
	return string(bytes)
}
