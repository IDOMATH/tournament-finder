package main

import (
	"fmt"
	"log"
	"net/http"

	render "github.com/IDOMATH/CheetahRender/Render"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	repo := Repository{}
	rr := render.NewRenderer("./templates", ".go.html", "./templates/partials", ".go.html", true)

	fmt.Println("Starting on port 8080")
	log.Fatal(server.ListenAndServe())
}
