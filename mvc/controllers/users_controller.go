package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kalfian/Go-Micro/mvc/utils"

	"github.com/kalfian/Go-Micro/mvc/services"
)

// GetUsers ...
func GetUsers(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(apiErr.Message))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}
