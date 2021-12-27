package user

type service struct {
	s Storage
}

func New(s Storage) Service {
	return &service{s: s}
}

func (s *service) CreateUser(userID string, user User) (*User, error) {
	return nil, nil
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
