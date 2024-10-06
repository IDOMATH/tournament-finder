package repository

import "github.com/IDOMATH/tournament-finder/db"

type TournamentHandler struct {
	tournamentStore db.TournamentStore
}
