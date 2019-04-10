package Data

import (
	"errors"
	"github.com/graphql-go/graphql"
)


var authenticateAdmin = func(p graphql.ResolveParams) (interface{}, error) {

	return nil, errors.New("authentication failed")
}