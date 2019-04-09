package Data

import "github.com/graphql-go/graphql"

// Admin user struct that serializes into json
type Admin struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Teacher user struct that serializes into json
type Teacher struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Student user struct that serializes into json
type Student struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create a graphql object called Login Admin that stores a username and password
var loginAdminType = graphql.NewObject(
	graphql.ObjectConfig{
	    Name: "LoginAdmin",
	    Fields: graphql.Fields{
		    "username": &graphql.Field{
			    Type: graphql.String,
		    },
		    "password": &graphql.Field{
			    Type: graphql.String,
		    },
	    },
    },
)