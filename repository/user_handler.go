package repository

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IDOMATH/tournament-finder/db"
)

type UserHandler struct {
	UserStore db.UserStore
}

func (repo *Repository) HandleGetUser(w http.ResponseWriter, r *http.Request) {
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
