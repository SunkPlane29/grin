package core

import (
	"io/ioutil"
	"time"

	"github.com/SunkPlane29/grin/pkg/auth/token"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type Keys struct {
	pubKey     []byte
	privateKey []byte
}

func NewKeysFromCertFiles(pubKeyFName, privateKeyFName string) (*Keys, error) {
	pubKey, err := ioutil.ReadFile(pubKeyFName)
	if err != nil {
		return nil, err
	}

	privateKey, err := ioutil.ReadFile(privateKeyFName)
	if err != nil {
		return nil, err
	}

	return &Keys{
		pubKey:     pubKey,
		privateKey: privateKey,
	}, nil
}

type service struct {
	tokenIssuer token.JWT
	storage     AuthenticationStorage
}

func NewAuthorizationService(s AuthenticationStorage, keys *Keys) AuthorizationService {
	return &service{
		tokenIssuer: token.NewJWT(keys.pubKey, keys.privateKey),
		storage:     s,
	}
}

func (s *service) CreateUser(username string, password []byte) (string, string, error) {
	pwHash, err := hashPw(password)
	if err != nil {
		return "", "", err
	}

	user := &User{
		ID:           ksuid.New().String(),
		Username:     username,
		PasswordHash: pwHash,
	}

	if err := s.storage.StoreUser(user); err != nil {
		return "", "", err
	}

	return s.generateTokens(user.ID)
}

func hashPw(pw []byte) ([]byte, error) {
	pwHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return pwHash, nil
}

//TODO: return jwt token and refresh token and then create appication part
func (s *service) AuthenticateUser(username string, password []byte) (string, string, error) {
	user, err := s.storage.GetUser(username)
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
