package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Declare the default database driver
const defaultDatabaseDriver string = "sqlite"

// Declare the valid database drivers
var validDatabaseDrivers []string = []string{"sqlite", "mysql"}

// Get the database type to use
func getDbType() string {
	// Grab the environment variable
	db, test := os.LookupEnv(`DATABASE_DRIVER`)

	// Check if the test has failed
	if !test {
		return defaultDatabaseDriver
	}

	for _, v := range validDatabaseDrivers {
		// Check if the provided database corresponds to this entry
		if v == db {
			// This is a valid database type
			return db
		}
	}

	// No matches
	return defaultDatabaseDriver
}

func GetConn() (*sql.DB, error) {
	// Check what kind of database we are looking for
	dbConnType := getDbType()

	if dbConnType == `mysql` {
		return GetMysqlConn()
	} else {
		return GetLiteConn()
	}

}

func GetSqliteDatabasePath() string {
	return "main.db"
}

func GetLiteConn() (*sql.DB, error) {
	// Declare the database file name
	dbFile := GetSqliteDatabasePath()

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetMysqlConn() (*sql.DB, error) {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dburl := os.Getenv("MYSQL_URL")
	dbname := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dburl, dbname))

	if err != nil {
		return nil, err
	}

	// Set options
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
