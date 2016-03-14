package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// APIRequestError holds an error returned from the API
type APIRequestError struct {
	APIError  apiErrorCode `json:"error"`
	ErrorCode string       `json:"error_code"`
}

type apiErrorCode struct {
	ErrorCode string `json:"code"`
	ErrorType string `json:"type"`
	Message   string `json:"message"`
	Cat       string `json:"cat"`
}

const (
	// HeaderMIMEJsonType is the MIME Type to be used in requests to the API
	HeaderMIMEJsonType = "application/json"
	transportProtocol  = "https://"
	pushbulletAPIURL   = "api.pushbullet.com"
	apiVersion         = "v2"
	usersAuthPath      = "users/me"
)

// PushbulletAPIURL is the current version of the API URL
var usersPushbulletAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, usersAuthPath)

// adds parameters to the header of a request that gets sent to the API
func headerEnrichment(req *http.Request, token string) {
	req.Header.Add("application/type", HeaderMIMEJsonType)
	req.Header.Add("Access-Token", token)
}

// CreateNewHTTPClient creates a new request client
func createNewHTTPClient() *http.Client {
	return &http.Client{}
}

// CreateNewRequest returns a new request to the with the specified parameters
func createNewRequest(HTTPMethod string, userToken string) (*http.Request, error) {
	req, err := http.NewRequest(HTTPMethod, usersPushbulletAPIURL, nil)
	if err != nil {
		return nil, err
	}

	headerEnrichment(req, userToken)

	return req, nil
}

// ProcessAPIRequest process requests made to the API by adding header
// enrichment, http operations, etc.
func ProcessAPIRequest(HTTPMethod string, userToken string) ([]byte, error) {
	requestClient := createNewHTTPClient()
	req, err := createNewRequest(HTTPMethod, userToken)
	if err != nil {
		return nil, err
	}

	resp, err := requestClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		var requestError APIRequestError
		err := json.Unmarshal(response, &requestError)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("There was an error authenticating you: %s", requestError.APIError.ErrorCode)
	}

	return response, nil
}
