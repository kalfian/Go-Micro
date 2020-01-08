package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kalfian/Go-Micro/src/api/clients/restclient"
	"github.com/kalfian/Go-Micro/src/api/domain/github"
)

const (
	headerAuth       = "Authorization"
	headerAuthFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthHeader(token string) string {
	return fmt.Sprintf(headerAuthFormat, token)
}

//CreateRepo ...
func CreateRepo(token string, req github.CreateRepoRequest) (*github.ResponseCreateRepo, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuth, getAuthHeader(token))

	res, err := restclient.Post(urlCreateRepo, req, headers)

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = res.StatusCode

		return nil, &errResponse
	}

	var result github.ResponseCreateRepo
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal repo successful response %s", err.Error())
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil

}
