package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/IDOMATH/tournament-finder/types"
)

type ScheduleStore struct {
	Db *sql.DB
}

func NewScheduleStore(db *sql.DB) *ScheduleStore {
	return &ScheduleStore{
		Db: db,
	}
}

func (s *ScheduleStore) AddTournamentToCoachSchedule(coachId, tournamentId int) error {

	return nil
}

func (s *TournamentStore) GetScheduleByCoachId(id int) ([]types.Tournament, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tournaments []types.Tournament
	query := `SELECT name, location_name,
	street_address, city, state,
	is_boy_varsity, is_girls_varsity, is_boys_jv, is_girls_js,
    is_boys_ms, is_girls_ms, is_boys_youth, is_girls_youth,
	organizer_id, is_full 
	FROM tournaments WHERE coach_id = $1`

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
func (s *ScheduleStore) DeleteTournamentFromSchedule(coachId, tournamentId string) error {
	query := `DELETE FROM schedule WHERE coach_id = ? AND tournament_id = ?`
}
