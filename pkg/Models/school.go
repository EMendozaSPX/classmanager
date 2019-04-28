package Models

import "github.com/graphql-go/graphql"

// School configuration for the year per year group struct
type YearConfig struct {
	Year           int             `json:"year"`
	YearGroup      int             `json:"yearGroup"`
	Terms          []Term          `json:"term"`
	PublicHolidays []PublicHoliday `json:"publicHolidays"`
	Periods        []Period        `json:"periods"`
}

// School Term struct
type Term struct {
	Term      int    `json:"term"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// Public Holidays struct
type PublicHoliday struct {
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
// Event struct for incursions, excursions and
type Event struct {
	Name            string  `json:"name"`
	Classes         []Class `json:"classes"`
	StartDateTime   string  `json:"startDateTime"`
	EndDateTime     string  `json:"startDateTime"`
}
// Periods struct
type Period struct {
	Name      string `json:"name"`
	BeginTime int    `json:"beginTime"`
	EndTime   int    `json:"endTime"`
}

var TermType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "term",
		Description: "Configuration for a term at school",
		Fields: graphql.Fields{
			"term": &graphql.Field{
				Type: graphql.String,
			},
			"startDate": &graphql.Field{
				Type: graphql.String,
			},
			"endDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var PublicHolidayType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "publicHoliday",
		Description: "Configuration for public holidays",
		Fields: graphql.Fields{
			""
		}
	})
