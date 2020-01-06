package services

import (
	"net/http"
	"testing"

	"github.com/kalfian/Go-Micro/mvc/domain"
	"github.com/kalfian/Go-Micro/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMoc      usersDaoMoc
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMoc{}
}

type usersDaoMoc struct{}

func (n *usersDaoMoc) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
			Code:       "not_found",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}
