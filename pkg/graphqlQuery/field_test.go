package graphqlQuery

import (
	"testing"
)

func TestStringField_renderSelection(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Some string field",
			fields{Value: "field"},
			"field",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &StringField{
				Value: tt.fields.Value,
			}
			if got := v.renderSelection(); got != tt.want {
				t.Errorf("StringField.renderSelection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFragment_renderSelection(t *testing.T) {
	type fields struct {
		Name     string
		Type     string
		Fields   []Selection
		Typename bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Some fragment",
			fields{
				Name:     "MyFragment",
				Type:     "MyType",
				Fields:   []Selection{&StringField{Value: "f1"}},
				Typename: true,
			},
			"...MyFragment",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Fragment{
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				Fields:   tt.fields.Fields,
				Typename: tt.fields.Typename,
			}
			if got := v.renderSelection(); got != tt.want {
				t.Errorf("Fragment.renderSelection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_RenderType(t *testing.T) {
	type fields struct {
		Name       string
		Alias      *Alias
		Arguments  []*Argument
		Fields     []Selection
		Directives []*Directive
		Typename   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simple field",
			fields{
				Name:       "fields",
				Alias:      nil,
				Arguments:  []*Argument{},
				Directives: []*Directive{},
				Fields:     []Selection{&StringField{Value: "f1"}, &StringField{Value: "f2"}},
				Typename:   false,
			},
			"fields {\n  f1\n  f2\n}",
		},
		{
			"simple field with typename",
			fields{
				Name:       "fields",
				Alias:      nil,
				Arguments:  []*Argument{},
				Directives: []*Directive{},
				Fields:     []Selection{&StringField{Value: "f1"}, &StringField{Value: "f2"}},
				Typename:   true,
			},
			"fields {\n  __typename\n  f1\n  f2\n}",
		},
		{
			"has arg",
			fields{
				Name:  "height",
				Alias: nil,
				Arguments: []*Argument{
					{Name: "unit", Value: &EnumValue{Value: "FOOT"}},
				},
				Directives: []*Directive{},
				Fields:     []Selection{},
				Typename:   false,
			},
			"height(\n  unit: FOOT\n)",
		},
		{
			"deep",
			fields{
				Name:  "friendsConnection",
				Alias: nil,
				Arguments: []*Argument{
					{Name: "first", Value: &Variable{Name: "first", Type: "Int", Default: nil}},
				},
				Directives: []*Directive{},
				Fields: []Selection{
					&StringField{Value: "totalCount"},
					&Field{
						Name:       "edges",
						Alias:      nil,
						Arguments:  []*Argument{},
						Directives: []*Directive{},
						Fields: []Selection{
							&Field{
								Name:       "node",
								Alias:      nil,
								Arguments:  []*Argument{},
								Directives: []*Directive{},
								Fields:     []Selection{&StringField{Value: "name"}},
								Typename:   false,
							},
						},
						Typename: false,
					},
				},
				Typename: false,
			},
			"friendsConnection(\n  first: $first\n) {\n  totalCount\n  edges {\n    node {\n      name\n    }\n  }\n}",
		},
		{
			"Two directives",
			fields{
				Name:      "friends",
				Fields:    []Selection{&StringField{Value: "name"}},
				Alias:     nil,
				Arguments: []*Argument{},
				Directives: []*Directive{
					{
						Name:      "include",
						Arguments: []*Argument{{Name: "if", Value: &BooleanValue{Value: true}}},
					},
					{
						Name:      "skip",
						Arguments: []*Argument{{Name: "if", Value: &BooleanValue{Value: false}}},
					},
				},
			},
			"friends @include(\n  if: true\n) @skip(\n  if: false\n) {\n  name\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Field{
				Name:       tt.fields.Name,
				Alias:      tt.fields.Alias,
				Arguments:  tt.fields.Arguments,
				Fields:     tt.fields.Fields,
				Directives: tt.fields.Directives,
				Typename:   tt.fields.Typename,
			}
			if got := field.RenderType(); got != tt.want {
				t.Errorf("Field.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
