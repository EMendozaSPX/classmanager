package Env

import (
	"crypto/rand"
	"fmt"
	"log"
)

// a struct type that holds all information not stored in a database
type EnvVariables struct {
	SECRET_KEY   []byte       `json:"secret_key"`
	DBUser       *DatabaseUser `json:"database_user"`
}

// Struct which holds user information
type DatabaseUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var variables = &EnvVariables{
	SECRET_KEY: generateSecretKey(),
	DBUser: dbUser,
}

// this function gets the secret key used to sign the json web tokens from the env file
func GetSecretKey() []byte {
	return variables.SECRET_KEY
}

// get database user information from hard coded variable
func GetDatabaseUser() *DatabaseUser {
	return variables.DBUser
}

// generates a secret key to sign auth tokens
func generateSecretKey() []byte {
	fmt.Println("Generating secret key")
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return key
}