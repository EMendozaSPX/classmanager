package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

// db queries
var insertBehaviourNote = `
INSERT INTO behaviour_notes (
    class_student_id,
    name,
    note,
    time_stamp
)
VALUES (
    $1,
    $2,
    $3,
    $4
);
`

var selectBehaviourNote = `
SELECT name, note, time_stamp
FROM behaviour_notes
WHERE id=$1;
`

var updateBehaviourNote = `
UPDATE behaviour_notes
SET name=$1, note=$2, time_stamp=$3
WHERE id=$4;
`

var selectBehaviourNotesQuery = `
SELECT id, name, note, time_stamp
FROM behaviour_notes
WHERE class_student_id=$1;
`


var listBehaviourNotesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// authorization function
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	// create a behaviour note list for later usage
	var behaviourNotes []Models.BehaviourNote

	// storing classStudentId in a variable
	classStudentId := params.Args["classStudentId"].(int)

	// query database for behaviour notes of a student
	rows, err := db.Query(selectBehaviourNotesQuery, classStudentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		// create a variable to store a behaviour note
		var behaviourNote Models.BehaviourNote
		err := rows.Scan(&behaviourNote.ID, &behaviourNote.Name, &behaviourNote.Note, &behaviourNote.TimeStamp)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// append behaviour note to behaviour notes list
		behaviourNotes = append(behaviourNotes, behaviourNote)
	}

	return behaviourNotes, nil
}

var createBehaviourNoteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var behaviourNote Models.BehaviourNote

	classStudentId := params.Args["classStudentId"].(int)
	behaviourNote.Name = params.Args["name"].(string)
	behaviourNote.Note = params.Args["note"].(string)
	behaviourNote.TimeStamp = time.Now()

	_, err := db.Exec(insertBehaviourNote,
		classStudentId, behaviourNote.Name, behaviourNote.Note, behaviourNote.TimeStamp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := db.QueryRow(`SELECT id FROM behaviour_notes WHERE name=$1`,
		behaviourNote.Name).Scan(&behaviourNote.ID); err != nil {
		log.Println(err)
		return nil, err
	}

	return behaviourNote, nil

}

var readBehaviourNoteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var behaviourNote Models.BehaviourNote
	behaviourNote.ID = params.Args["id"].(int)

	err := db.QueryRow(selectBehaviourNote,
		behaviourNote.ID).Scan(&behaviourNote.Name, &behaviourNote.Note, &behaviourNote.TimeStamp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return behaviourNote, err
}

var updateBehaviourNoteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var behaviourNote Models.BehaviourNote
	behaviourNote.ID = params.Args["id"].(int)
	behaviourNote.Name = params.Args["name"].(string)
	behaviourNote.Note = params.Args["note"].(string)
	behaviourNote.TimeStamp = time.Now()
	_, err := db.Exec(updateBehaviourNote,
		behaviourNote.Name, behaviourNote.Note, behaviourNote.TimeStamp, behaviourNote.ID)

	if err != nil {
		return nil, err
	}
	return behaviourNote, nil
}

var deleteBehaviourNoteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var id = params.Args["id"].(int)
	_, err := db.Exec(`DELETE FROM behaviour_notes WHERE id=$1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}