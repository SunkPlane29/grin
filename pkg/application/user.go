package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SunkPlane29/grim/pkg/user"
)

func (g *GrinAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UIDK).(string)

	defer r.Body.Close()

	var userObj user.User
	if err := json.NewDecoder(r.Body).Decode(&userObj); err != nil {
		w.Write([]byte("invalid json request body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userID != userObj.ID {
		w.Write([]byte(fmt.Sprintf("IDs not matching: req body: %s, authorization token: %s", userObj.ID, userID)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser, err := g.userService.CreateUser(userID, userObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
	w.WriteHeader(http.StatusCreated)
}
