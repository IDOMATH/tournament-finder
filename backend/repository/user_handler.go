package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/constants"
	"github.com/IDOMATH/tournament-finder/types"
)

func (repo *Repository) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error converting id"))
		return
	}

	user, err := repo.US.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting user"))
		return
	}
	resUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resUser)
}

func (repo *Repository) HandlePostNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("getting user from body"))
		return
	}

	err = repo.US.InsertUser(user)
	if err.Error() == constants.EmailInAlreadyInUse {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("email already in use"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error inserting user"))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Inserted new user"))
}

func (repo *Repository) HandleGetTournamentsByCoachId(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByCoachId", "error converting coach id to string", err)
		return
	}

	tournaments, err := repo.TS.GetAllTournamentsByCoachId(id)
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

func (repo *Repository) HandleAddTournamentToCoachSchedule(w http.ResponseWriter, r *http.Request) {
	strCoachId := r.PathValue("coachId")
	strTournamentId := r.PathValue("tournamentId")

	coachId, err := strconv.Atoi(strCoachId)
	if err != nil {
		repo.Logger.LogError("HandleAddTournamentToCoachSchedule", "error converting id to string", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *Repository) HandleGetTournamentsByOrganizerId(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByOrganizerId", "error converting organizer id to int", err)
		return
	}

	tournaments, err := repo.TS.GetAllTournamentsByOrganizerId(id)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByOrganizerId", "error getting tournament", err)
		return
	}

	res, err := json.Marshal(tournaments)
	if err != nil {
		repo.Logger.LogError("HandleGetTournamentsByOrganizerId", "error marshalling tournaments to json", err)
		return
	}

	w.Write(res)
}
