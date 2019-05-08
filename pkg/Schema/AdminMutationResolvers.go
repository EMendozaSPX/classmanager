package Schema

/*
var createYearConfigResolver = func(params graphql.ResolveParams) (interface{}, error) {
	{
		_, err := db.Exec("INSERT INTO classmanager.year (year, year_group) VALUES ($1, $2)",
			params.Args["year"].(int), params.Args["yearGroup"].(int))
		if err != nil {
			log.Println(err)
		}
	}
	{
		terms := params.Args["terms"].([]Models.Term)
		for _, term := range terms {
			_, err := db.Exec("INSERT INTO classmanager.terms (term, start_time, end_time) VALUES ($1, $2, $3)",
				term.Term, term.StartTime, term.EndTime)
			if err != nil {
				log.Println(err)
			}
		}
	}
	{

	}
}
 */

