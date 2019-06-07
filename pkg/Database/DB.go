package Database

import (
	"database/sql"
	"fmt"
	"github.com/emendoza/classmanager/pkg/Env"
	"log"

	_ "github.com/lib/pq"
)

// Database handling module

var db *sql.DB

func init() {
	var err error

	// Get database user config from env package
	dbConfig := Env.GetDatabaseConfig()

	// Create a postgres database configuration
	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v port=%v sslmode=disable",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Name, dbConfig.Port)
	fmt.Println(connStr)

	// Open database using configuration
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
}

func GetDB() *sql.DB {
	return db
}