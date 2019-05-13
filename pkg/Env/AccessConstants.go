package Env

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// a struct type that holds all information not stored in a database
type Config struct {
	SecretKey string         `yaml:"secret_key"`
	Database  DatabaseConfig `yaml:"database"`
}

// struct which holds database config information
type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
}

var fpath = "config.yml"

var config Config

func GetSecretKey() []byte {
	decodeYaml()
	encodeYaml()
	secretKeyStr, err := base64.StdEncoding.DecodeString(config.SecretKey)
	if err != nil {
		log.Println(err)
	}
	return secretKeyStr
}

func GetDatabaseConfig() DatabaseConfig {
	decodeYaml()
	encodeYaml()
	return config.Database
}

func decodeYaml() {
	fmt.Println("parsing yaml config file")
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}

	if config.SecretKey == "" {
		config.SecretKey = generateSecretKey()
	}
}

func encodeYaml() {
	fmt.Println("encoding toml config file")
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fpath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

// generates a secret key to sign auth tokens
func generateSecretKey() string {
	fmt.Println("Generating secret key")
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	keyString := base64.StdEncoding.EncodeToString(key)
	return keyString
}