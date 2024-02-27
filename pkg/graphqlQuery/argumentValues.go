package graphqlQuery

import (
	"strconv"
	"strings"
)

type ArgumentValue interface {
	// method for render of a value
	renderValue() string
}

var _ ArgumentValue = (*Variable)(nil)
var _ ArgumentValue = (*IntValue)(nil)
var _ ArgumentValue = (*FloatValue)(nil)
var _ ArgumentValue = (*StringValue)(nil)
var _ ArgumentValue = (*BooleanValue)(nil)
var _ ArgumentValue = (*NullValue)(nil)
var _ ArgumentValue = (*EnumValue)(nil)
var _ ArgumentValue = (*Argument)(nil)
var _ ArgumentValue = (*ListValueConst)(nil)
var _ ArgumentValue = (*ListValueArgument)(nil)
var _ ArgumentValue = (*ObjectValueConst)(nil)

type IntValue struct {
	Value int
}

type FloatValue struct {
	Value float64
}

type StringValue struct {
	Value string
}

type BooleanValue struct {
	Value bool
}

type NullValue struct{}

type EnumValue struct {
	Value string
}

type ListValueConst struct {
	Values []ArgumentValue
}

type ListValueArgument struct {
	Values []*Argument
}

type ObjectValueConst struct {
	Values []*ListValueArgument
}

func (v *Variable) renderValue() string {
	return "$" + v.Name
}

func (v *IntValue) renderValue() string {
	if v == nil {
		return ""
	}
	return strconv.FormatInt(int64(v.Value), 10)
}

func (v *FloatValue) renderValue() string {
	if v == nil {
		return ""
	}
	return strconv.FormatFloat(v.Value, 'g', -1, 64)
}

func (v *StringValue) renderValue() string {
	if v == nil {
		return ""
	}
	return "\"" + v.Value + "\""
}

func (v *BooleanValue) renderValue() string {
	if v == nil {
		return ""
	}
	return strconv.FormatBool(v.Value)
}

func (v *NullValue) renderValue() string {
	if v == nil {
		return ""
	}
	return "null"
}

func (v *EnumValue) renderValue() string {
	if v == nil {
		return ""
	}
	return v.Value
}

func (v *Argument) renderValue() string {
	if v == nil {
		return ""
	}
	return "{\n  " + lineShift(v.RenderType()) + "\n}"
}

func (v *ListValueConst) renderValue() string {
	if v == nil {
		return ""
	}

	renderedValues := make([]string, 0)

	for _, item := range v.Values {
		renderedValues = append(renderedValues, item.renderValue())
	}

	return "[" + strings.Join(renderedValues, ", ") + "]"
}

func (v *ListValueArgument) renderValue() string {
	if v == nil {
		return ""
	}

	if len(v.Values) == 0 {
		return "{}"
	}

	renderedValues := make([]string, 0)

	for _, item := range v.Values {
		renderedValues = append(renderedValues, item.RenderType())
	}

	return "{\n  " + lineShift(strings.Join(renderedValues, "\n")) + "\n}"
}

func (v *ObjectValueConst) renderValue() string {
	if v == nil {
		return ""
	}
	if len(v.Values) == 0 {
		return "[]"
	}

	renderedValues := make([]string, 0)

	for _, item := range v.Values {
		renderedValues = append(renderedValues, item.renderValue())
	}

	return "[\n  " + lineShift(strings.Join(renderedValues, "\n")) + "\n]"
}
