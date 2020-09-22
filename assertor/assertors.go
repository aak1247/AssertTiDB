package assertor

type Assertor struct {
	statement string
	expected  interface{}
}

type convertor interface {
	Convert(result interface{}) interface{}
}

var convertors map[interface{}]convertor = make(map[interface{}]convertor, 0)

func RegisterConvertor(resultType interface{}, convertor convertor) {
	convertors[resultType] = convertor
}

func (this *Assertor) Convert(result interface{}, resultType interface{}) interface{} {
	if con := convertors[resultType]; con != nil {
		return con.Convert(result)
	}
	return nil
}
