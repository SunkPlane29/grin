package application

import (
	"net/http"

	"github.com/SunkPlane29/grin/pkg/auth/core"
	"github.com/SunkPlane29/grin/pkg/util"
	"github.com/gorilla/mux"
)

const Prefix = "/api/auth"

var (
	CreateUserEndpoint       = Prefix + "/create-user"
	AuthenticateUserEndpoint = Prefix + "/authenticate"
	RefreshTokenEndpoint     = Prefix + "/refresh"
)

type AuthServer struct {
	AuthService core.AuthorizationService
	r           *mux.Router
}

func NewAuthServer(s core.AuthorizationService) *AuthServer {
	authServer := &AuthServer{AuthService: s, r: mux.NewRouter()}
	authServer.HandleRoutes()

	return authServer
}

func (as *AuthServer) HandleRoutes() {
	as.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	as.r.HandleFunc(CreateUserEndpoint, util.CORSMiddleware(as.CreateUserHandler)).Methods("POST")

	as.r.HandleFunc(AuthenticateUserEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	as.r.HandleFunc(AuthenticateUserEndpoint, util.CORSMiddleware(as.AuthenticateUserHandler)).Methods("POST")

	as.r.HandleFunc(RefreshTokenEndpoint, util.CORSMiddleware(util.PostMethodPreflightHandler)).Methods("OPTIONS")
	as.r.HandleFunc(RefreshTokenEndpoint, util.CORSMiddleware(as.RefreshTokenHandler)).Methods("POST")
}

func (as *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	as.r.ServeHTTP(w, r)
}
