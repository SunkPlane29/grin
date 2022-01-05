package util

import "net/http"

func PostMethodPreflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	w.WriteHeader(http.StatusOK)
}

func GetMethodPreflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	w.WriteHeader(http.StatusOK)
}
