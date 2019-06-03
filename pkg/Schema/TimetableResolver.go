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

var timetableQuery = `
SELECT class_id, period_name
FROM timetable
WHERE week_day=$1
`

var classIdQuery = `
SELECT class_id, teacher_id
FROM classes
WHERE id=$1
`

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

var viewTimetableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var timetables []Models.Timetable
	year, month, day := time.Now().Date()
	yearConfig := Env.GetYearConfig()

	var weekDay Models.Weekday

	if yearConfig.Year == year {
		for _, period := range Env.GetYearConfig().Periods {
			timetable := Models.Timetable{
				Period: period.PeriodName,
				Class: "Free Period",
				StartTime: period.StartTime,
				EndTime: period.EndTime,
			}
			timetables = append(timetables, timetable)
		}
		for _, term := range yearConfig.Terms {
			_, termStartMonth, termStartDay := term.StartDate.Date()
			_, termEndMonth, termEndDay := term.EndDate.Date()
			if termStartMonth >= month && termEndMonth <= month && termStartDay >= day && termEndDay <= day {
				weekDay = calculateDay(term.StartDate, false)
			} else {
				return nil, nil
			}
		}
		rows, err := db.Query(timetableQuery, weekDay)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			var (
				classId int
				periodName string
			)

			if err := rows.Scan(&classId, &periodName); err != nil {
				log.Println(err)
			}

			classRows, err := db.Query(classIdQuery, classId)
			if err != nil {
				log.Println(err)
			}
			for classRows.Next() {
				var (
					className string
					teacherId int
				)

				if err := classRows.Scan(&className, &teacherId); err != nil {
					log.Println(err)
				}

				if teacherId == params.Args["teacherId"] {
					for i, timetable := range timetables {
						if timetable.Period == periodName {
							timetables[i].Class = className
						}
					}
				}
			}
		}
		return timetables, nil
	}
	return nil, errors.New("year not configured")
}