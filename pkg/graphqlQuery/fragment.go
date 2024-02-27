package graphqlQuery

type Fragment struct {
	Name     string
	Type     string
	Fields   []Selection
	Typename bool
}

func (fragment *Fragment) RenderType() string {
	if fragment == nil {
		return ""
	}

	resultString := "fragment " + fragment.Name + " on " + fragment.Type + " {\n"

	if fragment.Typename {
		resultString += "  __typename\n"
	}

	for _, f := range fragment.Fields {
		resultString += "  " + lineShift(f.renderSelection()) + "\n"
	}

	resultString += "}"

	return resultString
}
