package repository

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.gom/IDOMATH/tournament-finder/log"
)

func (repo *Repository) HandleAddTournamentToCoachSchedule(w http.ResponseWriter, r *http.Request) {
	strCoachId := r.PathValue("coachId")
	strTournamentId := r.PathValue("tournamentId")

	coachId, err := strconv.Atoi(strCoachId)
	if err != nil {
		log.Error("HandleAddTournamentToCoachSchedule", "error converting coach id to string", err)
		return
	}

	tournamentId, err := strconv.Atoi(strTournamentId)
	if err != nil {
		log.Error("HandleAddTournamentToCoachSchedule", "error converting tournament id to string", err)
		return
	}

	err = repo.SS.AddTournamentToCoachSchedule(coachId, tournamentId)
	if err != nil {
		log.Error("HandleAddTournamentToCoachSchedule", "error adding tournament to schedule", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *Repository) HandleRemoveFromSchedule(w http.ResponseWriter, r *http.Request) {
	strCoachId := r.PathValue("coachId")
	strTournamentId := r.PathValue("tournamentId")

	coachId, err := strconv.Atoi(strCoachId)
	if err != nil {
		log.Error("HandleRemoveFromSchedule", "error converting coach id to string", err)
		return
	}

	tournamentId, err := strconv.Atoi(strTournamentId)
	if err != nil {
		log.Error("HandleRemoveFromSchedule", "error converting tournament id to string", err)
		return
	}

	err = repo.SS.DeleteTournamentFromSchedule(coachId, tournamentId)
	if err != nil {
		log.Error("HandleAddTournamentToCoachSchedule", "error adding tournament to schedule", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *Repository) HandleGetSchedule(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Error("HandleGetSchedule", "error converting coach id to string", err)
		return
	}

	tournaments, err := repo.TS.GetScheduleByCoachId(id)
	if err != nil {
		log.Error("HandleGetSchedule", "error getting tournaments", err)
		return
	}

	res, err := json.Marshal(tournaments)
	if err != nil {
		log.Error("HandleGetSchedule", "error marshalling tournaments to json", err)
		return
	}

	w.Write(res)
}
