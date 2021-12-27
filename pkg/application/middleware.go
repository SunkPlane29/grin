package application

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type (
	JWKS struct {
		Keys []JsonWebKey `json:"keys"`
	}

	JsonWebKey struct {
		KTY string   `json:"kty"`
		KID string   `json:"kid"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5C []string `json:"x5c"`
	}
)

type UserClaims struct {
	jwt.StandardClaims
}

type UserIDKey string

const UIDK UserIDKey = "user-id"

func GetAuth0Cert(domain string, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	if err != nil {
		return "", err
	}

	resp, err := http.Get("https://" + domain + "/.well-known/jwks.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jwks JWKS
	if err = json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return "", err
	}

	var cert string
	for _, key := range jwks.Keys {
		if key.KID == token.Header["kid"] {
			cert = "-----BEGIN CERTIFICATE-----\n" + key.X5C[0] + "\n-----END CERTIFICATE-----"
			break
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key")
	}

	return cert, nil
}

func Auth0Middleware(domain string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			w.Write([]byte("unauthorized access, missing access token"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(accessToken, "Bearer ") {
			w.Write([]byte("no bearer token"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(accessToken, " ")[1]

		cert, err := GetAuth0Cert(domain, tokenString)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cert), nil
		})

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(UserClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), UIDK, claims.Subject)
			f(w, r.WithContext(ctx))
		}
	}
}
