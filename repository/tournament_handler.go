package repository

import (
	"net/http"

	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
)

type TournamentHandler struct {
	tournamentStore db.TournamentStore
}

func (h *TournamentHandler) HandleGetTournaments(w http.ResponseWriter, r *http.Request) {
	td := types.TemplateData{
		PageName:  "All Tournaments",
		ObjectMap: make(map[string]interface{}),
	}
	tournaments, err := h.tournamentStore.GetAllTournaments()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	td.ObjectMap["tournaments"] = tournaments

	//TODO: render template
}
