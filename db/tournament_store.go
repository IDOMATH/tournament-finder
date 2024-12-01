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

func (s *TournamentStore) UpdateTournament(tournament types.Tournament) (types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedTournament types.Tournament

	//TODO: add fields to statement and execution
	statement := `update tournaments set ... where id = $1`

	res, err := s.DB.ExecContext(ctx, statement)

	if err != nil {
		return updatedTournament, err
	}

	return updatedTournament, nil
}

func (s *TournamentStore) GetAllTournaments() ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament

	//TODO: figure out fields in database
	query := `select * from tournaments`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.Name,
		)
		if err != nil {
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}

	if err = rows.Err(); err != nil {
		return tournaments, err
	}

	return tournaments, nil
}

func (s *TournamentStore) GetTournamentById(id int) (types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournament types.Tournament
	query := `select * from tournaments where id = $1`

	err := s.DB.QueryRowContext(ctx, query, id).Scan(tournament)
	if err == sql.ErrNoRows {
		return tournament, err
	}
	if err != nil {
		return tournament, err
	}
	return tournament, nil
}

func (s *TournamentStore) DeleteTournament(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournament types.Tournament
	query := `DELETE FROM tournaments where id = $1`

	err := s.DB.QueryRowContext(ctx, query, id).Scan(tournament)
	if err != nil {
		return err
	}
	return nil
}
