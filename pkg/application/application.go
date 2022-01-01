package application

import (
	"net/http"

	"github.com/SunkPlane29/grin/pkg/post"
	"github.com/SunkPlane29/grin/pkg/user"
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
	g.r.HandleFunc(CreateUserEndpoint, CORSMiddleware(g.postMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreateUserEndpoint, CORSMiddleware(Auth0Middleware(DOMAIN, g.CreateUserHandler))).Methods("POST")
	g.r.HandleFunc(CheckUserExistsEndpoint, CORSMiddleware(g.CheckUserExistsHandler)).Methods("GET")

	g.r.HandleFunc(CreateUserEndpoint, CORSMiddleware(g.postMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreatePostEndpoint, CORSMiddleware(Auth0Middleware(DOMAIN, g.CreatePostHandler))).Methods("POST")
	g.r.HandleFunc(GetPostsEndpoint, CORSMiddleware(g.GetPostsHandler)).Methods("GET")
	g.r.HandleFunc(GetPostEndpoint, CORSMiddleware(g.GetPostHandler)).Methods("GET")
}
