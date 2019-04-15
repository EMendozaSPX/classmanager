package Data

import "github.com/graphql-go/graphql"

// Admin user struct that serializes into json
type Admin struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Teacher user struct that serializes into json
type Teacher struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Student user struct that serializes into json
type Student struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}



var adminType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Admins",
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
	},
	)