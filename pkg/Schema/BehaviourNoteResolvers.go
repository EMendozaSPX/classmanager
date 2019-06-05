package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

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

	if err := db.QueryRow(`SELECT id FROM behaviour_notes WHERE name=$1`, behaviourNote.Name).Scan(&behaviourNote.ID); err != nil {
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