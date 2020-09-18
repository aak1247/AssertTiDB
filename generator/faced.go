package generator

type GeneratorFaced struct {
	generators map[interface{}]DataGenerator
}

var DefaultGeneratorFaced = GeneratorFaced{
	generators: map[interface{}]DataGenerator{},
}

func (generatorFaced *GeneratorFaced) Inject(_type interface{}, generator DataGenerator) {
	if generatorFaced.generators == nil {
		generatorFaced.generators = make(map[interface{}]DataGenerator)
	}
	generatorFaced.generators[_type] = generator
}

func (generatorFaced *GeneratorFaced) Generate(_type interface{}) interface{} {
	return (generatorFaced.generators[_type]).Generate()
}

func (this *GeneratorFaced) Length() int {
	return len(this.generators)
}
