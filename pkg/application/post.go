package application

import (
	"encoding/json"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/post"
	"github.com/gorilla/mux"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (g *GrinAPI) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	if v["creator-id"] == "" {
		w.Write([]byte("no userID given"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts, err := g.postService.GetPosts(v["creator-id"])
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (g *GrinAPI) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	if v["creator-id"] == "" {
		w.Write([]byte("no userID given"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if v["post-id"] == "" {
		w.Write([]byte("no postID given"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := g.postService.GetPost(v["creator-id"], v["post-id"])
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
