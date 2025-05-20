package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/types"
)

func (repo *Repository) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
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

	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")

	newId, err := repo.US.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error inserting user"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Inserted user with ID: %d", newId)))
}
