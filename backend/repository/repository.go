package repository

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/IDOMATH/CheetahMath/formulas"
	"github.com/IDOMATH/session/memorystore"
	"github.com/IDOMATH/tournament-finder/constants"
	"github.com/IDOMATH/tournament-finder/db"
	"github.com/IDOMATH/tournament-finder/types"
	"github.com/IDOMATH/tournament-finder/util"
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

	// var str strings.Builder
	// str.WriteString(strconv.Itoa(id))
	// str.WriteString(strconv.Itoa(int(time.Now().UnixNano())))
	t := int(time.Now().UnixMilli())*formulas.IntPow(10, formulas.GetDigits(id)) + id

	token := util.TenToThirtysix(t)
	// token := sha256.Sum256([]byte(stringToHash))
	// repo.Session.Insert(string(token[:]), []byte(strconv.Itoa(id)), time.Now().Add(time.Hour))
	// w.Header().Set(constants.AuthToken, string(token[:]))
	repo.Session.Insert(token, []byte(strconv.Itoa(id)), time.Now().Add(time.Hour))
	w.Header().Set(constants.AuthToken, token)
}
