package application

import (
	"encoding/json"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/post"
)

func (g *GrinAPI) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UIDK).(string)
	if !ok {
		w.Write([]byte("no access token")) //TODO: repeated code here, make variable
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()

	var postObj post.Post
	if err := json.NewDecoder(r.Body).Decode(&postObj); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postObj.CreatorID = userID

	createdPost, err := g.postService.CreatePost(userID, postObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}
