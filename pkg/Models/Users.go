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
	ID       int    `json:"id"`
	Role     Role   `json:"role"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ClassStudent struct {
	ID          int  `json:"id"`
	StudentInfo User `json:"studentInfo"`
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

var ClassStudentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClassStudent",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"studentInfo": &graphql.Field{
				Type: UserType,
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
