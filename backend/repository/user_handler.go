package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IDOMATH/tournament-finder/constants"
	"github.com/IDOMATH/tournament-finder/types"
)

func (repo *Repository) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error converting id"))
		return
	}

	user, err := repo.US.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting user"))
		return
	}
	resUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling json"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resUser)
}

func (repo *Repository) HandlePostNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("getting user from body"))
		return
	}

	err = repo.US.InsertUser(user)
	if err.Error() == constants.EmailInAlreadyInUse {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("email already in use"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error inserting user"))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Inserted new user"))
}

func HandleGetTournamentsByCoachId(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

}

func HandleGetTournamentsByOrganizerId(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
}
