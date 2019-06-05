package Models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TotalMarks  int       `json:"totalMarks"`
	Description string    `json:"description"`
	DueTime     time.Time `json:"dueTime"`
}

type TaskMark struct {
	ID           int       `json:"id"`
	AssignedTask Task      `json:"task"`
	TaskMark     int       `json:"taskMark"`
	Feedback     string    `json:"feedback"`
	TimeStamp    time.Time `json:"timeStamp"`
}

var TaskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
		Description: "Task Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"totalMarks": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"dueDate": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

var TaskMarkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TaskMark",
		Description: "Task Mark Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"task": &graphql.Field{
				Type: TaskType,
			},
			"taskMark": &graphql.Field{
				Type: graphql.Int,
			},
			"feedback": &graphql.Field{
				Type: graphql.String,
			},
			"timeStamp": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})
