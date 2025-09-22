package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/session/memorystore"
	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/middleware"
	"github.com/IDOMATH/tournament-finder/repository"
	"github.com/IDOMATH/tournament-finder/util"
)

func main() {

	server := setup()

	fmt.Println("Starting on port ", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func setup() *http.Server {

	serverPort := util.GetEnvValue("SERVEPORT", "8080")

	router := http.NewServeMux()
	server := http.Server{
		Addr:    fmt.Sprint(":", serverPort),
		Handler: router,
	}

	postgresDb := setupDbConnection()

	repo := repository.Repository{}

	repo.TS = *db.NewTournamentStore(postgresDb.SQL)
	repo.US = *db.NewUserStore(postgresDb.SQL)

	memstore := memorystore.New[string]()
	repo.Session = memstore

	repo.Logger = util.NewLogger("/logs/log.txt")

	registerRoutes(router, &repo)

	return &server
}

func setupDbConnection() *db.DB {
	dbHost := util.GetEnvValue("DBHOST", "localhost")
	dbPort := util.GetEnvValue("DBPORT", "5432")
	dbName := util.GetEnvValue("DBNAME", "tournament-finder")
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
	return postgresDb
}

func registerRoutes(router *http.ServeMux, repo *repository.Repository) {
	stack := []middleware.Middleware{middleware.Authenticate(repo), middleware.Log()}

	router.HandleFunc("GET /", middleware.Use(handleLoginTest, stack...))

	router.HandleFunc("GET /tournaments", middleware.Use(repo.HandleGetTournaments, stack...))
	router.HandleFunc("POST /tournaments", middleware.Use(repo.HandlePostTournament, stack...))
	router.HandleFunc("PUT /tournaments/{id}", repo.HandlePutTournament)
	router.HandleFunc("GET /tournaments/{id}", repo.HandleGetTournamentById)
	router.HandleFunc("DELETE /tournaments/{id}", repo.HandleDeleteTournament)
	router.HandleFunc("GET /user/{id}", repo.HandleGetUserById)
	router.HandleFunc("POST /user", repo.HandlePostNewUser)
	router.HandleFunc("POST /login", repo.HandleLogin)
	router.HandleFunc("POST /logout", repo.HandleLogout)
}

func handleLoginTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logged in"))
}
