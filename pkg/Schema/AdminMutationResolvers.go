package Schema

import (
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var createUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Args["token"].(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, errors.New("permission denied")
	}
	usernameInput := params.Args["username"].(string)
	{
		role := params.Args["role"].(string)
		email := params.Args["email"].(string)
		password := Auth.HashAndSalt(params.Args["password"].(string))
		_, err := db.Exec(
			`INSERT INTO classmanager.users (role, username, email, password) VALUES ($1, $2, $3, $4)`,
			role, usernameInput, email, password)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	var user Models.User
	err := db.QueryRow(`SELECT (id, role, username, email) FROM classmanager.users WHERE username=$1`,
		usernameInput).Scan(&user.ID, &user.Role, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
	}

	return user, nil
}

var updateUserResolver = func(params graphql.ResolveParams) (interface{}, error) {

}

var deleteUserResolver = func(params graphql.ResolveParams) (interface{}, error) {

}
