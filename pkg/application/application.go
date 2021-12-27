package application

import (
	"net/http"

	"github.com/SunkPlane29/grim/pkg/user"
	"github.com/gorilla/mux"
)

const APIPathPrefix = "/api"

var (
	CreateUserEndpoint = APIPathPrefix + "/users"
)

type GrinStorage interface {
	user.Storage
}

type GrinAPI struct {
	r           *mux.Router
	userService user.Service
}

func New(s GrinStorage) *GrinAPI {
	g := &GrinAPI{r: mux.NewRouter(), userService: user.New(s)}
	g.HandleRoutes()

	return g
}

func (g *GrinAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.r.ServeHTTP(w, r)
}

func (g *GrinAPI) HandleRoutes() {
	g.r.HandleFunc(CreateUserEndpoint, g.CreateUserHandler).Methods("POST")
}
