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

var selectClassTaskWithStudentId = `
SELECT task_marks.id, task_marks.task_mark, task_marks.feedback, task_marks.time_stamp, 
       tasks.id, tasks.task_name, tasks.task_description, tasks.total_mark, tasks.due_time
FROM task_marks
INNER JOIN tasks
ON task_marks.task_id=tasks.id
WHERE student_class_id=$1, tasks.id=$2;
`

var createTaskMarkResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// authenticate user
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	// variables to hold data-
	var (
		taskMark Models.TaskMark
		task Models.Task
		studentId int
	)

	// saving arguments to variables
	task.ID = params.Args["taskId"].(int)
	studentId = params.Args["studentClassId"].(int)
	taskMark.TaskMark = params.Args["taskMark"].(int)
	taskMark.Feedback = params.Args["feedback"].(string)
	taskMark.TimeStamp = time.Now()

	// verify the task mark is less than the total mark
	if taskMark.TaskMark > task.TotalMarks {
		return nil, errors.New("error task mark cannot be greater than total marks")
	}

	{
		// insert data into database
		_, err := db.Exec(insertTaskMark, task.ID, studentId, taskMark.TaskMark, taskMark.Feedback, taskMark.TimeStamp)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	{
		// get created table id
		if err := db.QueryRow(selectTaskMarkId, studentId, task.ID).Scan(&taskMark.ID); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	{
		// get the class task from the
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

var readClassTaskMarkResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// Authenticate user
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var taskMark Models.TaskMark

	err := db.QueryRow(selectClassTaskWithStudentId,
		params.Args["studentId"].(int), params.Args["taskId"].(int)).Scan(
			&taskMark.ID, &taskMark.TaskMark, &taskMark.Feedback, &taskMark.TimeStamp, &taskMark.AssignedTask.ID,
			&taskMark.AssignedTask.Name, &taskMark.AssignedTask.Description, &taskMark.AssignedTask.TotalMarks,
			&taskMark.AssignedTask.DueTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return taskMark, nil
}
