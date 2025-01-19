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

func (s *UserStore) GetUserById(id int) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u types.User

	query := `select name, email from users where id = $1`

	err := s.Db.QueryRowContext(ctx, query, id).Scan(&u.Name, &u.Email)
	return u, err
}

func (s *UserStore) UpdateUser(u types.User, id int) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedUser types.User

	query := `upate users set name = $1, email = $1 where id = $3
			  returning name, email, id`

	err := s.Db.QueryRowContext(ctx, query, u.Name, u.Email, u.Id).Scan(&updatedUser.Name, &updatedUser.Email, &updatedUser.Id)

	return updatedUser, err

}
