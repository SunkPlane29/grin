package application

import (
	"encoding/json"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/service/user"
)

func (g *GrinAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userObj user.User
	if err := json.NewDecoder(r.Body).Decode(&userObj); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser, err := g.userService.CreateUser(userObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (g *GrinAPI) CheckUserExistsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing id in request url params"))
	}

	if g.userService.CheckUserExists(id) {
		w.WriteHeader(http.StatusFound)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
