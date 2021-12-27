package user

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreateUser(userID string, user User) (*User, error) {
	return s.store.CreateUser(userID, user)
}

func (s *service) UpdateUsername(userID, newUsername string) error {
	return nil
}

func (s *service) UpdateAlias(userID, newAlias string) error {
	return nil
}

func (s *service) DeleteUser(userID string) error {
	return nil
}
