package application

import (
	"net/http"

	"github.com/SunkPlane29/grin/pkg/user"
	"github.com/gorilla/mux"
)

const DOMAIN = "sunkplane.us.auth0.com"

const APIPathPrefix = "/api"

var (
	CreateUserEndpoint = APIPathPrefix + "/users"
)

type GrinStorage struct {
	UserStorage user.Storage
}

func NewGrinStorage(userStorage user.Storage) *GrinStorage {
	return &GrinStorage{
		UserStorage: userStorage,
	}
}

type GrinAPI struct {
	r           *mux.Router
	userService user.Service
}

func New(s *GrinStorage) *GrinAPI {
	g := &GrinAPI{r: mux.NewRouter(), userService: user.New(s.UserStorage)}
	g.HandleRoutes()

	return g
}

func (g *GrinAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.r.ServeHTTP(w, r)
}

func (g *GrinAPI) HandleRoutes() {
	g.r.HandleFunc(CreateUserEndpoint, Auth0Middleware(DOMAIN, g.CreateUserHandler)).Methods("POST")
}
