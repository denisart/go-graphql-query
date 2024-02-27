package graphqlQuery

import "testing"

func TestInlineFragment_RenderType(t *testing.T) {
	type fields struct {
		Type      string
		Arguments []Argument
		Fields    []Selection
		Typename  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"inlineFragment without typename",
			fields{
				Type:      "name",
				Arguments: []Argument{},
				Fields: []Selection{
					&StringField{Value: "f1"},
					&StringField{Value: "f2"},
				},
				Typename: false,
			},
			"... on name {\n  f1\n  f2\n}",
		},
		{
			"inlineFragment with typename",
			fields{
				Type:      "name",
				Arguments: []Argument{},
				Fields: []Selection{
					&StringField{Value: "f1"},
					&StringField{Value: "f2"},
				},
				Typename: true,
			},
			"... on name {\n  __typename\n  f1\n  f2\n}",
		},
		{
			"inlineFragment with arguments",
			fields{
				Type: "name",
				Arguments: []Argument{
					{Name: "arg1", Value: &EnumValue{Value: "VALUE1"}},
					{Name: "arg2", Value: &EnumValue{Value: "VALUE2"}},
				},
				Fields: []Selection{
					&StringField{Value: "f1"},
					&StringField{Value: "f2"},
				},
				Typename: true,
			},
			"... on name(\n  arg1: VALUE1\n  arg2: VALUE2\n) {\n  __typename\n  f1\n  f2\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inlineFragment := &InlineFragment{
				Type:      tt.fields.Type,
				Arguments: tt.fields.Arguments,
				Fields:    tt.fields.Fields,
				Typename:  tt.fields.Typename,
			}
			if got := inlineFragment.RenderType(); got != tt.want {
				t.Errorf("InlineFragment.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
