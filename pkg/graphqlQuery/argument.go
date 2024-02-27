package graphqlQuery

type Argument struct {
	Name  string
	Value ArgumentValue
}

func (argument *Argument) RenderType() string {
	return argument.Name + ": " + argument.Value.renderValue()
}
