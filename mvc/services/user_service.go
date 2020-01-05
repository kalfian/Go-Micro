package services

import (
	"github.com/kalfian/Go-Micro/mvc/domain"
	"github.com/kalfian/Go-Micro/mvc/utils"
)

//GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
