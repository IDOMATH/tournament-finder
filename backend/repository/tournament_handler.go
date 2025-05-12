package repository

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
)

type TournamentHandler struct {
	TournamentStore db.TournamentStore
}

func parseTournamentForm(r *http.Request) types.Tournament {
	var tournament types.Tournament

	tournament.Name = r.PostFormValue("name")
	tournament.LocationName = r.PostFormValue("locationName")
	tournament.StreetAddress = r.PostFormValue("streetAddress")
	tournament.City = r.PostFormValue("city")
	tournament.State = r.PostFormValue("state")
	//TODO: Make some sort of date parser function.  Probably in the utils package
	// tournament.StartDate = util.ParseDate(r.PostFormValue("startDate"))
	// tournament.EndDate = util.ParseDate(r.PostFormValue("endDate"))
	tournament.BoysVarsity = r.PostFormValue("boysVarsity")
	tournament.GirlsVarsity = r.PostFormValue("girlsVarsity")
	tournament.BoysJv = r.PostFormValue("boysJv")
	tournament.GirlsJv = r.PostFormValue("girlsJv")
	tournament.BoysMs = r.PostFormValue("boysMs")
	tournament.GirlsMs = r.PostFormValue("girlsMs")
	tournament.BoysYouth = r.PostFormValue("boysYouth")
	tournament.GirlsYouth = r.PostFormValue("girlsYouth")

	return tournament
}

func (repo *Repository) HandlePostTournament(w http.ResponseWriter, r *http.Request) {
	tournament := parseTournamentForm(r)

	//TODO: get the organizerId from the logged in user.
	// tournament.OrganizerId = repo.Session.

	id, err := repo.TH.TournamentStore.InsertTournament(tournament)
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
	tournament.Name = r.FormValue("name")

	updatedTournament, err := repo.TH.TournamentStore.UpdateTournament(tournament)
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
	tournaments, err := repo.TH.TournamentStore.GetAllTournaments(pageNumber)
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
	tournament, err := repo.TH.TournamentStore.GetTournamentById(id)
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
	err = repo.TH.TournamentStore.DeleteTournament(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error deleting tournament: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
