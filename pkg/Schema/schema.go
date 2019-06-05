package Schema

import (
	"database/sql"
	"errors"
	"github.com/emendoza/classmanager/pkg/Database"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var (
	Schema graphql.Schema
	db     *sql.DB
)
var permissionDenied = errors.New("permission denied")

func init() {
	// create a error variable to handle errors
	var err error

	// get db instance from file
	db = Database.GetDB()

	// create a root query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"verifyAuthorization": &graphql.Field {
				Type: graphql.Boolean,
				Description: "Provides a user access to a particular site",
				Args: graphql.FieldConfigArgument{
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
					},
				},
				Resolve: verifyAuthorizationResolver,
			},
			"listUsers": &graphql.Field{
				Type: graphql.NewList(Models.UserType),
				Description: "Get a list users of a certain usertype",
				Resolve: listUsersResolver,
			},
			"readUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Get a user from database",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: getUsersResolver,
			},
			"listClasses": &graphql.Field{
				Type: graphql.NewList(Models.ClassType),
				Description: "Get a list of classes",
				Resolve: listClassesResolver,
			},
			"viewTimetable": &graphql.Field{
				Type: graphql.NewList(Models.TimetableType),
				Description: "Get a user from database",
				Args: graphql.FieldConfigArgument{
					"teacherId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: viewTimetableResolver,
			},
			"listClassesByTeacher": &graphql.Field{
				Type: graphql.NewList(Models.ClassType),
				Description: "Get a list of a teachers classes",
				Args: graphql.FieldConfigArgument{
					"teacherId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: listClassesByTeacher,
			},
			"readBehaviourNote": &graphql.Field{
				Type: Models.BehaviourNoteType,
				Description: "Read a behaviour note",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: readBehaviourNoteResolver,
			},
			"readClassTask": &graphql.Field{
				Type: Models.TaskType,
				Description: "Read a class task",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: readClassTaskResolver,
			},
		},
	})


	// Creates a graphql mutation type
	// graphql mutation types take arguments and return corresponding information
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: Models.LoginType,
				Description: "Sign in users through json web tokens.",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: loginResolver,
			},
			"createUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Create a new user.",
				Args: graphql.FieldConfigArgument{
					"role": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(Models.RoleEnum),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createUserResolver,
			},
			"updateUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Update a users information.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"role": &graphql.ArgumentConfig{
						Type: Models.RoleEnum,
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: updateUserResolver,
			},
			"deleteUser": &graphql.Field{
				Type: Models.UserType,
				Description: "Remove a user from database.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: deleteUserResolver,
			},
			"createClass": &graphql.Field{
				Type: Models.ClassType,
				Description: "Add a class to database",
				Args: graphql.FieldConfigArgument{
					"classId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"teacherId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: createClassResolver,
			},
			"createTimetableEntry": &graphql.Field{
				Type: Models.PeriodType,
				Description: "Create timetable entry",
				Args: graphql.FieldConfigArgument{
					"classId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"periodName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"weekday": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: createTimetableEntryResolver,
			},
			"createBehaviourNote": &graphql.Field{
				Type: Models.BehaviourNoteType,
				Description: "create a behaviour note for a student",
				Args: graphql.FieldConfigArgument{
					"classStudentId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"note": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createBehaviourNoteResolver,
			},
			"updateBehaviourNote": &graphql.Field{
				Type: Models.BehaviourNoteType,
				Description: "Update a behaviour note",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"note": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: updateBehaviourNoteResolver,
			},
			"deleteBehaviourNote": &graphql.Field{
				Type: Models.BehaviourNoteType,
				Description: "Delete a behaviour note",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: deleteBehaviourNoteResolver,
			},
			"createClassTask": &graphql.Field{
				Type: Models.TaskType,
				Description: "Create a class task",
				Args: graphql.FieldConfigArgument{
					"classId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"totalMarks": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"dueTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
				},
				Resolve: createClassTaskResolver,
			},
			"updateClassTask": &graphql.Field{
				Type: Models.TaskType,
				Description: "Update a class task",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"totalMarks": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"dueTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
				},
				Resolve: updateClassTaskResolver,
			},
			"deleteClassTask": &graphql.Field{
				Type: Models.TaskType,
				Description: "Delete class task",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: deleteClassTaskResolver,
			},
		},
	})

	// create a new graphql schema using the query and mutation types, if failed returns an error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
			Mutation: mutationType,
		})

	// print error to console
	if err != nil {
		log.Println(err)
	}
}
