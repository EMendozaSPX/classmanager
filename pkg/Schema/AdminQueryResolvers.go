package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

// resolves the listUser query
var listUsersResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// verify admin authorization
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, permissionDenied
	}

	// create array of users to store data
	var users []Models.User

	// get data from database
	rows, err := db.Query(`SELECT id, role, username, email FROM users;`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// save each user in row to array of users
	for rows.Next() {
		var user Models.User
		if err := rows.Scan(&user.ID, &user.Role, &user.Username, &user.Email); err != nil {
			log.Println(err)
		}

		// add user to list of users
		users = append(users, user)
	}
	return users, nil
}

var getUsersResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, permissionDenied
	}

	var user Models.User

	err := db.QueryRow(`SELECT id, role, username, email FROM users WHERE id=$1;`,
		params.Args["id"].(int)).Scan(&user.ID, &user.Role, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
	}

	return user, nil
}