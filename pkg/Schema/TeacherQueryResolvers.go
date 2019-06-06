package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var selectClassWithTeacherId = `
SELECT classes.id, classes.class_id, users.role, users.username, users.email
FROM classes
INNER JOIN users
ON classes.teacher_id=users.id
WHERE classes.teacher_id=$1;
`

var selectStudentFromClassStudent = `
SELECT users.id, users.role, users.username, users.email
FROM users
INNER JOIN class_student
ON class_student.student_id=users.id
WHERE class_student.class_id=$1;
`

/*
var selectClassQuery = `
SELECT classes.class_id, users.id, users.role, users.username, users.email
FROM classes
INNER JOIN users
ON classes.teacher_id=users.id
WHERE classes.id=$1;
`

var selectStudentClass = `
SELECT class_student.id, users.id, users.role, users.username, users.email
FROM class_student
INNER JOIN users
ON class_student.student_id=users.id
WHERE class_student.class_id=$1;
`

 */

var listTeachersClassesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var classes []Models.Class

	rows, err := db.Query(selectClassWithTeacherId, params.Args["teacherId"].(int))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var class Models.Class

		class.Teacher.ID = params.Args["teacherId"].(int64)
		err := rows.Scan(&class.ID, &class.ClassId, &class.Teacher.Role, &class.Teacher.Username, &class.Teacher.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		studentRows, err := db.Query(selectStudentFromClassStudent, class.ID)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for studentRows.Next() {
			var student Models.User

			err := rows.Scan(&student.ID, &student.Role, &student.Username, &student.Email)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			class.Students = append(class.Students, student)
		}
		classes = append(classes, class)
	}
	return classes, nil
}

/*
var viewClassResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var class Models.Class
	class.ID = params.Args["id"].(int)

	err := db.QueryRow(selectClassQuery, class.ID).Scan(
		&class.ClassID, &class.Teacher.ID, &class.Teacher.Role, &class.Teacher.Username, &class.Teacher.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	studentRows, err := db.Query(selectStudentClass, class.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var students []Models.ClassStudent

	for studentRows.Next() {
		var student Models.ClassStudent
		err := studentRows.Scan(
			&student.ID, &student.Student.ID, &student.Student.Role, &student.Student.Username, &student.Student.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}


	}
}

 */