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
	username := params.Args["username"].(string)
	password := params.Args["password"].(string)

	// create variables to save user data from database
	var user Models.User
	var passwordHash string

	err := db.QueryRow(
		`SELECT id, role, username, email, password
        FROM users 
        WHERE username = $1;`,
		username).Scan(&user.ID, &user.Role, &user.Username, &user.Email, &passwordHash)
	if err != nil {
		log.Println(err)
		return nil, errors.New("username not found")
	}

	// if the verify password is successful create token
	if Auth.VerifyPassword(password, passwordHash) {
		// store variables in token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"id": user.ID,
			"role": user.Role,
			"username": username,
			"email": user.Email,
		})

		// sign token with secret key
		tokenString, err := token.SignedString(Env.GetSecretKey())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		loginVar := Models.Login{
			Token: tokenString,
			User: user,
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
	role := params.Args["role"].(Models.Role)

	// return token verification boolean
	return Auth.VerifyToken(token, role), nil
}
