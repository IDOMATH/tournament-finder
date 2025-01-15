package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/IDOMATH/tournament-finder/types"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		Db: db,
	}
}

func (s *UserStore) InsertUser(user types.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	statement := `insert into users (name, email, created_at) values ($1, $2, $3)`

	err := s.Db.QueryRowContext(ctx, statement,
		user.Name,
		user.Email, time.Now()).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}
