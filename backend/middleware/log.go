package middleware

import (
	"net/http"

	"github.com/idomath/tournament-finder/repository"
)

func Log(next http.HandlerFunc, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add something to actually log here
		next(w, r)
	}
}
