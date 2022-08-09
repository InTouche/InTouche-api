package store

import "database/sql"

type userStore struct {
	db        *sql.DB
	tableName string
}
