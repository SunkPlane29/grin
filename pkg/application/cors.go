package application

import (
	"net/http"
)

func (g *GrinAPI) postMethodPreflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	w.WriteHeader(http.StatusOK)
}
