package graphqlQuery

import (
	"testing"
)

func TestVariable_renderValue(t *testing.T) {
	type fields struct {
		Name    string
		Type    string
		Default *VariableDefault
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Variable without default",
			fields{
				Name:    "id",
				Type:    "ID",
				Default: nil,
			},
			"$id",
		},
		{
			"Variable with default",
			fields{
				Name:    "id",
				Type:    "ID",
				Default: &VariableDefault{Value: "\"12345\""},
			},
			"$id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Variable{
				Name:    tt.fields.Name,
				Type:    tt.fields.Type,
				Default: tt.fields.Default,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("Variable.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntValue_renderValue(t *testing.T) {
	type fields struct {
		Value int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Positive number",
			fields{Value: 123},
			"123",
		},
		{
			"Zero",
			fields{Value: 0},
			"0",
		},
		{
			"Negative number",
			fields{Value: -5},
			"-5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &IntValue{
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("IntValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatValue_renderValue(t *testing.T) {
	type fields struct {
		Value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Positive number",
			fields{Value: 123.456},
			"123.456",
		},
		{
			"Zero",
			fields{Value: 0.},
			"0",
		},
		{
			"Negative number",
			fields{Value: -0.2},
			"-0.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FloatValue{
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("FloatValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValue_renderValue(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Some string",
			fields{Value: "AnId"},
			"\"AnId\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &StringValue{
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("StringValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooleanValue_renderValue(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"True value",
			fields{Value: true},
			"true",
		},
		{
			"False value",
			fields{Value: false},
			"false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &BooleanValue{
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("BooleanValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullValue_renderValue(t *testing.T) {
	tests := []struct {
		name string
		v    *NullValue
		want string
	}{
		{
			"Null value",
			&NullValue{},
			"null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NullValue{}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("NullValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumValue_renderValue(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Some value",
			fields{Value: "JEDI"},
			"JEDI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &EnumValue{
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("EnumValue.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgument_renderValue(t *testing.T) {
	type fields struct {
		Name  string
		Value ArgumentValue
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Some string argument",
			fields{
				Name:  "filter",
				Value: &StringValue{Value: "123"},
			},
			"{\n  filter: \"123\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Argument{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("Argument.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListValueConst_renderValue(t *testing.T) {
	type fields struct {
		Values []ArgumentValue
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Empty list",
			fields{
				Values: make([]ArgumentValue, 0),
			},
			"[]",
		},
		{
			"List of int (one)",
			fields{
				Values: []ArgumentValue{
					&IntValue{Value: 123},
				},
			},
			"[123]",
		},
		{
			"List of int (a few)",
			fields{
				Values: []ArgumentValue{
					&IntValue{Value: 123},
					&IntValue{Value: 0},
					&IntValue{Value: -1},
				},
			},
			"[123, 0, -1]",
		},
		{
			"List of float (one)",
			fields{
				Values: []ArgumentValue{
					&FloatValue{Value: 123.456},
				},
			},
			"[123.456]",
		},
		{
			"List of float (a few)",
			fields{
				Values: []ArgumentValue{
					&FloatValue{Value: 1.1},
					&FloatValue{Value: 0.},
					&FloatValue{Value: -5.2},
				},
			},
			"[1.1, 0, -5.2]",
		},
		{
			"List of string (one)",
			fields{
				Values: []ArgumentValue{
					&StringValue{Value: "123"},
				},
			},
			"[\"123\"]",
		},
		{
			"List of string (a few)",
			fields{
				Values: []ArgumentValue{
					&StringValue{Value: "1.1"},
					&StringValue{Value: "abc"},
				},
			},
			"[\"1.1\", \"abc\"]",
		},
		{
			"List of bool (one)",
			fields{
				Values: []ArgumentValue{
					&BooleanValue{Value: true},
				},
			},
			"[true]",
		},
		{
			"List of bool (a few)",
			fields{
				Values: []ArgumentValue{
					&BooleanValue{Value: false},
					&BooleanValue{Value: true},
				},
			},
			"[false, true]",
		},
		{
			"List of null (one)",
			fields{
				Values: []ArgumentValue{
					&NullValue{},
				},
			},
			"[null]",
		},
		{
			"List of bool (a few)",
			fields{
				Values: []ArgumentValue{
					&NullValue{},
					&NullValue{},
				},
			},
			"[null, null]",
		},
		{
			"List of enum (one)",
			fields{
				Values: []ArgumentValue{
					&EnumValue{Value: "JEDI"},
				},
			},
			"[JEDI]",
		},
		{
			"List of enum (a few)",
			fields{
				Values: []ArgumentValue{
					&EnumValue{Value: "JEDI"},
					&EnumValue{Value: "EMPIRE"},
				},
			},
			"[JEDI, EMPIRE]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ListValueConst{
				Values: tt.fields.Values,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("ListValueConst.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListValueArgument_renderValue(t *testing.T) {
	type fields struct {
		Values []*Argument
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Empty list",
			fields{Values: []*Argument{}},
			"{}",
		},
		{
			"One element",
			fields{
				Values: []*Argument{
					{Name: "arg1", Value: &EnumValue{Value: "VALUE1"}},
				},
			},
			"{\n  arg1: VALUE1\n}",
		},
		{
			"A few elements",
			fields{
				Values: []*Argument{
					{Name: "arg1", Value: &EnumValue{Value: "VALUE1"}},
					{Name: "arg2", Value: &EnumValue{Value: "VALUE2"}},
				},
			},
			"{\n  arg1: VALUE1\n  arg2: VALUE2\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ListValueArgument{
				Values: tt.fields.Values,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("ListValueArgument.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectValueConst_renderValue(t *testing.T) {
	type fields struct {
		Values []*ListValueArgument
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Empty list",
			fields{Values: []*ListValueArgument{}},
			"[]",
		},
		{
			"One element",
			fields{
				Values: []*ListValueArgument{
					{
						Values: []*Argument{
							{Name: "title", Value: &StringValue{Value: "lesson title"}},
							{Name: "filePath", Value: &StringValue{Value: "static-resource-path"}},
						},
					},
				},
			},
			"[\n  {\n    title: \"lesson title\"\n    filePath: \"static-resource-path\"\n  }\n]",
		},
		{
			"One element",
			fields{
				Values: []*ListValueArgument{
					{
						Values: []*Argument{
							{Name: "title", Value: &StringValue{Value: "lesson title"}},
							{Name: "filePath", Value: &StringValue{Value: "static-resource-path"}},
						},
					},
					{
						Values: []*Argument{
							{Name: "title", Value: &StringValue{Value: "lesson title 2"}},
							{Name: "filePath", Value: &StringValue{Value: "static-resource-path 2"}},
						},
					},
				},
			},
			"[\n  {\n    title: \"lesson title\"\n    filePath: \"static-resource-path\"\n  }\n  {\n    title: \"lesson title 2\"\n    filePath: \"static-resource-path 2\"\n  }\n]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ObjectValueConst{
				Values: tt.fields.Values,
			}
			if got := v.renderValue(); got != tt.want {
				t.Errorf("ObjectValueConst.renderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
