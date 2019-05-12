package Auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/emendoza/classmanager/pkg/Env"
	"github.com/emendoza/classmanager/pkg/Models"
)

// verify if token is valid for given user role
func VerifyToken(tokenString string, role Models.Role) bool {
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
	if claims["role"] == role {
		return true
	}
	fmt.Println("user not validated")
	return false
}