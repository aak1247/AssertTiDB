package generator

type DataGenerator interface {
	Generate() interface{}
}

type DefaultGenerator struct {
	// 	GenerateWithSeed(seed interface{}) interface{}
	// Init(config interface{}) interface{}
}
