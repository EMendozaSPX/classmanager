package Env

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/emendoza/classmanager/pkg/Auth"
	"io/ioutil"
	"log"
	"os"
)

// a struct type that holds all information not stored in a database
type EnvVariables struct {
	SECRET_KEY string `json:"secret_key"`
}

// this function gets the secret key used to sign the json web tokens from the env file
func GetSecretKey() []byte {
	// get env variables from env.json
	envVariables := importVariables()

	// decode string into a byte array
	secretKey, err := b64.StdEncoding.DecodeString(envVariables.SECRET_KEY)
	if err != nil {
		log.Println(err)
		return nil
	}
	return secretKey
}

// this function generates a env.json file at the root of the executable if it is not found
func createFile() {
	fmt.Println("Creating file env.json ...")

	// create env.json file
	f, err := os.Create("env.json")
	if err != nil {
		log.Println(err)
	}

	// close the file env.json
	if err := f.Close(); err != nil {
		log.Println(err)
	}
}

// this function writes variables into the env.json file
func writeFile(variables EnvVariables) {
	fmt.Println("Writing to file env.json ...")

	// serialize the the EnvVariables struct into indented json with a double space
	result, err := json.MarshalIndent(variables, "", "  ")
	if err != nil {
		log.Println(err)
	}

	// write json code to the file env.json
	if err := ioutil.WriteFile("env.json", result, 0644); err != nil {
		log.Println(err)
	}
}

// this function creates the variables not stored in a database
func createVariables() EnvVariables {
	fmt.Println("Generating secret key ...")
	secretKey := Auth.GenerateSecretKey()
	return EnvVariables{
		SECRET_KEY: secretKey,
	}
}

// this function imports and parses variables from the env.json file
func importVariables() EnvVariables {
	fmt.Println("Reading env.json file ...")

	// read env.json file
	envJson, err := ioutil.ReadFile("env.json")

	// if file does not exist create a file and add constants to it
	if os.IsNotExist(err) {
		createFile()
		envVariables := createVariables()
		writeFile(envVariables)
		return envVariables
	}
	if err != nil {
		log.Println(err)
	}

	// create a variable of type EnvVariables
	envVariables := EnvVariables{}

	// parse json from env.json into the envVariables variable
	if err := json.Unmarshal(envJson, &envVariables); err != nil {
		log.Println(err)
	}
	return envVariables
}