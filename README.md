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

**go-graphql-query** provides special structures for generate of GraphQL queries. Below are examples of using these classes for queries from GraphQL documentation [https://graphql.org/learn/queries/](https://graphql.org/learn/queries/)

### First query

**Operation** it is the general class for render of your GraphQL query or mutation. For the first query from [https://graphql.org/learn/queries/#fields](https://graphql.org/learn/queries/#fields)

```graphql
{
  hero {
    name
  }
}
```

we can to use **graphqlQuery.Operation** as like that

```go
heroQuery := graphqlQuery.Field{
    Name: "hero",
    Fields: []graphqlQuery.Selection{
        &graphqlQuery.StringField{Value: "name"},
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
    name
  }
}
*/
```

Same way for the query with sub-fields

```graphql
"""
{
  hero {
    name
    # Queries can have comments!
    friends {
      name
    }
  }
}
"""
```

we can to use `graphqlQuery.Selection` as like that

```go
friendsField := graphqlQuery.Field{
    Name: "friends",
    Fields: []graphqlQuery.Selection{
        &graphqlQuery.StringField{Value: "name"},
    },
}

heroQuery := graphqlQuery.Field{
    Name: "hero",
    Fields: []graphqlQuery.Selection{
        &graphqlQuery.StringField{Value: "name"},
        &friendsField,
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
    name
    friends {
      name
    }
  }
}
*/
```

## Arguments

For arguments in your query or fields ([https://graphql.org/learn/queries/#arguments](https://graphql.org/learn/queries/#arguments)) you can use **graphqlQuery.Argument**

```go
idArg := graphqlQuery.Argument{
    Name:  "id",
    Value: &graphqlQuery.StringValue{Value: "1000"},
}
unitArg := graphqlQuery.Argument{
    Name:  "unit",
    Value: &graphqlQuery.EnumValue{Value: "FOOT"},
}

humanQuery := graphqlQuery.Field{
    Name: "human",
    Arguments: []*graphqlQuery.Argument{
        &idArg,
    },
    Fields: []graphqlQuery.Selection{
        &graphqlQuery.StringField{Value: "name"},
        &graphqlQuery.Field{
            Name: "height",
            Arguments: []*graphqlQuery.Argument{
                &unitArg,
            },
        },
    },
}

operation := graphqlQuery.Operation{
    Type:   graphqlQuery.QUERY,
    Name:   nil,
    Fields: []graphqlQuery.Selection{&humanQuery},
}

/*
query {
  human(
    id: "1000"
  ) {
    name
    height(
      unit: FOOT
    )
  }
}
*/
```
