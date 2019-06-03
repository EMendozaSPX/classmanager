package Models

import "github.com/graphql-go/graphql"

// Golang equivalent to enum of usertype declaration
type Role string
const (
	Admin   Role = "admin"
	Teacher Role = "teacher"
	Student Role = "student"
)

// User struct that serializes into json
type User struct {
	ID       int64    `json:"id"`
	Role     Role   `json:"role"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Login struct serializes into json
type Login struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// Graphql user roles enum type definition
var RoleEnum = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "role",
		Description: "A enum selection of user roles",
		Values: graphql.EnumValueConfigMap{
			"admin": &graphql.EnumValueConfig{
				Value: Admin,
				Description: "The Admin User",
			},
			"teacher": &graphql.EnumValueConfig{
				Value: Teacher,
				Description: "The Teacher User",
			},
			"student": &graphql.EnumValueConfig{
				Value: Student,
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
			"role": &graphql.Field{
				Type: RoleEnum,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: UserType,
			},
		},
	})
