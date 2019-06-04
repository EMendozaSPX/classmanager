package Models

import (
	"github.com/graphql-go/graphql"
)

type Weekday int
const (
	Weekend    Weekday = 0
	Monday1    Weekday = 1
	Tuesday1   Weekday = 2
	Wednesday1 Weekday = 3
	Thursday1  Weekday = 4
	Friday1    Weekday = 5
	Monday2    Weekday = 6
	Tuesday2   Weekday = 7
	Wednesday2 Weekday = 8
	Thursday2  Weekday = 9
	Friday2    Weekday = 10
)

// A struct that serializes into a timetable entry
type Timetable struct {
	Period    string `json:"period"`
	Class     string `json:"class"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// graphql timetable type declaration
var TimetableType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "timetable",
		Description: "A timetable entry object type",
		Fields: graphql.Fields{
			"period": &graphql.Field{
				Type: graphql.String,
			},
			"class": &graphql.Field{
				Type: graphql.String,
			},
			"startTime": &graphql.Field{
				Type: graphql.String,
			},
			"endTime": &graphql.Field{
				Type: graphql.String,
			},
		},
	})