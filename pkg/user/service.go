package user

type service struct {
	s Storage
}

func NewService(s Storage) Service {
	return &service{s: s}
}

func (s *service) CreateUser(u User) (*User, error) {
	return nil, nil
}

func (s *service) GetUser(userID string) (*User, error) {
	return nil, nil
}

func (s *service) ChangeAlias(userID string, newAlias string) (*User, error) {
	return nil, nil
}

func (s *service) ChangeUsername(userID string, newUsername string) (*User, error) {
	return nil, nil
}

func (s *service) DeleteUser(userID string) error {
	return nil
}

//TODO: should we "addSubscriber" or should the user "subscribe"
func (s *service) AddSubscriber(userID string, subscriberID string) error {
	return nil
}

func (s *service) GetSubscribed(userID string) (*[]User, error) {
	return nil, nil
}

func (s *service) RemoveSubscriber(userID string, subscriberID string) error {
	return nil
}
