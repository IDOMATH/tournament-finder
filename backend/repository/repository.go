package repository

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/IDOMATH/session/memorystore"
	"github.com/IDOMATH/tournament-finder/constants"
	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
)

type Repository struct {
	TS      db.TournamentStore
	US      db.UserStore
	Session *memorystore.MemoryStore
}

func (repo *Repository) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser types.LoginFormUser
	json.NewDecoder(r.Body).Decode(&loginUser)

	id, err := repo.US.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	stringToHash := strconv.Itoa(id) + time.Now().String()
	token := sha256.Sum256([]byte(stringToHash))
	repo.Session.Insert(string(token[:]), []byte(strconv.Itoa(id)), time.Now().Add(time.Hour))
	w.Header().Set(constants.AuthToken, string(token[:]))
}
