package domain

import (
	"fmt"
	"net/http"

	"github.com/kalfian/Go-Micro/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Fede", LastName: "Leon", Email: "testemail@gmail.com"},
	}
)

//GetUser ...
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
