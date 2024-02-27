package main

import (
	"fmt"

	"github.com/denisart/go-graphql-query/pkg/graphqlQuery"
)

func main() {
	fmt.Println("Hello to go-graphql-query")

	argValue := graphqlQuery.StringValue{Value: "123"}
	arg1 := graphqlQuery.Argument{
		Name:  "id",
		Value: &argValue,
	}

	fmt.Println(arg1.RenderType())
}
