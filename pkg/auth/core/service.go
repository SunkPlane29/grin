package core

import (
	"context"
	"time"

	"github.com/SunkPlane29/grin/pkg/auth/token"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	tokenIssuer token.JWT
	storage     AuthenticationStorage
}

func NewAuthorizationService(s AuthenticationStorage, keys *token.Keys) AuthorizationService {
	return &service{
		tokenIssuer: token.NewJWT(keys.PubKey, keys.PrivateKey),
		storage:     s,
	}
}

func (s *service) CreateUser(ctx context.Context, username string, password []byte) error {
	pwHash, err := hashPw(password)
	if err != nil {
		return err
	}

	user := User{
		ID:           ksuid.New().String(),
		Username:     username,
		PasswordHash: pwHash,
	}

	if err := s.storage.StoreUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func hashPw(pw []byte) ([]byte, error) {
	pwHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return pwHash, nil
}

//TODO: return jwt token and refresh token and then create appication part
func (s *service) AuthenticateUser(ctx context.Context, username string, password []byte) (string, string, error) {
	user, err := s.storage.GetUser(ctx, username)
	if err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, password); err != nil {
		return "", "", err
	}

	return s.generateTokens(user.ID)
}

func (s *service) generateTokens(userID string) (string, string, error) {
	accessToken, err := s.tokenIssuer.Create(time.Minute*60, userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenIssuer.Create(time.Hour*24, "")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) RefreshToken(refreshToken string) (string, string, error) {
	id, err := s.validateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	return s.generateTokens(id)
}

func (s *service) validateRefreshToken(token string) (string, error) {
	id, err := s.tokenIssuer.Validate(token)
	if err != nil {
		return "", err
	}

	return id, nil
}
