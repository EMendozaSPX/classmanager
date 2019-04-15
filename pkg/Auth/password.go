package Auth

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// generates a secret key used for authentication
func GenerateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(key)
}

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