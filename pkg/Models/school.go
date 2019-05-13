package Models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type TermID string
const (
	Term1 TermID = "term1"
	Term2 TermID = "term2"
	Term3 TermID = "term3"
	Term4 TermID = "term4"
)

var TermEnum = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "enumID",
		Description: "A enum of possible terms",
		Values: graphql.EnumValueConfigMap{
			"term1": &graphql.EnumValueConfig{
				Value: Term1,
				Description: "Term one",
			},
			"term2": &graphql.EnumValueConfig{
				Value: Term2,
				Description: "Term two",
			},
			"term3": &graphql.EnumValueConfig{
				Value: Term3,
				Description: "Term three",
			},
			"term4": &graphql.EnumValueConfig{
				Value: Term4,
				Description: "Term four",
			},
		},
	})

// School configuration for the year per year group struct
type YearConfig struct {
	ID             int             `json:"id"`
	Year           int             `json:"year"`
	YearGroup      int             `json:"yearGroup"`
	Terms          []Term          `json:"terms"`
	PublicHolidays []PublicHoliday `json:"publicHolidays"`
	Events         []Event         `json:"events"`
	Periods        []Period        `json:"periods"`
}

// School Term struct
type Term struct {
	Name      TermID    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
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

// Periods struct
type Period struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

var termType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "term",
		Description: "School Term object type",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: TermEnum,
			},
			"startTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"endTime": &graphql.Field{
				Type: graphql.DateTime,
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

var periodType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "period",
		Description: "A period object type",
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

var YearConfigType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "yearConfig",
		Description: "A year configuration object type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"yearGroup": &graphql.Field{
				Type: graphql.Int,
			},
			"terms": &graphql.Field{
				Type: graphql.NewList(termType),
			},
			"publicHolidays": &graphql.Field{
				Type: graphql.NewList(publicHolidayType),
			},
			"events": &graphql.Field{
				Type: graphql.NewList(eventType),
			},
			"periods": &graphql.Field{
				Type: graphql.NewList(periodType),
			},
		},
	})