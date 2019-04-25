package Data

import (
	"database/sql"
	"fmt"
	"github.com/emendoza/classmanager/pkg/Env"
	"github.com/graphql-go/graphql"
	"log"

	_ "github.com/lib/pq"
)

var (
	Schema graphql.Schema
	db     *sql.DB
)


func init() {
	// create a error variable to handle errors
	var err error

	// Setup database connection
	// get database user information from the env.json file
	dbUser := Env.GetDatabaseUser()

	// create a postgres database config string using the user information from the previous line
	connStr := fmt.Sprintf("user=%v password=%v dbname=classmanager port=5433 sslmode=disable",
		dbUser.Username, dbUser.Password)

	// open sql database with the set configuration
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}

	// defer closing the database to the end of the program
	// defer db.Close()

	// create a root query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"listAdmins": &graphql.Field {
				Type: graphql.NewList(adminType),
				Description: "Get a list of admin users",
				Resolve: listAdminsResolver,
			},
		},
	})


	// Creates a graphql mutation type
	// graphql mutation types take arguments and return corresponding information
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"signin": &graphql.Field{
				Type: graphql.String,
				Description: "Sign in users through json web tokens.",
				Args: graphql.FieldConfigArgument{
					"userType": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(userTypeEnum),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: authenticateUser,
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
