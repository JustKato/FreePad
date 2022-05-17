package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateMysql() error {

	m, err := migrate.New(
		"file://db/migrations/",
		GetMysqlString(),
	)
	if err != nil {
		return err
	}

	// Migrate
	err = m.Up()
	if err != nil {
		return err
	}

	return m.Run()
}

// Run migrations to ensure tables exist
func MigrationUpdate() *migrate.Logger {
	// Get the path to the sqlite database
	databasePath := fmt.Sprintf("sqlite://%s", GetSqliteDatabasePath())

	// Try and create a new migration
	m, err := migrate.New(
		"file://../db/migrations_sqlite",
		databasePath,
	)

	if err != nil {
		// End the whole thing if migrations fail
		panic(err)
	}

	// Run the update
	err = m.Up()
	m.Run()

	m.Force(1)

	return &m.Log

}
