package Models

import "github.com/graphql-go/graphql"

var TermInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "termInput",
		Description: "A term input type",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(TermEnum),
			},
			"startTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var PublicHolidayInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "publicHolidayInput",
		Description: "A public holiday input type",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var EventInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "eventInput",
		Description: "A event input type",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"classes": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(ClassType),
			},
			"startTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})

var PeriodInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "periodInput",
		Description: "A period input type",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"endTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},
	})