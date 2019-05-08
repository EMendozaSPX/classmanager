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
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
					},
				},
				Resolve: verifyAuthorizationResolver,
			},
			"listUsers": &graphql.Field {
				Type: graphql.NewList(Models.UserType),
				Description: "Get a list users of a certain usertype",
				Args: graphql.FieldConfigArgument{
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
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
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
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
			/*
			"createYearConfiguration": &graphql.Field{
				Type: Models.YearConfigType,
				Description: "Create a year configuration for the current year",
				Args: graphql.FieldConfigArgument{
					"year": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"yearGroup": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"terms": &graphql.ArgumentConfig{
						Type: graphql.NewList(Models.TermInput),
					},
					"publicHolidays": &graphql.ArgumentConfig{
						Type: graphql.NewList(Models.PublicHolidayInput),
					},
					"events": &graphql.ArgumentConfig{
						Type: graphql.NewList(Models.EventInput),
					},
					"periods": &graphql.ArgumentConfig{
						Type: graphql.NewList(Models.PeriodInput),
					},
				},
				Resolve: createYearConfigResolver,
			},
			"createUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Create a new user.",
				Args: graphql.FieldConfigArgument{
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
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

			 */
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
