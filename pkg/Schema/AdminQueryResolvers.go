package Schema

import (
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

// resolves the listUser query
var listUsersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	// if either a teacher or a admin
	conditionOne := Auth.VerifyToken(p.Context.Value("token").(string), "admin")
	conditionTwo := Auth.VerifyToken(p.Context.Value("token").(string), "teacher")
	conditionThree := p.Args["role"].(Models.Role) == Models.Student
	if !conditionOne && (!conditionTwo || !conditionThree ) {
		return nil, errors.New("user not verified")
	}

	// create array of users to store data
	var users []Models.User

	// get data from database
	rows, err := db.Query(`SELECT (id, role, username, email) FROM classmanager.users`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// save each user in row to array of users
	for rows.Next() {
		var (
			id       int64
			role     Models.Role
			username string
			email    string
		)
		if err := rows.Scan(&id, role, &username, &email); err != nil {
			log.Println(err)
		}

		// add user to list of users
		users = append(users, Models.User{
			ID:       id,
			Role:     role,
			Username: username,
			Email:    email,
		})
	}
	return users, nil
}