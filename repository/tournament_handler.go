package repository

import "github.com/IDOMATH/tournament-finder/db"

type TournamentHandler struct {
	tournamentStore db.TournamentStore
}

func (h *TournamentHandler) HandleGetAllTournaments() {
	tournaments := h.tournamentStore.GetAllTournaments
}
