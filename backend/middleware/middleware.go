package middleware

import (
	"fmt"
	"net/http"

	"github.com/IDOMATH/tournament-finder/repository"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add something to actually log here
		fmt.Println("Logging...")
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc, repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("NOT AUTHENTICATED")
		next(w, r)
	}
}
