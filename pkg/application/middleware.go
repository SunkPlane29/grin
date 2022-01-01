package application

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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

func IssuedAtCheck(err error, token *jwt.Token) error {
	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorIssuedAt {
			token.Valid = true
			return nil
		}
	}

	return err
}

type UserIDKey string

const UIDK UserIDKey = "user-id"

func GetAuth0Cert(domain string, tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwt.MapClaims{})

	err = IssuedAtCheck(err, token)
	if err != nil {
		return "", err
	}

	resp, err := http.Get("https://" + domain + "/.well-known/jwks.json")
	if err != nil {
		log.New(os.Stdout, "", log.Lshortfile).Println("here")
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

		err = IssuedAtCheck(err, token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			sub, ok := claims["sub"].(string)
			if sub == "" || !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid sub claim"))
				return
			}

			subSplit := strings.Split(sub, "|")
			if len(subSplit) < 2 {
				sub = subSplit[0]
			} else {
				sub = subSplit[1]
			}

			ctx := context.WithValue(r.Context(), UIDK, strings.TrimRight(sub, "@clients"))
			f(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func CORSMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") //TODO: hardcoded origin
		f(w, r)
	}
}

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		data *responseData
	}
)

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lw.ResponseWriter.Write(b)
	lw.data.size += size
	return size, err
}

func (lw *loggingResponseWriter) WriteHeader(statusCode int) {
	lw.ResponseWriter.WriteHeader(statusCode)
	lw.data.status = statusCode
}

func LoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := &loggingResponseWriter{
			ResponseWriter: w,
			data:           responseData,
		}

		h.ServeHTTP(lw, r)

		duration := time.Since(start)

		logger := log.New(os.Stdout, "grin-api | ", log.Ldate|log.Ltime)
		logger.Printf("[%s] %s, status: %d, size: %d, elapsed: %dms\n",
			method,
			uri,
			responseData.status,
			responseData.size,
			duration.Round(time.Millisecond).Milliseconds(),
		)
	})
}

func RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				panic(err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
