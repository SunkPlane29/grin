package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	pubKey     []byte
	privateKey []byte
}

func NewJWT(pubKey []byte, privateKey []byte) JWT {
	return JWT{
		pubKey:     pubKey,
		privateKey: privateKey,
	}
}

func (j JWT) Create(ttl time.Duration, sub interface{}) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = sub                 // Our custom data.
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (j JWT) Validate(token string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return "", fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return "", fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", fmt.Errorf("validate: invalid")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("unable to parse claims from token")
	}

	return sub, nil
}

// DEPRECATED
//
// this function should be used to bypass the ValidationErrorIssuedAt, that happens
// when the token is used before issued, this error can happen due to synchronization
// issues
func IssuedAtCheck(err error, token *jwt.Token) error {
	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorIssuedAt {
			token.Valid = true
			return nil
		}
	}

	return err
}
