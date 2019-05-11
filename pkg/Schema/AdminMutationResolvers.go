package Schema

import (
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var createYearConfigResolver = func(params graphql.ResolveParams) (interface{}, error) {
	{
		_, err := db.Exec(`INSERT INTO classmanager.year (year, year_group) VALUES ($1, $2)`,
			params.Args["year"].(int), params.Args["yearGroup"].(int))
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	var yearId int
	{
		err := db.QueryRow(`SELECT id FROM classmanager.year WHERE year=$1 AND year_group=$2`,
			params.Args["year"].(int), params.Args["yearGroup"].(int)).Scan(&yearId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	{
		terms := params.Args["terms"].([]Models.Term)
		for _, term := range terms {
			_, err := db.Exec(
				`INSERT INTO classmanager.terms (year_id, name, start_time, end_time) 
                VALUES ($1, $2, $3, $4);`,
				    yearId, term.Name, term.StartTime, term.EndTime)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}
	{
		publicHolidays := params.Args["publicHolidays"].([]Models.PublicHoliday)
		for _, publicHoliday := range publicHolidays {
			_, err := db.Exec(
				`INSERT INTO classmanager.public_holidays (year_id, name, start_time, end_time)
                VALUES ($1, $2, $3, $4);`,
                    yearId, publicHoliday.Name, publicHoliday.StartTime, publicHoliday.EndTime)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}
	{
		events := params.Args["events"].([]Models.Event)
		for _, event := range events {
			for _, class := range event.Classes {
				_, err := db.Exec(
					`INSERT INTO classmanager.events (year_id, class_id, name, start_time, end_time)
                    VALUES ($1, $2, $3, $4, $5);`,
					    yearId, class, event.Name, event.StartTime, event.EndTime)
				if err != nil {
					log.Println(err)
					return nil, err
				}
			}
		}
	}
	{
		periods := params.Args["periods"].([]Models.Period)
		for _, period := range periods {
			_, err := db.Exec(
				`INSERT INTO classmanager.periods (year_id, name, start_time, end_time)
                VALUES ($1, $2, $3, $4);`,
                    yearId, period.Name, period.StartTime, period.EndTime)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}

	savedVars := Models.YearConfig{
		ID: yearId,
		Year: params.Args["year"].(int),
		YearGroup: params.Args["yearGroup"].(int),
		Terms: params.Args["terms"].([]Models.Term),
		PublicHolidays: params.Args["publicHolidays"].([]Models.PublicHoliday),
		Events: params.Args["events"].([]Models.Event),
		Periods: params.Args["periods"].([]Models.Period),
	}
	return savedVars, nil
}

