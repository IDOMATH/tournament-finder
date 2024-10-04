package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting on port 8080")
	log.Fatal(server.ListenAndServe())
}
