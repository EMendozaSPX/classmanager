package Schema

import (
	"database/sql"
	"github.com/emendoza/classmanager/pkg/Database"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var (
	Schema graphql.Schema
	db     *sql.DB
)

func init() {
	// create a error variable to handle errors
	var err error

	// get db instance from file
	db = Database.GetDB()

	// create a root query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"verifyAuthorization": &graphql.Field {
				Type: graphql.Boolean,
				Description: "Provides a user access to a particular site",
				Args: graphql.FieldConfigArgument{
					"usertype": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.UserTypeEnum),
					},
				},
				Resolve: verifyAuthorizationResolver,
			},
			"listUsers": &graphql.Field {
				Type: graphql.NewList(Models.UserType),
				Description: "Get a list users of a certain usertype",
				Args: graphql.FieldConfigArgument{
					"usertype": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.UserTypeEnum),
					},
				},
				Resolve: listUsersResolver,
			},
		},
	})


	// Creates a graphql mutation type
	// graphql mutation types take arguments and return corresponding information
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: graphql.String,
				Description: "Sign in users through json web tokens.",
				Args: graphql.FieldConfigArgument{
					"usertype": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.UserTypeEnum),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: loginResolver,
			},
			"createYearConfiguration": &graphql.Field{
				Type: Models.
			}
			"createUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Create a new user.",
				Args: graphql.FieldConfigArgument{
					"usertype": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.UserTypeEnum),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createUserResolver,
			},
		},
	})

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
