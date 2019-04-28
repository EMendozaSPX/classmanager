package Models

import "github.com/graphql-go/graphql"

// User struct that serializes into json
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Graphql user enum type definition
var UserTypeEnum = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "usertype",
		Description: "A enum selection of user types",
		Values: graphql.EnumValueConfigMap{
			"admin": &graphql.EnumValueConfig{
				Value: 1,
				Description: "The Admin User",
			},
			"teacher": &graphql.EnumValueConfig{
				Value: 2,
				Description: "The Teacher User",
			},
			"student": &graphql.EnumValueConfig{
				Value: 3,
				Description: "The Student User",
			},
		},
	})

// graphql User Type definition
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
