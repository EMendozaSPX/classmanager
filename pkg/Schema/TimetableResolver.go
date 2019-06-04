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
SELECT classes.class_id
FROM timetable
INNER JOIN classes
ON timetable.class_id=classes.id
WHERE (timetable.period_name=$1 AND timetable.week_day=$2 AND classes.teacher_id=$3);
`

// resolver function to view timetable query
var viewTimetableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// teachers authorization
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var timetables []Models.Timetable

	// variable to hold timetable entries
	var periods []Models.Period

	// seperate time into day month and year variables
	year, _, _ := time.Now().Date()

	// get year config
	yearConfig := Env.GetYearConfig()

	// verify that year is configured for the current year
	if yearConfig.Year != year {
		return nil, errors.New("year not configured")
	}

	for w := Models.Monday1; w <= Models.Friday2; w++ {
		// loop through periods in year config and add to periods
		for _, confPeriod := range Env.GetYearConfig().Periods {
			var class string
			err := db.QueryRow(timetableQuery, confPeriod.PeriodName, w, params.Args["teacherId"].(int)).Scan(&class)
			if err != nil {
				log.Println(err)
			}

			if class == "" {
				class = "Free Period"
			}

			period := Models.Period{
				PeriodName: confPeriod.PeriodName,
				Class: class,
				StartTime: confPeriod.StartTime,
				EndTime: confPeriod.EndTime,
			}
			periods = append(periods, period)

		}

		timetable := Models.Timetable{
			Weekday: w,
			Periods: periods,
		}

		timetables = append(timetables, timetable)
	}

	return timetables, nil
}