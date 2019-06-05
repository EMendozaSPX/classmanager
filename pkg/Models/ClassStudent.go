package Models

import "github.com/graphql-go/graphql"

type ClassStudent struct {
	ID             int             `json:"id"`
	Student        User            `json:"student"`
	BehaviourNotes []BehaviourNote `json:"behaviourNotes"`
	TaskMarks      []TaskMark      `json:"taskMarks"`
}

var ClassStudentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClassStudent",
		Description: "Class Student Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"student": &graphql.Field{
				Type: UserType,
			},
			"behaviourNotes": &graphql.Field{
				Type: BehaviourNoteType,
			},
			"taskMarks": &graphql.Field{
				Type: TaskMarkType,
			},
		},
	})
