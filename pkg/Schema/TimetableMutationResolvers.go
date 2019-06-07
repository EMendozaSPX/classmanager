package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var createTimetableInsert = `
INSERT INTO timetable (
    class_id,
    period_name,
    week_day
)
VALUES (
    $1,
    $2,
    $3
);
`
var createTimetableEntryResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, permissionDenied
	}

	_, err := db.Exec(createTimetableInsert,
		params.Args["classId"].(int), params.Args["periodName"].(string), params.Args["weekday"].(int))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}
