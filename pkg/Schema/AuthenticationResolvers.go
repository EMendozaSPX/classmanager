package Schema

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/emendoza/classmanager/pkg/Auth"
	"github.com/emendoza/classmanager/pkg/Env"
	"github.com/emendoza/classmanager/pkg/Models"
	"github.com/graphql-go/graphql"
	"log"
)

// returns the jwt authorization token
var loginResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// save parameters to variables for convenience
	usernameInput := params.Args["username"].(string)
	passwordInput := params.Args["password"].(string)

	// create variables to save user data from database
	var role Models.Role
	var username, email, password string

	err := db.QueryRow(
		`SELECT role, username, email, password
        FROM classmanager.users 
        WHERE username = $1;`,
		usernameInput).Scan(&role, &username, &email, &password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("username not found")
	}

	// if the verify password is successful create token
	if Auth.VerifyPassword(passwordInput, password) {
		// store variables in token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"role": role,
			"username": username,
			"email": email,
		})

		// sign token with secret key
		tokenString, err := token.SignedString(Env.GetSecretKey())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		loginVar := Models.Login{
			Token: tokenString,
			Role: role,
		}

		// return token
		return loginVar, nil
	}

	// return authentication failed if password verification failed
	return nil, errors.New("authentication failed")
}

// verifies users authorization to access certain pages on the website
var verifyAuthorizationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// save parameters as variables for convenience
	token := params.Context.Value("token").(string)
	role := params.Context.Value("role").(Models.Role)

	// return token verification boolean
	return Auth.VerifyToken(token, role), nil
}
