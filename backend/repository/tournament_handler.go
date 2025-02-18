package repository

import (
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
)

type TournamentHandler struct {
	TournamentStore db.TournamentStore
}

func (repo *Repository) HandleGetNewTournamentForm(w http.ResponseWriter, r *http.Request) {
	td := types.TemplateData{
		PageName:  "New Tournament",
		ObjectMap: make(map[string]interface{}),
	}

	repo.RR.Render(w, r, "tournament-form.go.html", td)
}

func (repo *Repository) HandlePostTournament(w http.ResponseWriter, r *http.Request) {
	var tournament types.Tournament
	tournament.Name = r.FormValue("name")

	repo.TH.TournamentStore.InsertTournament(tournament)
}

func (repo *Repository) HandlePutTournament(w http.ResponseWriter, r *http.Request) {
	var tournament types.Tournament
	tournament.Name = r.FormValue("name")

	repo.TH.TournamentStore.UpdateTournament(tournament)

	td := types.TemplateData{
		PageName: tournament.Name,
	}

	repo.RR.Render(w, r, "tournament.go.html", td)
}

func (repo *Repository) HandleGetTournaments(w http.ResponseWriter, r *http.Request) {
	td := types.TemplateData{
		PageName:  "All Tournaments",
		ObjectMap: make(map[string]interface{}),
	}
	tournaments, err := repo.TH.TournamentStore.GetAllTournaments()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	td.ObjectMap["tournaments"] = tournaments

	repo.RR.Render(w, r, "all-tournaments.go.html", td)
}

func (repo *Repository) HandleGetTournamentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("error converting ID from string to int: " + err.Error()))
		return
	}
	tournament, err := repo.TH.TournamentStore.GetTournamentById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	td := types.TemplateData{
		PageName:  tournament.Name,
		ObjectMap: make(map[string]interface{}),
	}

	repo.RR.Render(w, r, "tournament.go.html", td)
}

func (repo *Repository) HandleDeleteTournament(w http.ResponseWriter, r *http.Request) {
	td := types.TemplateData{
		PageName:  "Deleted",
		ObjectMap: make(map[string]interface{}),
	}

	w.Write([]byte("Deleting tournament"))
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("error converting ID from string to int: " + err.Error()))
		td.ObjectMap["error"] = err.Error()
		return
	}
	err = repo.TH.TournamentStore.DeleteTournament(id)
	if err != nil {
		w.Write([]byte("error deleting tournament: " + err.Error()))
		td.ObjectMap["error"] = err.Error()
		// TODO: Render some error page or the tournament details page
		return
	}

	repo.RR.Render(w, r, "tournament-deleted.go.html", td)
}
