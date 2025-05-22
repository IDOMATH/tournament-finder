package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("getting user from body"))
		return
	}

	newId, err := repo.US.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error inserting user"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Inserted user with ID: %d", newId)))
}
