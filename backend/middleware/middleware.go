package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/IDOMATH/tournament-finder/repository"
	"github.com/IDOMATH/tournament-finder/util"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Use(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func Log() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%v:\t%s\t%s", time.Now(), r.RemoteAddr, r.URL)
			next(w, r)
		}
	}
}

func Authenticate(repo *repository.Repository) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Make and hook up session storage to repository
			t, found, _ := repo.Session.Get(r.Header.Get("cheetauth"))
			if !found {
				fmt.Println("NOT AUTHENTICATED")
				// Potentially do some rerouting if the endpoint is protected
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(r.Header.Get("cheetauth")))
				return
			}
			id := util.IntifyId(t)
			fmt.Println(id)

			next(w, r)
		}
	}
}

func Authorize(repo *repository.Repository) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			t, found, _ := repo.Session.Get(r.Header.Get("cheetauth"))
			if !found {
				fmt.Println("NOT AUTHENTICATED")
				// Potentially do some rerouting if the endpoint is protected
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(r.Header.Get("cheetauth")))
				return
			}
			id := util.IntifyId(t)
			tournamentId, err := strconv.Atoi(r.PathValue("id"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				// Something like route to the previous page?
				return
			}
			organizerId, err := repo.TS.GetTournamentOrganizerId(tournamentId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if id != organizerId {
				w.WriteHeader(http.StatusUnauthorized)
				// Something like route to previous page?
				return
			}

			next(w, r)
		}
	}
}
