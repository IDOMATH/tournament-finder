package repository

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"time"

	"github.com/IDOMATH/session/memorystore"
	"github.com/IDOMATH/tournament-finder/types"
)

type Repository struct {
	TH      TournamentHandler
	UH      UserHandler
	Session *memorystore.MemoryStore
}

func (repo *Repository) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser types.LoginFormUser
	json.NewDecoder(r.Body).Decode(&loginUser)

	err := repo.UH.UserStore.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	stringToHash := loginUser.Email + time.Now().String()
	token := sha256.Sum256([]byte(stringToHash))
	repo.Session.Insert(string(token[:]), []byte(loginUser.Email), time.Now().Add(time.Hour))
}
