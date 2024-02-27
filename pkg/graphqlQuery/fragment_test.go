package graphqlQuery

import "testing"

func TestFragment_RenderType(t *testing.T) {
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
			"fragment without typename",
			fields{
				Name: "comparisonFields",
				Type: "Character",
				Fields: []Selection{
					&StringField{Value: "name"},
					&StringField{Value: "appearsIn"},
				},
				Typename: false,
			},
			"fragment comparisonFields on Character {\n  name\n  appearsIn\n}",
		},
		{
			"fragment with typename",
			fields{
				Name: "comparisonFields",
				Type: "Character",
				Fields: []Selection{
					&StringField{Value: "name"},
					&StringField{Value: "appearsIn"},
				},
				Typename: true,
			},
			"fragment comparisonFields on Character {\n  __typename\n  name\n  appearsIn\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fragment := &Fragment{
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				Fields:   tt.fields.Fields,
				Typename: tt.fields.Typename,
			}
			if got := fragment.RenderType(); got != tt.want {
				t.Errorf("Fragment.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
