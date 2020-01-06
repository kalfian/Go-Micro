package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kalfian/Go-Micro/mvc/utils"

	"github.com/kalfian/Go-Micro/mvc/services"
)

// GetUsers ...
func GetUsers(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.Response(c, apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)

	if apiErr != nil {
		utils.Response(c, http.StatusBadRequest, apiErr)
		return
	}

	// return user to client
	utils.Response(c, http.StatusOK, user)

}
