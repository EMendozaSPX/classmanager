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
var loginResolver = func(p graphql.ResolveParams) (interface{}, error) {
	// save username and password parameters to local variables
	usernameInput := p.Args["username"].(string)
	passwordInput := p.Args["password"].(string)

	// create variables to save user data from database
	var username, email, password string

	// get username, email and password from database
	// if the query failed send the error username not found to client
	switch p.Args["usertype"].(Models.Usertype) {
	case Models.Admin:
		err := db.QueryRow(`SELECT username, email, password FROM classmanager.admins WHERE username = $1`,
			usernameInput).Scan(&username, &email, &password)
		if err != nil {
			log.Println(err)
			return nil, errors.New("username of usertype admin not found")
		}
	case Models.Teacher:
		err := db.QueryRow(`SELECT username, email, password FROM classmanager.teachers WHERE username = $1`,
			usernameInput).Scan(&username, &email, &password)
		if err != nil {
			log.Println(err)
			return nil, errors.New("username of usertype teacher not found")
		}
	case Models.Student:
		err := db.QueryRow(`SELECT username, email, password FROM classmanager.students WHERE username = $1`,
			usernameInput).Scan(&username, &email, &password)
		if err != nil {
			log.Println(err)
			return nil, errors.New("username of usertype student not found")
		}
	default:
		return nil, errors.New("usertype not found")
	}

	// if the verify password is successful create token
	if Auth.VerifyPassword(passwordInput, password) {
		// store variables in token
		var token *jwt.Token
		switch p.Args["usertype"].(Models.Usertype) {
		case Models.Admin:
			token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"usertype": "admin",
				"username": username,
				"email":    email,
			})
		case Models.Teacher:
			token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"usertype": "teacher",
				"username": username,
				"email":    email,
			})
		case Models.Student:
			token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"usertype": "student",
				"username": username,
				"email":    email,
			})
		default:
			return nil, errors.New("failed to sign token")
		}

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

// verifies users authorization to access certain pages on the website
var verifyAuthorizationResolver = func(p graphql.ResolveParams) (interface{}, error) {
	token := p.Context.Value("token").(string)

	// check which usertype needs to be authorized
	switch p.Args["usertype"].(Models.Usertype) {
	case Models.Admin:
		return Auth.VerifyToken(token, "admin"), nil
	case Models.Teacher:
		return Auth.VerifyToken(token, "teacher"), nil
	case Models.Student:
		return Auth.VerifyToken(token, "student"), nil
	default:
		return nil, errors.New("user does not have required permissions")
	}
}
