package database

import "database/sql"

type DB struct {
	CategoryDB *sql.DB
}

func NewCategoryDB(db *sql.DB) DB {
	return DB{
		CategoryDB: db,
	}
}
