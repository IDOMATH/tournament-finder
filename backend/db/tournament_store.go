package db

import (
	"context"
	"database/sql"
	"errors"
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
	(name, location_name, street_address, city, state, organizer_id,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	is_full) 
	values ($1, $2, $3, $4, $5, $6, $7)`

	err := s.DB.QueryRowContext(ctx, statement,
		tournament.Name,
		tournament.LocationName,
		tournament.StreetAddress,
		tournament.City,
		tournament.State,
		&tournament.BoysVarsity,
		&tournament.GirlsVarsity,
		&tournament.BoysJv,
		&tournament.GirlsJv,
		&tournament.BoysMs,
		&tournament.GirlsMs,
		&tournament.BoysYouth,
		&tournament.GirlsYouth,
		tournament.OrganizerId).Scan(&newId)

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
	set name = $1, location_name = $2,
	street_address = $3, city = $4, state = $5,
	is_boy_varsity = $6, is_girls_varsity = $7, is_boys_jv = $8, is_girls_js = $9,
    is_boys_ms = $10, is_girls_ms = $11, is_boys_youth = $12, is_girls_youth = $13,
	organizer_id = $14, is_full = $15
	where id = $16`

	err := s.DB.QueryRowContext(ctx, statement,
		tournament.Name,
		tournament.LocationName,
		tournament.StreetAddress,
		tournament.City,
		tournament.State,
		&tournament.BoysVarsity,
		&tournament.GirlsVarsity,
		&tournament.BoysJv,
		&tournament.GirlsJv,
		&tournament.BoysMs,
		&tournament.GirlsMs,
		&tournament.BoysYouth,
		&tournament.GirlsYouth,
		tournament.OrganizerId,
		tournament.Id).Scan(&updatedTournament)

	fmt.Println(updatedTournament)

	if err != nil {
		return updatedTournament, err
	}

	return updatedTournament, nil
}

func (s *TournamentStore) GetAllTournaments(page int) ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament

	offset := 25 * (page - 1)

	// TODO: Add pagination
	query := `select name, location_name,
	street_address, city, state,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	organizer_id, is_full,
	start_date
	from tournaments
	ORDER BY start_date
	LIMIT 25
	OFFSET $1`

	rows, err := s.DB.QueryContext(ctx, query, offset)
	if err != nil {
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.Name,
			&tournament.LocationName,
			&tournament.StreetAddress,
			&tournament.City,
			&tournament.State,
			&tournament.BoysVarsity,
			&tournament.GirlsVarsity,
			&tournament.BoysJv,
			&tournament.GirlsJv,
			&tournament.BoysMs,
			&tournament.GirlsMs,
			&tournament.BoysYouth,
			&tournament.GirlsYouth,
			&tournament.OrganizerId,
			&tournament.IsFull,
			&tournament.StartDate)
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
	var activeFilters []any

	query := `select name, location_name, 
	street_address, city, state,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	organizer_id, is_full 
	from tournaments where`

	// This might want to be LIKE instead of =
	if filter.Name != "" {
		activeFilters = append(activeFilters, filter.Name)
		query = query + fmt.Sprintf("name = $%d", len(activeFilters)+1)
	}

	if filter.StreetAddress != "" {
		activeFilters = append(activeFilters, filter.StreetAddress)
		query = query + fmt.Sprintf("street_address = $%d", len(activeFilters)+1)
	}

	if filter.City != "" {
		activeFilters = append(activeFilters, filter.City)
		query = query + fmt.Sprintf("city = $%d", len(activeFilters)+1)
	}

	if filter.State != "" {
		activeFilters = append(activeFilters, filter.State)
		query = query + fmt.Sprintf("location_state = $%d", len(activeFilters)+1)
	}

	if !filter.StartDate.IsZero() {
		activeFilters = append(activeFilters, filter.StartDate)
		query = query + fmt.Sprintf("start_date = $%d", len(activeFilters)+1)
	}

	if !filter.EndDate.IsZero() {
		activeFilters = append(activeFilters, filter.EndDate)
		query = query + fmt.Sprintf("end_date = $%d", len(activeFilters)+1)
	}

	if filter.BoysVarsity > 0 {
		query = query + "is_boys_varsity = true"
	}
	if filter.GirlsVarsity > 0 {
		query = query + "is_girls_varsity = true"
	}
	if filter.BoysJv > 0 {
		query = query + "is_boys_jv = true"
	}
	if filter.GirlsJv > 0 {
		query = query + "is_girls_jv = true"
	}
	if filter.BoysMs > 0 {
		query = query + "is_boys_ms = true"
	}
	if filter.GirlsMs > 0 {
		query = query + "is_girls_ms = true"
	}
	if filter.BoysYouth > 0 {
		query = query + "is_boys_youth = true"
	}
	if filter.GirlsYouth > 0 {
		query = query + "is_girls_youth = true"
	}

	if len(activeFilters) == 0 {
		return tournaments, errors.New("no filters given")
	}

	rows, err := s.DB.QueryContext(ctx, query, activeFilters...)
	if err != nil {
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.Name,
			&tournament.LocationName,
			&tournament.StreetAddress,
			&tournament.City,
			&tournament.State,
			&tournament.BoysVarsity,
			&tournament.GirlsVarsity,
			&tournament.BoysJv,
			&tournament.GirlsJv,
			&tournament.BoysMs,
			&tournament.GirlsMs,
			&tournament.BoysYouth,
			&tournament.GirlsYouth,
			&tournament.OrganizerId,
		)
		if err != nil {
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func (s *TournamentStore) GetTournamentById(id int) (types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournament types.Tournament
	query := `select name, location_name,
	street_address, city, state,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	organizer_id, is_full 
	from tournaments where id = $1`

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&tournament.Name,
		&tournament.LocationName,
		&tournament.StreetAddress,
		&tournament.City,
		&tournament.State,
		&tournament.BoysVarsity,
		&tournament.GirlsVarsity,
		&tournament.BoysJv,
		&tournament.GirlsJv,
		&tournament.BoysMs,
		&tournament.GirlsMs,
		&tournament.BoysYouth,
		&tournament.GirlsYouth,
		&tournament.OrganizerId,
		&tournament.IsFull,
	)
	return tournament, err
}

func (s *TournamentStore) GetTournamentOrganizerId(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var organizerId int
	query := `select organizer_id 
	from tournaments where id = $1`

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&organizerId,
	)
	return organizerId, err
}

func (s *TournamentStore) GetAllTournamentsByOrganizerId(id int) ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament
	query := `SELECT name, location_name,
	street_address, city, state,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	organizer_id, is_full 
	FROM tournaments WHERE organizer_id = $1`

	rows, err := s.DB.QueryContext(ctx, query, id)
	if err != nil {
		return tournaments, err
	}

	defer rows.Close()

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.Name,
			&tournament.LocationName,
			&tournament.StreetAddress,
			&tournament.City,
			&tournament.State,
			&tournament.BoysVarsity,
			&tournament.GirlsVarsity,
			&tournament.BoysJv,
			&tournament.GirlsJv,
			&tournament.BoysMs,
			&tournament.GirlsMs,
			&tournament.BoysYouth,
			&tournament.GirlsYouth,
			&tournament.OrganizerId,
			&tournament.IsFull,
		)
		if err != nil {
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}
	return tournaments, nil
}

func (s *TournamentStore) DeleteTournament(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM tournaments where id = $1`

	_, err := s.DB.ExecContext(ctx, query, id)
	return err
}
