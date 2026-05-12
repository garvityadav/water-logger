package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// database connection initiator
func Connect(connectionString string) (*sql.DB, error) {
	db, dbErr := sql.Open("postgres", connectionString)
	if dbErr != nil {
		return nil, fmt.Errorf("database.Connect: %w", dbErr)
	}

	// ping DB
	dbPingErr := db.Ping()
	if dbPingErr != nil {
		return nil, fmt.Errorf("database.Connect: %w", dbPingErr)
	}
	return db, nil
}
