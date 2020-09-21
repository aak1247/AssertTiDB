package sql

import "strconv"

type logistic int

const (
	and logistic = iota
	or
	not
)

type compareOpe int

const (
	eq compareOpe = iota
	lt
	gt
	in
)

type ConditionContext int

type query struct {
	condition Conditions
}

func Query(con Conditions) query {
	return query{condition: con}
}

type conditionable interface {
	String() string
}

func encodeConditionable(con conditionable) string {
	switch con.(type) {
	case condition:
		return con.(condition).String()
	case Conditions:
		c := con.(Conditions)
		if len(c.conditions) == 1 {
			return c.__String()
		} else {
			return "( " + c.__String() + ") "
		}
	default:
		return ""
	}
}

type Conditions struct {
	context    conditionable
	conditions []conditionable // condition or condition
	ope        logistic
}

func Where(con conditionable) Conditions {
	l := make([]conditionable, 1)
	conditions := Conditions{}
	l[0] = con
	conditions.conditions = l
	conditions.ope = and
	// conditions.context = nil
	return conditions
}

func (this Conditions) And(con conditionable) Conditions {

	if this.ope == or {
		// ope == or
		// `and` is more prior
		newConditions := Conditions{}
		newConditions.conditions = []conditionable{this.conditions[len(this.conditions)-1], con}
		newConditions.ope = and
		this.conditions[len(this.conditions)-1] = newConditions
		newConditions.context = this
		return newConditions
	} else if this.ope == and {
		switch con.(type) {
		case Conditions:
			t := con.(Conditions)
			t.context = this
		}
		this.conditions = append(this.conditions, con)
		return this
	}
	return this
}

func (this Conditions) Or(con conditionable) Conditions {
	if this.ope == or {
		this.conditions = append(this.conditions, con)
		return this
	} else if this.ope == and {
		// ope == or
		// `and` is more prior
		// make and sub of or
		newConditions := Conditions{}
		newConditions.ope = or
		newConditions.conditions = []conditionable{this, con}
		this.context = newConditions
		return newConditions
	}
	return this
}

func (this Conditions) __String() string {
	var res string
	var opeStr string
	switch this.ope {
	default:
		fallthrough
	case and:
		opeStr = "AND"
	case or:
		opeStr = "OR"
	case not:
		if len(this.conditions) != 1 {
			panic("not must be applied to one conditionable")
		}
		return "NOT " + encodeConditionable(this.conditions[0])
	}

	for i, con := range this.conditions {
		if i > 0 {
			res += opeStr + " "
		}
		res += encodeConditionable(con)
	}

	return res
}

func (this Conditions) String() string {
	if this.context == nil {
		return this.__String()
	} else {
		return this.context.String()
	}

}

type condition struct {
	field string
	rel   compareOpe
	data  []interface{}
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

func (this condition) String() string {
	switch this.rel {
	case eq:
		return this.field + " = " + encodeData(this.data)
	case lt:
		return this.field + " < " + encodeData(this.data)
	case gt:
		return this.field + " > " + encodeData(this.data)
	case in:
		return this.field + " IN ( " + encodeData(this.data) + ") "
	}
	return ""
}

func Eq(field string, data interface{}) condition {
	return condition{field: field, rel: eq, data: []interface{}{data}}
}

func LessThan(field string, data interface{}) condition {
	return condition{field: field, rel: lt, data: []interface{}{data}}
}

func GreaterThan(field string, data interface{}) condition {
	return condition{field: field, rel: gt, data: []interface{}{data}}
}

func In(field string, data []interface{}) condition {
	return condition{field: field, rel: in, data: data}
}
