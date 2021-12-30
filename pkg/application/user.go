package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/user"
)

func (g *GrinAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UIDK).(string)
	if !ok {
		w.Write([]byte("no access token"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()

	var userObj user.User
	if err := json.NewDecoder(r.Body).Decode(&userObj); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userObj.ID = userID
	log.Printf("CreateUser request: %v\n", userObj)

	createdUser, err := g.userService.CreateUser(userID, userObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
