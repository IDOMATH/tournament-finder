package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/IDOMATH/tournament-finder/repository"
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
			// repo.Session.GetToken(r.Header.Get("cheetauth"))

			fmt.Println("NOT AUTHENTICATED")
			next(w, r)
		}
	}
}
