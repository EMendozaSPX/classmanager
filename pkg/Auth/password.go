package Auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(password string) string {
	// Cast password to a byte array
	passwordByteArray := []byte(password)

	// hash password and log any errors
	hash, err := bcrypt.GenerateFromPassword(passwordByteArray, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	// return hash casted to a string
	return string(hash)
}

func VerifyPassword(password string, hash string) bool {
	byteHash := []byte(hash)
	bytePass := []byte(password)
	if err := bcrypt.CompareHashAndPassword(byteHash, bytePass); err != nil {
		log.Println(err)
		return false
	}
	return true
}