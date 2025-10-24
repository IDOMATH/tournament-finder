package repository

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/IDOMATH/session/memorystore"
	"github.com/IDOMATH/tournament-finder/constants"
	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
	"github.com/IDOMATH/tournament-finder/util"
)

type Repository struct {
	TS      db.TournamentStore
	US      db.UserStore
	Session *memorystore.MemoryStore[string]
	Logger  *util.Logger
}

func (repo *Repository) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser types.LoginFormUser
	json.NewDecoder(r.Body).Decode(&loginUser)

	id, err := repo.US.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		repo.Logger.LogError("HandleLogin", "could not log in", err)
		return
	}

	token := util.MakeToken(id)

	repo.Session.Insert(token, strconv.Itoa(id), time.Now().Add(time.Hour))
	w.Header().Set(constants.AuthToken, token)
}

func (repo *Repository) HandleLogout(w http.ResponseWriter, r *http.Request) {
	token := w.Header().Get(constants.AuthToken)
	repo.Session.Delete(token)
	w.WriteHeader(http.StatusNoContent)
}
