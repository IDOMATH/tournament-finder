package repository

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("error converting id"))
	}

	w.Write([]byte(fmt.Sprintf("getting user with id: %d", id)))
}
