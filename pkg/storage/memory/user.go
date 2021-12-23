package memory

import (
	"errors"

	"github.com/SunkPlane29/solitude/pkg/user"
)

type UserStorage struct {
	d map[string]userStub
}

type userStub struct {
	_user       *user.User
	subscribers map[string]string
}

func (us *UserStorage) CreateUser(u user.User) (*user.User, error) {
	us.d[u.ID] = userStub{_user: &u, subscribers: make(map[string]string)}
	return &u, nil
}

func (us *UserStorage) GetUser(userID string) (*user.User, error) {
	u := us.d[userID]._user
	if u == nil {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (us *UserStorage) ChangeAlias(userID string, newAlias string) (*user.User, error) {
	u := us.d[userID]._user
	u.Alias = newAlias

	return u, nil
}

func (us *UserStorage) ChangeUsername(userID string, newUsername string) (*user.User, error) {
	u := us.d[userID]._user
	u.Username = newUsername

	return u, nil
}

func (us *UserStorage) DeleteUser(userID string) error {
	delete(us.d, userID)
	return nil
}

func (us *UserStorage) AddSubscriber(userID string, subscriberID string) error {
	u := us.d[userID]._user
	u.Subscribers = append(u.Subscribers, userID)
	us.d[userID].subscribers[subscriberID] = subscriberID

	return nil
}

func (us *UserStorage) GetSubscribed(userID string) (*[]user.User, error) {
	subscribed := []user.User{}

	for _, u := range us.d {
		if u._user.ID == userID {
			continue
		}
		if u.subscribers[userID] != "" {
			subscribed = append(subscribed, *u._user)
		}
	}

	return &subscribed, nil
}

func (us *UserStorage) RemoveSubscriber(userID string, subscriberID string) error {
	u := us.d[userID]._user
	u.Subscribers = removeSub(subscriberID, u.Subscribers)
	delete(us.d[userID].subscribers, subscriberID)

	return nil
}

func removeSub(subID string, subs []string) []string {
	var index int
	for i, v := range subs {
		if v == subID {
			index = i
		}
	}

	subs[index] = subs[len(subs)-1]
	return subs[:len(subs)-1]
}
