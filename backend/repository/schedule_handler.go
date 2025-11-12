package repository

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (repo *Repository) HandleAddTournamentToCoachSchedule(w http.ResponseWriter, r *http.Request) {
	strCoachId := r.PathValue("coachId")
	strTournamentId := r.PathValue("tournamentId")

	coachId, err := strconv.Atoi(strCoachId)
	if err != nil {
		repo.Logger.LogError("HandleAddTournamentToCoachSchedule", "error converting coach id to string", err)
		return
	}

	tournamentId, err := strconv.Atoi(strTournamentId)
	if err != nil {
		repo.Logger.LogError("HandleAddTournamentToCoachSchedule", "error converting tournament id to string", err)
		return
	}

	err = repo.SS.AddTournamentToCoachSchedule(coachId, tournamentId)
	if err != nil {
		repo.Logger.LogError("HandleAddTournamentToCoachSchedule", "error adding tournament to schedule", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *Repository) HandleGetTournamentsByCoachId(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByCoachId", "error converting coach id to string", err)
		return
	}

	tournaments, err := repo.TS.GetScheduleByCoachId(id)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByCoachId", "error getting tournaments", err)
		return
	}

	res, err := json.Marshal(tournaments)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByCoachId", "error marshalling tournaments to json", err)
		return
	}

	w.Write(res)
}
