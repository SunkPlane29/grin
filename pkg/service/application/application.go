package application

import (
	"net/http"

	"github.com/SunkPlane29/grin/pkg/service/post"
	"github.com/SunkPlane29/grin/pkg/service/user"
	"github.com/SunkPlane29/grin/pkg/util"
	"github.com/gorilla/mux"
)

const DOMAIN = "sunkplane.us.auth0.com"

const APIPathPrefix = "/api"

var (
	CreateUserEndpoint      = APIPathPrefix + "/users"
	CheckUserExistsEndpoint = APIPathPrefix + "/user-exists"
)

var (
	CreatePostEndpoint = APIPathPrefix + "/posts"
	GetPostsEndpoint   = APIPathPrefix + "/users" + "/{creator-id}" + "/posts"
	GetPostEndpoint    = APIPathPrefix + "/users" + "/{creator-id}" + "/posts" + "/{post-id}"
)

type GrinStorage struct {
	UserStorage user.Storage
	PostStorage post.Storage
}

type StorageOptions struct {
	UserStorage user.Storage
	PostStorage post.Storage
}

func NewGrinStorage(storageOpt StorageOptions) *GrinStorage {
	return &GrinStorage{
		UserStorage: storageOpt.UserStorage,
		PostStorage: storageOpt.PostStorage,
	}
}

type GrinAPI struct {
	r           *mux.Router
	userService user.Service
	postService post.Service
}

func New(s *GrinStorage) *GrinAPI {
	g := &GrinAPI{
		r:           mux.NewRouter(),
		userService: user.New(s.UserStorage),
		postService: post.New(s.PostStorage),
	}
	g.HandleRoutes()

	return g
}

func (g *GrinAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.r.ServeHTTP(w, r)
}

func (g *GrinAPI) HandleRoutes() {
	g.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(g.CreateUserHandler)).Methods("POST")
	g.r.HandleFunc(CheckUserExistsEndpoint, util.CORSMiddleware(g.CheckUserExistsHandler)).Methods("GET")

	g.r.HandleFunc(CreatePostEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreatePostEndpoint, util.CORSMiddleware(StubAuthMiddleware(g.CreatePostHandler))).Methods("POST")
	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(util.GetMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(g.GetPostsHandler)).Methods("GET")
	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(util.GetMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(GetPostEndpoint, util.CORSMiddleware(g.GetPostHandler)).Methods("GET")
}
