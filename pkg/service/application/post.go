package application

import (
	"encoding/json"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/service/post"
	"github.com/SunkPlane29/grin/pkg/util"
	"github.com/gorilla/mux"
)

func (g *GrinAPI) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(util.IDK).(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("no access token")) //TODO: repeated code here, make variable
		return
	}

	defer r.Body.Close()

	var postObj post.Post
	if err := json.NewDecoder(r.Body).Decode(&postObj); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	postObj.CreatorID = userID

	createdPost, err := g.postService.CreatePost(userID, postObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (g *GrinAPI) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	if v["creator-id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no userID given"))
		return
	}

	posts, err := g.postService.GetPosts(v["creator-id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (g *GrinAPI) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	if v["creator-id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no userID given"))
		return
	} else if v["post-id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no postID given"))
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

func (g *GrinAPI) GetPostsSubscribedHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(util.IDK).(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("no access token")) //TODO: repeated code here, make variable
		return
	}

	posts, err := g.postService.GetPosts(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
