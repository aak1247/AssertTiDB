package generator

import (
	"go/types"
	"time"
)

type TimeGenerator struct {
	integerGenerator IntegerGenerator
}

// [min, max)
func (this *TimeGenerator) Init(minTime, maxTime time.Time) {
	this.integerGenerator = IntegerGenerator{}
	this.integerGenerator.Init(minTime.Unix(), maxTime.Unix(), types.Int64)
}

// return type time.Time
func (this *TimeGenerator) Generate() interface{} {
	timestamp := this.integerGenerator.Generate().(int64)
	return time.Unix(timestamp, 0)
}
