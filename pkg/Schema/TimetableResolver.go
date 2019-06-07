package Schema

import (
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Env"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
	"time"
)


// Query Classes in the given weekday
var timetableQuery = `
SELECT timetable.period_name, classes.class_id
FROM timetable
INNER JOIN classes
ON timetable.class_id=classes.id
WHERE timetable.week_day= $1 AND classes.teacher_id=$2;
`

// resolver function to view timetable query
var viewTimetableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// teachers authorization
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	// variable to hold week days
	var weekdays []Models.Weekday

	// variable to hold period entries
	var periods []Models.Period

	// variable to hold classes entries
	var classesList [][]string

	// seperate time into day month and year variables
	year, _, _ := time.Now().Date()

	// get year config
	yearConfig := Env.GetYearConfig()

	// verify that year is configured for the current year
	if yearConfig.Year != year {
		return nil, errors.New("year not configured")
	}

	for _, EnvPeriod := range Env.GetYearConfig().Periods {
		period := Models.Period{
			PeriodName: EnvPeriod.PeriodName,
			StartTime: EnvPeriod.StartTime,
			EndTime: EnvPeriod.EndTime,
		}
		periods = append(periods, period)
	}

	for w := Models.Monday1; w <= Models.Friday2; w++ {
		var classes []string

		// query classes in timetable
		rows, err := db.Query(timetableQuery, w, params.Args["teacherId"].(int))
		if err != nil {
			log.Println(err)
		}

		for rows.Next() {
			var (
				periodName string
				className  string
			)
			if err := rows.Scan(&periodName, &className); err != nil {
				log.Println(err)
			}

			for _, confPeriod := range Env.GetYearConfig().Periods {
				var class string

				// create a class variable
				if periodName == confPeriod.PeriodName {
					class = className
				} else {
					class = "Free Period"
				}

				// add variables to sub lists
				classes = append(classes, class)
			}
		}
		if len(classes) <= 0 {
			for i := 0; i < len(Env.GetYearConfig().Periods); i++ {
				classes = append(classes,"Free Period")
			}
		}

		// add variables to lists
		weekdays = append(weekdays, w)
		classesList = append(classesList, classes)
	}

	// create timetable struct
	timetable := Models.Timetable{
		Weekdays: weekdays,
		Periods: periods,
		Classes: classesList,
	}

	return timetable, nil
}