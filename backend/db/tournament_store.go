package db

import (
	"context"
	"database/sql"
	"fmt"
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
	statement := `
	insert into tournaments 
	(name, location_name, location_address, organizer_id, age_division, is_full) 
	values ($1, $2, $3, $4, $5, $6)`

	err := s.DB.QueryRowContext(ctx, statement,
		tournament.Name,
		tournament.LocationName,
		tournament.LocationAddress,
		tournament.OrganizerId,
		tournament.AgeDivisionArrayToInt(),
		tournament.IsFull).Scan(&newId)

	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (s *TournamentStore) UpdateTournament(tournament types.Tournament) (types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedTournament types.Tournament

	statement := `update tournaments 
	set name = $1, location_name = $2, location_address = $3, 
	organizer_id = $4, age_division = $5, is_full = $6
	where id = $7`

	err := s.DB.QueryRowContext(ctx, statement,
		tournament.Name,
		tournament.LocationName,
		tournament.LocationAddress,
		tournament.OrganizerId,
		tournament.AgeDivisionArrayToInt(),
		tournament.IsFull,
		tournament.Id).Scan(&updatedTournament)

	fmt.Println(updatedTournament)

	if err != nil {
		return updatedTournament, err
	}

	return updatedTournament, nil
}

func (s *TournamentStore) GetAllTournaments() ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament

	query := `select name, location_name, location_address, organizer_id, age_division, is_full 
	from tournaments`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return tournaments, err
	}
	defer rows.Close()

	var ageDivision int

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.Name,
			&tournament.LocationName,
			&tournament.LocationAddress,
			&tournament.OrganizerId,
			&ageDivision,
			&tournament.IsFull,
		)
		tournament.AgeDivisionIntToArray(ageDivision)
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

func (s *TournamentStore) FilterTournaments(filter types.Tournament) ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament
	numActiveFilters := 0
	var activeFilters []any

	query := `select name, location_name, location_address, organizer_id,
	 age_division, is_full from tournaments where`

	if filter.Name != "" {
		numActiveFilters++
		query = query + fmt.Sprintf("name = $%d", numActiveFilters)
		activeFilters = append(activeFilters, filter.Name)
	}

	rows, err := s.DB.QueryContext(ctx, query, activeFilters...)
	if err != nil {
		return tournaments, err
	}
	defer rows.Close()

	return tournaments, nil
}

func (s *TournamentStore) GetTournamentById(id int) (types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournament types.Tournament
	query := `select name, location_name, location_address, organizer_id,
	 age_division, is_full from tournaments where id = $1`

	var ageDivision int

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&tournament.Name,
		&tournament.LocationName,
		&tournament.LocationAddress,
		&tournament.OrganizerId,
		&ageDivision,
		&tournament.IsFull,
	)
	tournament.AgeDivisionIntToArray(ageDivision)
	return tournament, err
}

func (s *TournamentStore) DeleteTournament(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM tournaments where id = $1`

	_, err := s.DB.ExecContext(ctx, query, id)
	return err
}
