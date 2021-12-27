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

type UserIDKey string

const UIDK UserIDKey = "user-id"

func GetAuth0Cert(domain string, tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

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
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized access, missing access token"))
			return
		}

		if !strings.HasPrefix(accessToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no bearer token"))
			return
		}

		tokenString := strings.Split(accessToken, " ")[1]

		cert, err := GetAuth0Cert(domain, tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return "", err
			}
			return key, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub, ok := claims["sub"].(string)
			if sub == "" || !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid sub claim"))
				return
			}

			ctx := context.WithValue(r.Context(), UIDK, strings.TrimRight(sub, "@clients"))
			f(w, r.WithContext(ctx))
		}
	}
}
