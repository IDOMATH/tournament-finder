package repository

import (
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
		w.Write([]byte("error converting id"))
		return
	}

	user, err := repo.UH.UserStore.GetUserById(id)
	if err != nil {
		w.Write([]byte("error getting user"))
		return
	}

	w.Write([]byte(fmt.Sprintf("getting user with id: %d", user.Id)))
}

func (repo *Repository) HandlePosNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")

	newId, err := repo.UH.UserStore.InsertUser(user)
	if err != nil {
		w.Write([]byte("error inserting user"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("Inserted user with ID: %d", newId)))
}
