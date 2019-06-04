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

// A struct that serializes into a period entry
type Period struct {
	PeriodName    string `json:"periodName"`
	Class         string `json:"class"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}

type Timetable struct {
	Weekday Weekday  `json:"weekday"`
	Periods []Period `json:"periods"`
}

// graphql timetable type declaration
var PeriodType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "period",
		Description: "A period entry object type",
		Fields: graphql.Fields{
			"periodName": &graphql.Field{
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

var TimetableType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "timetable",
		Description: "A timetble entry object",
		Fields: graphql.Fields{
			"weekday": &graphql.Field{
				Type: graphql.String,
			},
			"periods": &graphql.Field{
				Type: graphql.NewList(PeriodType),
			},
		},
	})