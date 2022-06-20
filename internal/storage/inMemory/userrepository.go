package inMemory

import (
	"errors"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/storage"
)

type UsrRepository struct {
	users map[int]*storage.UsersModel
}

func (u UsrRepository) Set(s *storage.UsersModel) error {
	fmt.Printf("добавляем пользователя %v \n", s.Login)
	s.ID = len(u.users) + 1
	u.users[s.ID] = s

	return nil
}

func (u UsrRepository) Get(login string) (*storage.UsersModel, error) {

	for _, u := range u.users {
		if u.Login == login {
			return u, nil
		}
	}
	return nil, errors.New("не найден пользователь " + login)
}
