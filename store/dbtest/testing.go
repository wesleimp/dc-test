package dbtest

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/wesleimp/dc-test/store/migrate/postgres"
)

const (
	dbDsn     = "host=localhost port=5432 user=postgres password=postgres dbname=dctest sslmode=disable"
	testDbDsn = "host=localhost port=5432 user=postgres password=postgres dbname=dctest_test_%s sslmode=disable"
)

// Setup creates a database test
func Setup(name string) *sqlx.DB {
	err := setup(name)
	if err != nil {
		panic("Error setting up database test")
	}
	db, _ := migrate(name)
	return db
}

// TearDown deletes the database test
func TearDown(db *sqlx.DB, name string) {
	err := tearDown(db, name)
	if err != nil {
		panic("Error shutting down database")
	}
}

func setup(name string) error {
	log.Println("Setup Database")
	db := connect(dbDsn)
	defer db.Close()
	return exec(db, fmt.Sprintf(createTestDatabase, name))
}

func migrate(name string) (*sqlx.DB, error) {
	db := connect(fmt.Sprintf(testDbDsn, name))
	err := postgres.Migrate(db)
	if err != nil {
		fmt.Printf("Unable to create migrations\n%s\n", err.Error())
	}
	return db, err
}

func tearDown(db *sqlx.DB, name string) error {
	log.Println("Teardown Database")
	db.Close()
	db = connect(dbDsn)
	defer db.Close()
	return exec(db, fmt.Sprintf(dropTestDatabase, name))
}

func connect(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Printf("error connecting to database\n%s\n", err.Error())
	}
	return db
}

func exec(db *sqlx.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("error handling test database\n%s\n", err.Error())
	}
	return nil
}

var createTestDatabase = `
CREATE DATABASE dctest_test_%s
	WITH 
	OWNER = postgres
	ENCODING = 'UTF8'
	LC_COLLATE = 'en_US.utf8'
	LC_CTYPE = 'en_US.utf8'
	TABLESPACE = pg_default
	CONNECTION LIMIT = -1;
`

var dropTestDatabase = `
DROP DATABASE IF EXISTS dctest_test_%s;
`
