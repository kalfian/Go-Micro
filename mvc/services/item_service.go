package services

import (
	"net/http"

	"github.com/kalfian/Go-Micro/mvc/domain"
	"github.com/kalfian/Go-Micro/mvc/utils"
)

type itemService struct{}

var (
	// ItemService ...
	ItemService itemService
)

// GetItem ...
func (s *itemService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "not implemented",
		StatusCode: http.StatusInternalServerError,
	}
}
