package core

type AuthenticationStorage interface {
	StoreUser(*User) error
	GetUser(string) (*User, error)
}

type AuthorizationService interface {
	CreateUser(username string, password []byte) error
	AuthenticateUser(username string, password []byte) (accessToken, refreshToken string, err error)
	RefreshToken(token string) (accessToken, refreshToken string, err error)
}
