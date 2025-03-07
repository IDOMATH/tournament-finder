package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
)

type UserHandler struct {
	UserStore db.UserStore
}

func (repo *Repository) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error converting id"))
		return
	}

	user, err := repo.UH.UserStore.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting user"))
		return
	}
	resUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resUser)
}

func (repo *Repository) HandlePostNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")

	newId, err := repo.UH.UserStore.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Inserted user with ID: %d", newId)))
}
