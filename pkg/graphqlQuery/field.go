package graphqlQuery

type Selection interface {
	renderSelection() string
}

var _ Selection = (*StringField)(nil)
var _ Selection = (*Field)(nil)
var _ Selection = (*InlineFragment)(nil)
var _ Selection = (*Fragment)(nil)

type StringField struct {
	Value string
}

func (v *StringField) renderSelection() string {
	return v.Value
}

func (v *Field) renderSelection() string {
	return v.RenderType()
}

func (v *InlineFragment) renderSelection() string {
	return lineShift(v.RenderType())
}

func (v *Fragment) renderSelection() string {
	return "..." + v.Name
}

type Field struct {
	Name       string
	Alias      *Alias
	Arguments  []*Argument
	Fields     []Selection
	Directives []*Directive
	Typename   bool
}

func (field *Field) RenderType() string {
	if field == nil {
		return ""
	}

	resultString := ""

	if field.Alias != nil {
		resultString += field.Alias.Value + ": "
	}
	resultString += field.Name

	if len(field.Arguments) > 0 {
		resultString += "(\n"

		for _, arg := range field.Arguments {
			resultString += "  " + lineShift(arg.RenderType()) + "\n"
		}

		resultString += ")"
	}

	if len(field.Directives) > 0 {
		for _, d := range field.Directives {
			resultString += " " + d.RenderType()
		}
	}

	if len(field.Fields) > 0 || field.Typename {
		resultString += " {\n"

		if field.Typename {
			resultString += "  __typename\n"
		}

		for _, f := range field.Fields {
			resultString += "  " + lineShift(f.renderSelection()) + "\n"
		}

		resultString += "}"
	}

	return resultString
}
