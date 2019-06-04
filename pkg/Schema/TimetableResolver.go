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

/*
// Query Classes in the given weekday
var timetableQuery = `
SELECT timetable.period_name, classes.class_id, classes.teacher_id
FROM timetable
INNER JOIN classes
ON timetable.class_id=classes.id
WHERE week_day=$1
`

// calculate the day of the week
// recursive function will recurse through each day of the week starting
// from the start of the term
// turns a golang weekday into a bi-weekday
func calculateDay(startTime time.Time, biweek bool) Models.Weekday {
	day := startTime.Weekday()
	var biWeekDay Models.Weekday
	switch day {
	case time.Monday:
		if biweek {
			biWeekDay = Models.Monday2
		} else {
			biWeekDay = Models.Monday1
		}
		break
	case time.Tuesday:
		if biweek {
			biWeekDay = Models.Tuesday2
		} else {
			biWeekDay = Models.Tuesday1
		}
		break
	case time.Wednesday:
		if biweek {
			biWeekDay = Models.Wednesday2
		} else {
			biWeekDay = Models.Wednesday1
		}
		break
	case time.Thursday:
		if biweek {
			biWeekDay = Models.Thursday2
		} else {
			biWeekDay = Models.Thursday1
		}
		break
	case time.Friday:
		if biweek {
			biWeekDay = Models.Friday2
		} else {
			biWeekDay = Models.Friday1
		}
	default:
		biWeekDay = Models.Weekend
		break
	}
	if startTime == time.Now() {
		return biWeekDay
	}
	return calculateDay(startTime.AddDate(0, 0, 1), !biweek)
}

// resolver function to view timetable query
var viewTimetableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// teachers authorization
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	// variable to hold timetable entries
	var timetables []Models.Timetable

	// seperate time into day month and year variables
	year, month, day := time.Now().Date()

	// get year config
	yearConfig := Env.GetYearConfig()

	// verify that year is configured for the current year
	if yearConfig.Year != year {
		return nil, errors.New("year not configured")
	}

	// loop through periods in year config and add to timetables
	for _, period := range Env.GetYearConfig().Periods {
		timetable := Models.Timetable{
			Period: period.PeriodName,
			Class: "Free Period",
			StartTime: period.StartTime,
			EndTime: period.EndTime,
		}
		timetables = append(timetables, timetable)
	}

	// create a weekday variable to hold the bi-weekday
	var weekDay Models.Weekday

	// loop through terms in year config and verify if this date is in year
	for _, term := range yearConfig.Terms {
		_, termStartMonth, termStartDay := term.StartDate.Date()
		_, termEndMonth, termEndDay := term.EndDate.Date()
		if termStartMonth >= month && termEndMonth <= month && termStartDay >= day && termEndDay <= day {
			// calculate bi-weekday and save to weekday variable for later use
			weekDay = calculateDay(term.StartDate, false)

			// return nothing if is during weekend
			if weekDay == Models.Weekend {
				return nil, nil
			}
		} else {
			return nil, nil
		}
	}

	// query timetable information from database
	rows, err := db.Query(timetableQuery, weekDay)
	if err != nil {
		log.Println(err)
	}

	// loop through timetable entries in database
	for rows.Next() {
		var (
			periodName string
			className string
			teacherId int
			)

		// save to variables
		if err := rows.Scan(&periodName, &className, &teacherId); err != nil {
			log.Println(err)
		}

		// checks if the timetable entry applies to current teacher
		if teacherId == params.Args["teacherId"] {
			for i, timetable := range timetables {
				if timetable.Period == periodName {
					// save information as a result in timetable
					timetables[i].Class = className
				}
			}
		}
	}
	return timetables, nil
}

 */