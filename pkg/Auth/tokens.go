package Auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/emendoza/classmanager/pkg/Env"
)

// verify if token is valid for given usertype
func VerifyToken(tokenString string, usertype string) bool {
	// Verify if token exists
	if tokenString == "" {
		fmt.Println("Token not found")
		return false
	}

	// parse json web token and validate using secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return Env.GetSecretKey(), nil
	})

	// check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("could not verify token")
		return false
	}

	// map json web token claims to claims variable
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("could not map claims")
		return false
	}

	// verify user
	if claims["usertype"] == usertype {
		return true
	}
	fmt.Println("user not validated")
	return false
}