package graphqlQuery

import "testing"

func TestArgument_RenderType(t *testing.T) {
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
			"Value is Variable",
			fields{
				Name:  "episode",
				Value: &Variable{Name: "episode", Type: "Episode", Default: nil},
			},
			"episode: $episode",
		},
		{
			"Value is int",
			fields{
				Name:  "arg",
				Value: &IntValue{Value: 123},
			},
			"arg: 123",
		},
		{
			"Value is float",
			fields{
				Name:  "arg",
				Value: &FloatValue{Value: 1.12},
			},
			"arg: 1.12",
		},
		{
			"Value is string",
			fields{
				Name:  "arg",
				Value: &StringValue{Value: "123"},
			},
			"arg: \"123\"",
		},
		{
			"Value is bool",
			fields{
				Name:  "arg",
				Value: &BooleanValue{Value: true},
			},
			"arg: true",
		},
		{
			"Value is null",
			fields{
				Name:  "arg",
				Value: &NullValue{},
			},
			"arg: null",
		},
		{
			"Value is enum",
			fields{
				Name:  "episode",
				Value: &EnumValue{Value: "JEDI"},
			},
			"episode: JEDI",
		},
		{
			"Value is argument",
			fields{
				Name:  "filter",
				Value: &Argument{Name: "field", Value: &IntValue{Value: 123}},
			},
			"filter: {\n  field: 123\n}",
		},
		{
			"Value is list of string",
			fields{
				Name: "someListArgument",
				Value: &ListValueConst{
					Values: []ArgumentValue{
						&StringValue{Value: "123"},
						&StringValue{Value: "456"},
					},
				},
			},
			"someListArgument: [\"123\", \"456\"]",
		},
		{
			"Value is list of arguments",
			fields{
				Name: "someListArgument",
				Value: &ListValueArgument{
					Values: []*Argument{
						{Name: "arg1", Value: &EnumValue{Value: "VALUE1"}},
						{Name: "arg2", Value: &EnumValue{Value: "VALUE2"}},
					},
				},
			},
			"someListArgument: {\n  arg1: VALUE1\n  arg2: VALUE2\n}",
		},
		{
			"Value is list of objects",
			fields{
				Name: "someListArgument",
				Value: &ObjectValueConst{
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
			},
			"someListArgument: [\n  {\n    title: \"lesson title\"\n    filePath: \"static-resource-path\"\n  }\n  {\n    title: \"lesson title 2\"\n    filePath: \"static-resource-path 2\"\n  }\n]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argument := &Argument{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := argument.RenderType(); got != tt.want {
				t.Errorf("Argument.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
