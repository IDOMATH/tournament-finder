package repository

import (
	"net/http"

	render "github.com/IDOMATH/CheetahRender/Render"
	"github.com/IDOMATH/tournament-finder/types"
)

type Repository struct {
	RR render.Renderer
	TH TournamentHandler
	UH UserHandler
}

func (repo *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	repo.RR.Render(w, r, "home.go.html", types.TemplateData{PageName: "Home"})
}
