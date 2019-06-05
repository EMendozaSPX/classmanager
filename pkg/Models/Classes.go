package Models

import "github.com/graphql-go/graphql"

type Class struct {
	ID       int            `json:"id"`
	ClassID  string         `json:"classID"`
	Teacher  User           `json:"teacher"`
	Students []ClassStudent `json:"students"`
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
			"classID": &graphql.Field{
				Type: graphql.String,
			},
			"teacher": &graphql.Field{
				Type: UserType,
			},
			"students": &graphql.Field{
				Type: graphql.NewList(ClassStudentType),
			},
			"tasks": &graphql.Field{
				Type: graphql.NewList(TaskType),
			},
		},
	})