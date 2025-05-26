package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/IDOMATH/tournament-finder/types"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		Db: db,
	}
}

func (s *UserStore) InsertUser(user types.NewUser) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	statement := `insert into users (name, email, updated_at, created_at) values ($1, $2, $3)`

	err := s.Db.QueryRowContext(ctx, statement,
		user.Name,
		user.Email, time.Now(), time.Now()).Scan(&newId)
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

	query := `upate users set name = $1, email = $1, updated_at = $4 where id = $3
			  returning name, email, id`

	err := s.Db.QueryRowContext(ctx, query, u.Name, u.Email, time.Now(), u.Id).Scan(&updatedUser.Name, &updatedUser.Email, &updatedUser.Id)

	return updatedUser, err
}

func (s *UserStore) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from users where id = $1`

	_, err := s.Db.ExecContext(ctx, query, id)
	return err
}

func (s *UserStore) Login(email, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u types.User

	query := `select email, password_hash, id from users where email = $1`
	err := s.Db.QueryRowContext(ctx, query, email).Scan(&u.Email, &u.PasswordHash, &u.Id)
	if err != nil {
		return 0, errors.New("error getting user from database")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	if err != nil {
		return 0, errors.New("incorrect password")
	}
	return u.Id, nil
}
