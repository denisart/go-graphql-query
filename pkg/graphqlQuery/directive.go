package graphqlQuery

type Directive struct {
	Name      string
	Arguments []*Argument
}

func (directive *Directive) RenderType() string {
	directiveString := "@" + directive.Name

	if len(directive.Arguments) > 0 {
		directiveString += "("

		for _, arg := range directive.Arguments {
			directiveString += "\n  " + arg.RenderType()
		}

		directiveString += "\n)"
	}

	return directiveString
}
