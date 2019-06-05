package Models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type BehaviourNote struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Note           string    `json:"note"`
	TimeStamp      time.Time `json:"timeStamp"`
}

var BehaviourNoteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BehaviourNote",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"note": &graphql.Field{
				Type: graphql.String,
			},
			"timeStamp": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})