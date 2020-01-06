package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kalfian/Go-Micro/mvc/utils"
)

type userDaoInterface interface {
	GetUser(userID int64) (*User, *utils.ApplicationError)
}

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Fede", LastName: "Leon", Email: "testemail@gmail.com"},
	}

	//UserDao ...
	UserDao userDaoInterface
)

type userDao struct{}

func init() {
	UserDao = &userDao{}
}

//GetUser ...
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	//Access DB
	log.Println("accesing DB")

	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
