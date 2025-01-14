package db

import "database/sql"

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		Db: db,
	}
}
