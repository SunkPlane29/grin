package application

import (
	"net/http"

	"github.com/SunkPlane29/grin/pkg/auth/token"
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
	CreatePostEndpoint         = APIPathPrefix + "/posts"
	GetPostsSubscribedEndpoint = APIPathPrefix + "/posts/subscribed"
	GetPostsEndpoint           = APIPathPrefix + "/users" + "/{creator-id}" + "/posts"
	GetPostEndpoint            = APIPathPrefix + "/users" + "/{creator-id}" + "/posts" + "/{post-id}"
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
	tokenIssuer token.JWT
	r           *mux.Router
	userService user.Service
	postService post.Service
}

func New(s *GrinStorage, tokenIssuer token.JWT) *GrinAPI {
	g := &GrinAPI{
		tokenIssuer: tokenIssuer,
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

//FIXME: too many middlewares
func (g *GrinAPI) HandleRoutes() {
	g.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(util.AuthMiddleware(g.tokenIssuer, g.CreateUserHandler))).Methods("POST")

	g.r.HandleFunc(CheckUserExistsEndpoint, util.CORSMiddleware(g.CheckUserExistsHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CheckUserExistsEndpoint, util.CORSMiddleware(g.CheckUserExistsHandler)).Methods("GET")

	g.r.HandleFunc(CreatePostEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(CreatePostEndpoint, util.CORSMiddleware(util.AuthMiddleware(g.tokenIssuer, g.CreatePostHandler))).Methods("POST")

	g.r.HandleFunc(GetPostsSubscribedEndpoint, util.CORSMiddleware(util.GetMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(GetPostsSubscribedEndpoint, util.CORSMiddleware(util.AuthMiddleware(g.tokenIssuer, g.GetPostsSubscribedHandler))).Methods("GET")

	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(util.GetMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(g.GetPostsHandler)).Methods("GET")

	g.r.HandleFunc(GetPostsEndpoint, util.CORSMiddleware(util.GetMethodPreflightHandler)).Methods("OPTIONS")
	g.r.HandleFunc(GetPostEndpoint, util.CORSMiddleware(g.GetPostHandler)).Methods("GET")
}
