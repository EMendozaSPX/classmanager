package Schema

import (
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

var insertTaskMark = `
INSERT INTO task_marks (task_id, student_class_id, task_mark, feedback, time_stamp)
VALUES ($1, $2, $3, $4, $5);
`

var selectTaskMarkId = `
SELECT id
FROM task_marks
WHERE student_class_id=$1 AND task_id=$2
`

var createTaskMarkResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var (
		taskMark Models.TaskMark
		task Models.Task
		studentId int
	)
	task.ID = params.Args["taskId"].(int)
	studentId = params.Args["studentClassId"].(int)
	taskMark.TaskMark = params.Args["taskMark"].(int)
	taskMark.Feedback = params.Args["feedback"].(string)
	taskMark.TimeStamp = time.Now()

	if taskMark.TaskMark > task.TotalMarks {
		return nil, errors.New("error task mark cannot be greater than total marks")
	}

	{
		_, err := db.Exec(insertTaskMark, task.ID, studentId, taskMark.TaskMark, taskMark.Feedback, taskMark.TimeStamp)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	{
		if err := db.QueryRow(selectTaskMarkId, studentId, task.ID).Scan(&taskMark.ID); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	{
		err := db.QueryRow(selectClassTaskQuery,
			task.ID).Scan(&task.Name, &task.Description, &task.TotalMarks, &task.DueTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	taskMark.AssignedTask = task
	return taskMark, nil
}


