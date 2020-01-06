package services

import (
	"github.com/kalfian/Go-Micro/mvc/domain"
	"github.com/kalfian/Go-Micro/mvc/utils"
)

type userService struct{}

var (
	// UserService ...
	UserService userService
)

//GetUser ...
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
