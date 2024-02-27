package graphqlQuery

// an abstract type with `RenderType` method
type GraphQLQueryType interface {
	RenderType() string
}
