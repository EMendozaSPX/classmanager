package Data

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Env"
	"github.com/graphql-go/graphql"
	"log"
)

// returns the jwt authorization token for admin users
var authenticateAdmin = func(p graphql.ResolveParams) (interface{}, error) {
	// save username and password parameters to local variables
	usernameInput := p.Args["username"].(string)
	passwordInput := p.Args["password"].(string)

	// create variables to save user data from database
	var username, email, password string

	// get username, email and password from database
	err := db.QueryRow(`SELECT username, email, password FROM classmanager.admins WHERE username = $1`,
		usernameInput).Scan(&username, &email, &password)

	// if the query failed send the error username not found to client
	if err != nil {
		log.Println(err)
		return nil, errors.New("username not found")
	}

	// if the verify password is successful create token
	if Auth.VerifyPassword(passwordInput, password) {
		// store variables in token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"user_type": "Admin",
			"username": username,
			"email": email,
		})

		// sign token with secret key
		tokenString, err := token.SignedString(Env.GetSecretKey())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// return token
		return tokenString, nil
	}

	// return authentication failed if password verification failed
	return nil, errors.New("authentication failed")
}