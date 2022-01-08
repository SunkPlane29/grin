package application

import (
	"encoding/json"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/service/user"
	"github.com/SunkPlane29/grin/pkg/util"
)

func (g *GrinAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(util.IDK).(string) //util.idontknow?
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("no access token")) //TODO: repeated code here, make variable
		return
	}

	defer r.Body.Close()

	var userObj user.User
	if err := json.NewDecoder(r.Body).Decode(&userObj); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if userObj.Alias == "" || userObj.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty request body"))
		return
	}

	createdUser, err := g.userService.CreateUser(userID, userObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
