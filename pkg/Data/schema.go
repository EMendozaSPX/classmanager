package Data

import (
	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func init() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"AdminLogin": &graphql.Field{
				Type: loginAdminType,
				Resolve:
			}
		}
	})
}