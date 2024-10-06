package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/IDOMATH/tournament-finder/types"
)

type TournamentStore struct {
	DB *sql.DB
}

func NewTournamentStore(db *sql.DB) *TournamentStore {
	return &TournamentStore{
		DB: db,
	}
}

func (s *TournamentStore) InsertTournament(tournament types.Tournament) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	//TODO: figure out how we want to store this all in a database.
	statement := `insert into tournaments `

	err := s.DB.QueryRowContext(ctx, statement).Scan(&newId)

	if err != nil {
		return 0, err
	}
	return newId, nil
}
