package Schema

import (
	"errors"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

var createUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// verify users authorization
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, errors.New("permission denied")
	}

	// save username input to variable for convenient access
	usernameInput := params.Args["username"].(string)

	// block scoping sql insert statement so variables will go out of scope within the function
	{
		// save variables for convenient access
		role := params.Args["role"].(Models.Role)
		email := params.Args["email"].(string)
		password := Auth.HashAndSalt(params.Args["password"].(string))

		// insert data into database
		_, err := db.Exec(
			`INSERT INTO classmanager.users (role, username, email, password) VALUES ($1, $2, $3, $4)`,
			role, usernameInput, email, password)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	// return the created user
	var user Models.User
	err := db.QueryRow(`SELECT id, role, username, email FROM classmanager.users WHERE username=$1`,
		usernameInput).Scan(&user.ID, &user.Role, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
	}

	return user, nil
}

var updateUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)
	if !Auth.VerifyToken(token, Models.Admin) {
		return nil, errors.New("permission denied")
	}

	id := params.Args["id"].(int)
	query := `UPDATE classmanager.users SET $1=$2 WHERE id=$3`

	if role := params.Args["role"].(Models.Role); role != "" {
		_, err := db.Exec(query, "role", role, id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if username := params.Args["username"].(string); username != "" {
		_, err := db.Exec(query, "username", username, id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if email := params.Args["email"].(string); email != "" {
		_, err := db.Exec(query, "email", email, id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if passwordHash := Auth.HashAndSalt(params.Args["password"].(string)); passwordHash != "" {
		_, err := db.Exec(query, "password", passwordHash, id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	var user Models.User
	err := db.QueryRow(`SELECT id, role, username, email FROM classmanager.users WHERE id=$1`,
		id).Scan(&user.ID, &user.Role, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

var deleteUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("token").(string)

	if Auth.VerifyToken(token, Models.Admin) {
		return nil, errors.New("permission denied")
	}

	id := params.Args["id"].(int)
	_, err := db.Exec(`DELETE FROM classmanager.users WHERE id=$1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}
