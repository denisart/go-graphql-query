package graphqlQuery

type InlineFragment struct {
	Type      string
	Arguments []Argument
	Fields    []Selection
	Typename  bool
}

func (inlineFragment *InlineFragment) RenderType() string {
	if inlineFragment == nil {
		return ""
	}

	resultString := "... on " + inlineFragment.Type

	if len(inlineFragment.Arguments) > 0 {
		resultString += "(\n"

		for _, arg := range inlineFragment.Arguments {
			resultString += "  " + lineShift(arg.RenderType()) + "\n"
		}

		resultString += ")"
	}

	resultString += " {\n"

	if inlineFragment.Typename {
		resultString += "  __typename\n"
	}

	for _, f := range inlineFragment.Fields {
		resultString += "  " + lineShift(f.renderSelection()) + "\n"
	}

	resultString += "}"

	return resultString
}
