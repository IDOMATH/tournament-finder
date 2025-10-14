package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (s *UserStore) InsertUser(user types.NewUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	statement := `insert into users (name, email, password_hash, is_organizer, is_coach, updated_at, created_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err = s.Db.ExecContext(ctx, statement,
		user.Name,
		user.Email,
		passwordHash,
		user.IsOrganizer,
		user.IsCoach,
		time.Now(),
		time.Now())

	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetUserById(id int) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u types.User

	query := `select name, email, is_organizer, is_coach from users where id = $1`

	err := s.Db.QueryRowContext(ctx, query, id).Scan(&u.Name, &u.Email, &u.IsOrganizer, &u.IsCoach)
	return u, err
}

func (s *UserStore) UpdateUser(u types.User, id int) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedUser types.User

	query := `upate users set name = $1, email = $2, is_organizer = $3, is_coach = $4, updated_at = $5 where id = $6
			  returning name, email, is_organizer, is_coach, id`

	err := s.Db.QueryRowContext(ctx, query, u.Name, u.Email, u.IsOrganizer, u.IsCoach, time.Now(), u.Id).Scan(&updatedUser.Name, &updatedUser.Email, &updatedUser.IsOrganizer, &updatedUser.IsCoach, &updatedUser.Id)

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
		fmt.Println(err.Error())
		return 0, errors.New("error getting user from database")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	if err != nil {
		return 0, errors.New("incorrect password")
	}
	return u.Id, nil
}
