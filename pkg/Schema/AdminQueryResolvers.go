package Schema

import (
	"database/sql"
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

// resolves the listUser query
var listUsersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	// verify token
	conditionOne := Auth.VerifyToken(p.Context.Value("token").(string), "admin")
	conditionTwo := Auth.VerifyToken(p.Context.Value("token").(string), "teacher")
	conditionThree := p.Args["usertype"].(int) == 3
	if !conditionOne && (!conditionTwo || !conditionThree ) {
		return nil, errors.New("user not verified")
	}

	// create array of users to store data
	var users []Models.User

	// get data from database
	var (
		rows *sql.Rows
		err error
	)
	switch p.Args["usertype"].(Models.Usertype) {
	case Models.Admin:
		rows, err = db.Query("SELECT id, username, email FROM classmanager.admins")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	case Models.Teacher:
		rows, err = db.Query("SELECT id, username, email FROM classmanager.teachers")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	case Models.Student:
		rows, err = db.Query("SELECT id, username, email FROM classmanager.students")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	default:
		return nil, errors.New("usertype not found error")
	}

	// save each user in row to array of users
	for rows.Next() {
		var (
			id       int64
			username string
			email    string
		)
		if err := rows.Scan(&id, &username, &email); err != nil {
			log.Println(err)
		}
		users = append(users, Models.User{
			ID:       id,
			Username: username,
			Email:    email,
		})
	}
	return users, nil
}