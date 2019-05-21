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
	Port      string         `yaml:"port"`
	SecretKey string         `yaml:"secret_key"`
	Database  DatabaseConfig `yaml:"database"`
}

// struct which holds database config information
type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
}

var fpath = "config.yml"

var config Config

func init() {
	if err := decodeYaml(); err != nil {
		log.Fatal(err)
	}
	if err := encodeYaml(); err != nil {
		log.Fatal(err)
	}
}

func GetSecretKey() []byte {
	secretKeyStr, err := base64.StdEncoding.DecodeString(config.SecretKey)
	if err != nil {
		log.Println(err)
	}
	return secretKeyStr
}

func GetPort() string {
	return config.Port
}

func GetDatabaseConfig() DatabaseConfig {
	return config.Database
}

func decodeYaml() error {
	fmt.Println("parsing yaml config file")
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		return err
	}

	if config.SecretKey == "" {
		config.SecretKey = generateSecretKey()
	}
	return nil
}

func encodeYaml() error {
	fmt.Println("encoding yaml config file")
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fpath, data, 0644); err != nil {
		return err
	}
	return nil
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