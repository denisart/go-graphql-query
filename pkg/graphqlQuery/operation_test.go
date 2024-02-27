package graphqlQuery

import "testing"

func TestOperation_RenderType(t *testing.T) {
	type fields struct {
		Type      OperationType
		Name      *OperationName
		Variables []Variable
		Fields    []Selection
		Fragments []Fragment
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simple operation",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name:   "hero",
						Fields: []Selection{&StringField{Value: "name"}},
					},
				},
				Fragments: []Fragment{},
			},
			"query {\n  hero {\n    name\n  }\n}",
		},
		{
			"hero friends query",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name: "hero",
						Fields: []Selection{
							&StringField{Value: "name"},
							&Field{
								Name:   "friends",
								Fields: []Selection{&StringField{Value: "name"}},
							},
						},
					},
				},
				Fragments: []Fragment{},
			},
			"query {\n  hero {\n    name\n    friends {\n      name\n    }\n  }\n}",
		},
		{
			"human query",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name: "human",
						Arguments: []*Argument{
							{Name: "id", Value: &StringValue{Value: "1000"}},
						},
						Fields: []Selection{
							&StringField{Value: "name"},
							&StringField{Value: "height"},
						},
					},
				},
				Fragments: []Fragment{},
			},
			"query {\n  human(\n    id: \"1000\"\n  ) {\n    name\n    height\n  }\n}",
		},
		{
			"human query with field argument",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name: "human",
						Arguments: []*Argument{
							{Name: "id", Value: &StringValue{Value: "1000"}},
						},
						Fields: []Selection{
							&StringField{Value: "name"},
							&Field{
								Name: "height",
								Arguments: []*Argument{
									{Name: "unit", Value: &EnumValue{Value: "FOOT"}},
								},
								Fields: []Selection{},
							},
						},
					},
				},
				Fragments: []Fragment{},
			},
			"query {\n  human(\n    id: \"1000\"\n  ) {\n    name\n    height(\n      unit: FOOT\n    )\n  }\n}",
		},
		{
			"multyqueries",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name:  "hero",
						Alias: &Alias{Value: "empireHero"},
						Arguments: []*Argument{
							{Name: "episode", Value: &EnumValue{Value: "EMPIRE"}},
						},
						Fields: []Selection{&StringField{Value: "name"}},
					},
					&Field{
						Name:  "hero",
						Alias: &Alias{Value: "jediHero"},
						Arguments: []*Argument{
							{Name: "episode", Value: &EnumValue{Value: "JEDI"}},
						},
						Fields: []Selection{&StringField{Value: "name"}},
					},
				},
			},
			"query {\n  empireHero: hero(\n    episode: EMPIRE\n  ) {\n    name\n  }\n  jediHero: hero(\n    episode: JEDI\n  ) {\n    name\n  }\n}",
		},
		{
			"with fragment",
			fields{
				Type:      QUERY,
				Name:      nil,
				Variables: []Variable{},
				Fields: []Selection{
					&Field{
						Name:  "hero",
						Alias: &Alias{Value: "leftComparison"},
						Arguments: []*Argument{
							{Name: "episode", Value: &EnumValue{Value: "EMPIRE"}},
						},
						Fields: []Selection{
							&Fragment{
								Name: "comparisonFields",
								Type: "Character",
								Fields: []Selection{
									&StringField{Value: "name"},
									&StringField{Value: "appearsIn"},
									&Field{
										Name:   "friends",
										Fields: []Selection{&StringField{Value: "name"}},
									},
								},
							},
						},
					},
					&Field{
						Name:  "hero",
						Alias: &Alias{Value: "rightComparison"},
						Arguments: []*Argument{
							{Name: "episode", Value: &EnumValue{Value: "JEDI"}},
						},
						Fields: []Selection{
							&Fragment{
								Name: "comparisonFields",
								Type: "Character",
								Fields: []Selection{
									&StringField{Value: "name"},
									&StringField{Value: "appearsIn"},
									&Field{
										Name:   "friends",
										Fields: []Selection{&StringField{Value: "name"}},
									},
								},
							},
						},
					},
				},
				Fragments: []Fragment{
					{
						Name: "comparisonFields",
						Type: "Character",
						Fields: []Selection{
							&StringField{Value: "name"},
							&StringField{Value: "appearsIn"},
							&Field{
								Name:   "friends",
								Fields: []Selection{&StringField{Value: "name"}},
							},
						},
					},
				},
			},
			"query {\n  leftComparison: hero(\n    episode: EMPIRE\n  ) {\n    ...comparisonFields\n  }\n  rightComparison: hero(\n    episode: JEDI\n  ) {\n    ...comparisonFields\n  }\n}\n\nfragment comparisonFields on Character {\n  name\n  appearsIn\n  friends {\n    name\n  }\n}\n",
		},
		{
			"CreateReviewForEpisode",
			fields{
				Type: MUTATION,
				Name: &OperationName{Value: "CreateReviewForEpisode"},
				Variables: []Variable{
					{Name: "ep", Type: "Episode!", Default: nil},
					{Name: "review", Type: "ReviewInput!", Default: nil},
				},
				Fields: []Selection{
					&Field{
						Name: "createReview",
						Arguments: []*Argument{
							{Name: "episode", Value: &Variable{Name: "ep", Type: "Episode!", Default: nil}},
							{Name: "review", Value: &Variable{Name: "review", Type: "ReviewInput!", Default: nil}},
						},
						Fields: []Selection{
							&StringField{Value: "stars"},
							&StringField{Value: "commentary"},
						},
					},
				},
			},
			"mutation CreateReviewForEpisode(\n  $ep: Episode!\n  $review: ReviewInput!\n) {\n  createReview(\n    episode: $ep\n    review: $review\n  ) {\n    stars\n    commentary\n  }\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := &Operation{
				Type:      tt.fields.Type,
				Name:      tt.fields.Name,
				Variables: tt.fields.Variables,
				Fields:    tt.fields.Fields,
				Fragments: tt.fields.Fragments,
			}
			if got := operation.RenderType(); got != tt.want {
				t.Errorf("Operation.RenderType() = %v, want %v", got, tt.want)
			}
		})
	}
}
