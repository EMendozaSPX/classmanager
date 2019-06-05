package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

var insertClassTask = `
INSERT INTO tasks (class_id, task_name, task_description, total_mark, due_time)
VALUES ($1, $2, $3, $4, $5);
`

var selectIdFromClassTask = `
SELECT id 
FROM tasks 
WHERE task_name=$1;
`

var selectClassTaskQuery = `
SELECT task_name, task_description, total_mark, due_time
FROM tasks
WHERE id=$1;
`

var updateClassTask = `
UPDATE tasks
SET task_name=$2, task_description=$3, total_mark=$4, due_time=$5
WHERE id=$1;
`

var createClassTaskResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var newTask Models.Task

	classId := params.Args["classId"].(int)

	newTask.Name = params.Args["name"].(string)
	newTask.Description = params.Args["description"].(string)
	newTask.TotalMarks = params.Args["totalMarks"].(int)
	newTask.DueTime = params.Args["dueTime"].(time.Time)

	{
		_, err := db.Exec(insertClassTask,
			classId, newTask.Name, newTask.Description, newTask.TotalMarks, newTask.DueTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	if err := db.QueryRow(selectIdFromClassTask, newTask.Name).Scan(&newTask.ID); err != nil {
		log.Println(err)
		return nil, err
	}

	return newTask, nil
}

var readClassTaskResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var task Models.Task
	task.ID = params.Args["id"].(int)

	err := db.QueryRow(selectClassTaskQuery,
		task.ID).Scan(&task.Name, &task.Description, &task.TotalMarks, &task.DueTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return task, nil
}

var updateClassTaskResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	var task Models.Task
	task.ID = params.Args["id"].(int)
	task.Name = params.Args["name"].(string)
	task.Description = params.Args["description"].(string)
	task.TotalMarks = params.Args["totalMarks"].(int)
	task.DueTime = params.Args["dueTime"].(time.Time)

	_, err := db.Exec(updateClassTask, task.ID, task.Name, task.Description, task.TotalMarks, task.DueTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return task, nil
}

var deleteClassTaskResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}

	taskId := params.Args["id"].(int)
	_, err := db.Exec(`DELETE FROM tasks WHERE id=$1`, taskId)
	if err != nil {
		log.Println(err)
	}
	return nil, nil
}
