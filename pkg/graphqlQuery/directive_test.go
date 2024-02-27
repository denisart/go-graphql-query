package graphqlQuery

import "testing"

func TestDirective_RenderType(t *testing.T) {
	type fields struct {
		Name      string
		Arguments []*Argument
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"skip directive",
			fields{
				Name:      "skip",
				Arguments: []*Argument{},
			},
			"@skip",
		},
		{
			"if directive",
			fields{
				Name:      "skip",
				Arguments: []*Argument{{Name: "if", Value: &BooleanValue{Value: true}}},
			},
			"@skip(\n  if: true\n)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			directive := &Directive{
				Name:      tt.fields.Name,
				Arguments: tt.fields.Arguments,
			}
			if got := directive.RenderType(); got != tt.want {
				t.Errorf("Directive.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
