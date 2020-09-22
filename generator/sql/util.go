package sql

import "strconv"

func encodeFields(fields []string) string {
	var res string
	for i := 0; i < len(fields); i++ {
		if i > 0 {
			res += ", "
		}
		res += fields[i]
	}
	return res + " "
}

func encodeData(data []interface{}) string {
	var res string
	for i := 0; i < len(data); i++ {
		if i > 0 {
			res += ", "
		}
		res += encodeDatum(data[i])
	}
	return res
}

func encodeDatum(datum interface{}) string {
	switch datum.(type) {
	case string:
		return "\"" + datum.(string) + "\" "
	case int:
		return strconv.Itoa(datum.(int)) + " "
	case int16:
		return strconv.FormatInt(int64(datum.(int16)), 10) + " "
	case int32:
		return strconv.FormatInt(int64(datum.(int32)), 10) + " "
	case int64:
		return strconv.FormatInt(datum.(int64), 10) + " "
	case uint:
		return strconv.FormatInt(int64(datum.(uint)), 10) + " "
	case uint8:
		return strconv.FormatInt(int64(datum.(uint8)), 10) + " "
	case uint16:
		return strconv.FormatInt(int64(datum.(uint16)), 10) + " "
	case uint32:
		return strconv.FormatInt(int64(datum.(uint32)), 10) + " "
	case uint64:
		return strconv.FormatInt(int64(datum.(uint64)), 10) + " "
	case bool:
		if datum.(bool) {
			return "1 "
		} else {
			return "0 "
		}
	default:
		panic("not implemented")
	}
}

func GetFieldsFromColumns(columns []Column) []string {
	res := make([]string, len(columns))
	for i, column := range columns {
		res[i] = column.Name
	}
	return res
}

func GetFieldsFromColumnData(columnData []ColumnData) []string {
	res := make([]string, len(columnData))
	for i, col := range columnData {
		res[i] = col.Name
	}
	return res
}

func GetValuesFromColumnData(columnData []ColumnData) []interface{} {
	res := make([]interface{}, len(columnData))
	for i, col := range columnData {
		res[i] = col.Data
	}
	return res
}

type StringList []string

func (this StringList) join(char string) string {
	res := ""
	for i, str := range this {
		if i > 0 {
			res += char
		}
		res += str
	}
	return res
}
