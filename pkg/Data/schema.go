package Data

import (
	"github.com/graphql-go/graphql"
	"log"
)

var Schema graphql.Schema

func init() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			Type: graphql.
		}
	})


	// Creates a graphql mutation type
	// graphql mutation types take arguments and return corresponding information
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"signinAdmin": &graphql.Field{
				Type: graphql.String,
				Description: "sign in admin users through json web tokens",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: authenticateAdmin,
			},
		},
	})

	// create an error variable
	var err error

	// create a new graphql schema using the query and mutation types, if failed returns an error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
			Mutation: mutationType,
		})

	// print error to console
	if err != nil {
		log.Println(err)
	}
}