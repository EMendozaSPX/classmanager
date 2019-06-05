package Schema

import (
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"time"
)

var createBehaviourNoteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Teacher) {
		return nil, permissionDenied
	}
	classStudentId := params.Args["classStudentId"].(string)
	name := params.Args["name"].(string)
	note := params.Args["note"].(string)
	time := time.Now()

	
}