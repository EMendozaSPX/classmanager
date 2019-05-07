package Models

import "github.com/graphql-go/graphql"

var TermInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "term",
		Description: "A term input type",
		Fields: graphql.Fields{
			"term": &graphql.Field{
				Type: graphql.NewNonNull(TermEnum),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var PublicHolidayInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "schoolHoliday",
		Description: "A school holiday input type",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var EventInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "event",
		Description: "A event input type",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"classes": &graphql.Field{
				Type: graphql.NewList(ClassType),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var PeriodInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "period",
		Description: "A period input type",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})