package application

import (
	"encoding/json"
	"net/http"
)

type UserData struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (as *AuthServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData UserData

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err := as.AuthService.CreateUser(r.Context(), userData.Username, []byte(userData.Password)) //there is no encryption, this should rely on https hehe
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (as *AuthServer) AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData UserData

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	accessToken, refreshToken, err := as.AuthService.AuthenticateUser(r.Context(), userData.Username, []byte(userData.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	resp := TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (as *AuthServer) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing refresh token in url params"))
		return
	}

	accessToken, refreshToken, err := as.AuthService.RefreshToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	resp := TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
