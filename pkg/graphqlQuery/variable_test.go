package graphqlQuery

import "testing"

func TestVariable_RenderType(t *testing.T) {
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
			"Without default",
			fields{
				Name:    "episode",
				Type:    "Episode",
				Default: nil,
			},
			"$episode: Episode",
		},
		{
			"With default",
			fields{
				Name:    "episode",
				Type:    "Episode",
				Default: &VariableDefault{"JEDI"},
			},
			"$episode: Episode = JEDI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			variable := &Variable{
				Name:    tt.fields.Name,
				Type:    tt.fields.Type,
				Default: tt.fields.Default,
			}
			if got := variable.RenderType(); got != tt.want {
				t.Errorf("Variable.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
