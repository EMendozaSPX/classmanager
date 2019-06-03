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
	rows, err := db.Query(`SELECT id, role, username, email FROM classmanager.users;`)
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

	err := db.QueryRow(`SELECT id, role, username, email FROM classmanager.users WHERE id=$1;`,
		params.Args["id"].(int)).Scan(&user.ID, &user.Role, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
	}

	return user, nil
}


var listClassesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, permissionDenied
	}

	var classes []Models.Class

	rows, err := db.Query(`SELECT * FROM classmanager.classes;`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var (
			class Models.Class
			teacherID int
		)
		if err := rows.Scan(&class.ID, &class.ClassID, &teacherID); err != nil {
			log.Println(err)
		}

		err := db.QueryRow(`SELECT id, role, username, email FROM classmanager.users WHERE id=$1;`,
			teacherID).Scan(&class.Teacher.ID, &class.Teacher.Role, &class.Teacher.Username, &class.Teacher.Email)
		if err != nil {
			log.Println(err)
		}

		studentRows, err := db.Query(`SELECT student_id FROM classmanager.class_student WHERE class_id=$1;`,
			class.ID)
		if err != nil {
			log.Println(err)
		}

		for studentRows.Next() {
			var (
				studentID int
				student Models.User
				)
			if err := studentRows.Scan(&studentID); err != nil {
				log.Println(err)
			}
			err := db.QueryRow(`SELECT id, role, username, email FROM classmanager.users WHERE id=$1;`,
				studentID).Scan(&student.ID, &student.Role, &student.Username, &student.Email)
			if err != nil {
				log.Println(err)
			}

			class.Students = append(class.Students, student)
		}

		classes = append(classes, class)
	}
	return classes, nil
}