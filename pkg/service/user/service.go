package user

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreateUser(userID string, user User) (*User, error) {
	user.ID = userID

	return s.store.CreateUser(user)
}

func (s *service) CheckUserExists(userID string) bool {
	user, err := s.store.GetUser(userID)
	if err != nil || user == nil {
		return false
	}

	return true
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
