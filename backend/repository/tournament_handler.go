package repository

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/types"
	"github.com/IDOMATH/tournament-finder/util"
)

func (repo *Repository) HandlePostTournament(w http.ResponseWriter, r *http.Request) {
	var tournament types.Tournament
	json.NewDecoder(r.Body).Decode(&tournament)

	//TODO: get the organizerId from the logged in user.
	tokenVal, found, err := repo.Session.Get(r.Header["cheetauth"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tournament.OrganizerId = util.IntifyId(tokenVal)
	if tournament.OrganizerId == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := repo.TS.InsertTournament(tournament)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (repo *Repository) HandlePutTournament(w http.ResponseWriter, r *http.Request) {
	var tournament types.Tournament
	json.NewDecoder(r.Body).Decode(&tournament.Name)

	updatedTournament, err := repo.TS.UpdateTournament(tournament)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(updatedTournament)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (repo *Repository) HandleGetTournaments(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.PathValue("page"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tournaments, err := repo.TS.GetAllTournaments(pageNumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resTourments, err := json.Marshal(tournaments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resTourments)
}

func (repo *Repository) HandleGetTournamentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid id"))
		return
	}
	tournament, err := repo.TS.GetTournamentById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resTournament, err := json.Marshal(tournament)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resTournament)
}

func (repo *Repository) HandleDeleteTournament(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error converting ID from string to int: " + err.Error()))
		return
	}
	err = repo.TS.DeleteTournament(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error deleting tournament: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
