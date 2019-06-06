package Models

import "github.com/graphql-go/graphql"

type Class struct {
	ID       int            `json:"id"`
	ClassId  string         `json:"classID"`
	Teacher  User           `json:"teacher"`
	Students []User         `json:"students"`
	Tasks    []Task         `json:"tasks"`
}

var ClassType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Class",
		Description: "Class Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"classId": &graphql.Field{
				Type: graphql.String,
			},
			"teacher": &graphql.Field{
				Type: UserType,
			},
			"students": &graphql.Field{
				Type: graphql.NewList(UserType),
			},
		},
	})