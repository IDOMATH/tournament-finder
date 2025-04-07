package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/middleware"
	"github.com/IDOMATH/tournament-finder/repository"
	"github.com/IDOMATH/tournament-finder/util"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	dbHost := util.GetEnvValue("DBHOST", "localhost")
	dbPort := util.GetEnvValue("DBPORT", "5432")
	dbName := util.GetEnvValue("DBNAME", "portfolio")
	dbUser := util.GetEnvValue("DBUSER", "postgres")
	dbPass := util.GetEnvValue("DBPASS", "postgres")
	dbSsl := util.GetEnvValue("DBSSL", "disable")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSsl)
	fmt.Println("Connecting to Postgres")
	postgresDb, err := db.ConnectSql(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Postgres")

	repo := repository.Repository{}

	TournamentStore := *db.NewTournamentStore(postgresDb.SQL)
	repo.TH = repository.TournamentHandler{TournamentStore: TournamentStore}

	router.HandleFunc("GET /tournaments", middleware.Log(middleware.Authenticate(repo.HandleGetTournaments, &repo)))
	router.HandleFunc("POST /tournaments", repo.HandlePostTournament)
	router.HandleFunc("PUT /tournaments/{id}", repo.HandlePutTournament)
	router.HandleFunc("GET /tournaments/{id}", repo.HandleGetTournamentById)
	router.HandleFunc("DELETE /tournaments/{id}", repo.HandleDeleteTournament)
	router.HandleFunc("GET /user/{id}", repo.HandleGetUserById)
	router.HandleFunc("POST /user", repo.HandlePostNewUser)

	fmt.Println("Starting on port 8080")
	log.Fatal(server.ListenAndServe())
}
