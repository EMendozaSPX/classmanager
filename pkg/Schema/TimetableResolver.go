package Schema

/*
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

var viewTimetableResolver = func(params graphql.Params) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var timetable Models.Timetable
	holiday := true

	year, month, day := time.Now().Date()
	yearConfig := Env.GetYearConfig()
	if yearConfig.Year != year {
		for _, term := range yearConfig.Terms {
			_, termStartMonth, termStartDay := term.StartDate.Date()
			_, termEndMonth, termEndDay := term.EndDate.Date()
			if termStartMonth >= month && termEndMonth <= month && termStartDay >= day && termEndDay <= day {
				day := calculateDay(term.StartDate, false)
			} else {
				return nil, nil
			}
		}
	}
	return nil, errors.New("year not configured")
}


 */