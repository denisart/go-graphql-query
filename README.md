# go-graphql-query

Complete Domain Specific Language (DSL) for GraphQL query in go. With this package you can to

- generate a correct GraphQL query string from a go structures;
- use and share similar Arguments, Variables and e.t.c between different queries;
- easily add new fields to your query;
- add Fragments and Directives to queries;

## Quick start

Install package

```bash
$ go get github.com/denisart/go-graphql-query
```

### Simple query

```go
package main

import (
	"fmt"

	"github.com/denisart/go-graphql-query/pkg/graphqlQuery"
)

func main() {
	heroQuery := graphqlQuery.Field{
		Name: "hero",
		Fields: []graphqlQuery.Selection{
			&graphqlQuery.Field{
				Name: "hero",
				Fields: []graphqlQuery.Selection{
					&graphqlQuery.StringField{Value: "name"},
				},
			},
		},
	}

	operation := graphqlQuery.Operation{
		Type:   graphqlQuery.QUERY,
		Name:   nil,
		Fields: []graphqlQuery.Selection{&heroQuery},
	}

	fmt.Println(operation.RenderType())
	/*
		query {
		  hero {
		    hero {
		      name
		    }
		  }
		}
	*/
}
```

## How to use


