package store

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/wesleimp/dc-test/store/migrate/postgres"
)

var (
	driver = "postgres"
)

// Connect to a database and verify with a ping.
func Connect(datasource string) (*sqlx.DB, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, driver)
	if err := setupDatabase(dbx); err != nil {
		return nil, err
	}

	return dbx, nil
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}

func setupDatabase(db *sqlx.DB) error {
	return postgres.Migrate(db)
}
