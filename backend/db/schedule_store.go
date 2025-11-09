package db

import "database/sql"

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
