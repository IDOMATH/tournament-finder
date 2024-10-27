package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/tournament-finder/repository"

	render "github.com/IDOMATH/CheetahRender/Render"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	repo := repository.Repository{}
	rr := render.NewRenderer("./templates", ".go.html", "./templates/partials", ".go.html", true)

	repo.RR = *rr

	router.HandleFunc("GET /", repo.HandleHome)

	fmt.Println("Starting on port 8080")
	log.Fatal(server.ListenAndServe())
}
