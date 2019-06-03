package Models

import (
	"github.com/graphql-go/graphql"
	"time"
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

type Timetable struct {
	Period    string `json:"period"`
	Class     string `json:"class"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
// Public Holidays struct
type PublicHoliday struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// Event struct for incursions, excursions and
type Event struct {
	Name            string    `json:"name"`
	Classes         []Class   `json:"classes"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
}

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

var publicHolidayType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "publicHoliday",
		Description: "Public holidays object type",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"startTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"endTime": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

var eventType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "event",
		Description: "Excursions and incursions, or general events of school",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"classes": &graphql.Field{
				Type: graphql.NewList(ClassType),
			},
			"startTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"endTime": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})