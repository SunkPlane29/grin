package user

type service struct {
	s Storage
}

func NewService(s Storage) Service {
	return &service{s: s}
}

func (s *service) CreateUser(user User) {

}

func (s *service) ChangeAlias(userID string, newAlias string) {

}

func (s *service) ChangeUsername(userID string, newUsername string) {

}

//TODO: should we "addSubscriber" or should the user "subscribe"
func (s *service) AddSubscriber(userID string) {

}

func (s *service) RemoveSubscriber(userID string) {

}
