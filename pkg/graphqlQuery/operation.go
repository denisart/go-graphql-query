package graphqlQuery

type OperationType string

const (
	QUERY        OperationType = "query"
	MUTATION     OperationType = "mutation"
	SUBSCRIPTION OperationType = "subscription"
)

type OperationName struct {
	Value string
}

type Operation struct {
	Type      OperationType
	Name      *OperationName
	Variables []Variable
	Fields    []Selection
	Fragments []Fragment
}

func (operation *Operation) RenderType() string {
	if operation == nil {
		return ""
	}

	resultString := string(operation.Type)

	if operation.Name != nil {
		resultString += " " + operation.Name.Value
	}

	if len(operation.Variables) > 0 {
		resultString += "(\n"

		for _, v := range operation.Variables {
			resultString += "  " + v.RenderType() + "\n"
		}

		resultString += ")"
	}

	resultString += " {\n"

	for _, f := range operation.Fields {
		resultString += "  " + lineShift(f.renderSelection()) + "\n"
	}

	resultString += "}"

	if len(operation.Fragments) > 0 {
		resultString += "\n\n"

		for _, f := range operation.Fragments {
			resultString += f.RenderType() + "\n"
		}
	}

	return resultString
}
