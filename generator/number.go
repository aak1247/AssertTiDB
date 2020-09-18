package generator

import (
	"go/types"
	"math"
	"math/rand"
	"time"

	"moul.io/srand"
)

type numberGenerator struct {
	min  int64 // close
	max  int64 // open
	seed int64 // inner state
}

type IntegerGenerator struct {
	numberGenerator
	dataType types.BasicKind
}

func (this *IntegerGenerator) Init(min int64, max int64, dataType types.BasicKind) {
	this.min = min
	this.max = max
	this.dataType = dataType
}

func (this *IntegerGenerator) Generate() interface{} {
	if this.seed == 0 {
		rand.Seed(srand.MustSecure()) // UnixNano()表示纳秒
	} else {
		rand.Seed(time.Now().UnixNano() + this.seed)
	}
	n := this.max - this.min
	switch this.dataType {
	default:
		fallthrough
	case types.Int:
		num := rand.Intn(int(n)) + int(this.min)
		this.seed = int64(num)
		return num

	case types.Int32:
		num32 := rand.Int31n(int32(n)) + int32(this.min)
		this.seed = int64(num32)
		return num32

	case types.Int64:
		num64 := rand.Int63n(n) + this.min
		this.seed = num64
		return num64

	case types.Int16:
		num16 := int16(rand.Int31n(int32(n))) + int16(this.min)
		this.seed = int64(num16)
		return num16

	case types.Int8:
		num8 := int8(rand.Int31n(int32(n))) + int8(this.min)
		this.seed = int64(num8)
		return num8

	case types.Uint8:
		numu8 := uint8(rand.Intn(int(n))) + uint8(this.min)
		this.seed = int64(numu8)
		return numu8
	}
}

type FloatGenerator struct {
	numberGenerator
	min      float64
	max      float64
	dataType types.BasicKind
}

func (this *FloatGenerator) Generate() interface{} {
	if this.seed == 0 {
		rand.Seed(srand.MustSecure())
	} else {
		rand.Seed(time.Now().UnixNano() + this.seed)
	}
	n := this.max - this.min
	switch this.dataType {
	case types.Float32:
		num32 := rand.Float32()*float32(n) + float32(this.min)
		this.seed = int64(math.Round(float64(num32)))
		return num32
	case types.Float64:
		num64 := rand.Float64()*n + this.min
		this.seed = int64(math.Round(num64))
		return num64
	default:
		return 0
	}
}

func (this *FloatGenerator) Init(min, max float64, dataType types.BasicKind) {
	this.min = min
	this.max = max
	this.dataType = dataType
}
