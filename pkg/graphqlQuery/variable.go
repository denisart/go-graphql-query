package graphqlQuery

type VariableDefault struct {
	Value string
}

type Variable struct {
	Name    string
	Type    string
	Default *VariableDefault
}

func (variable *Variable) RenderType() string {
	result := "$" + variable.Name + ": " + variable.Type

	if variable.Default != nil {
		result = result + " = " + variable.Default.Value
	}

	return result
}
