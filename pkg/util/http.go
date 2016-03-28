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
	// HeaderMIMEJsonType is the MIME Type to be used in requests to the
	// API
	HeaderMIMEJsonType = "application/json"
	transportProtocol  = "https://"
	pushbulletAPIURL   = "api.pushbullet.com"
	apiVersion         = "v2"
	usersAuthPath      = "users/me"
	pushesPath         = "pushes"
	devicesPath        = "devices"
)

// UsersAPIURL is the current version of the API URL
var UsersAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, usersAuthPath)

// PushesAPIURL is the path that returns the active and non active pushes from
// an account
var PushesAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, pushesPath)

// DevicesAPIURL is the path that returns the devices attached to your account.
var DevicesAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, devicesPath)

// Client is a client that does http request to the API and returns a response
// or an error
type Client interface {
	Do(r *http.Request) (*http.Request, error)
}

// adds parameters to the header of a request that gets sent to the API. The
// type of the response and the saved access token are added by default to each
// request
func headerEnrichment(req *http.Request, token string, headerOpt map[string]string) {
	req.Header.Add("application/type", HeaderMIMEJsonType)
	req.Header.Add("Access-Token", token)
	for headerOpt, value := range headerOpt {
		req.Header.Add(headerOpt, value)
	}
}

// CreateNewHTTPClient creates a new request client
func createNewHTTPClient() *http.Client {
	return &http.Client{}
}

// ProcessAPIRequest process requests made to the API by adding header
// enrichment, http operations, etc.
func ProcessAPIRequest(HTTPMethod string, URLPath string, userToken string, headerOpt map[string]string) ([]byte, error) {
	requestClient := createNewHTTPClient()
	req, err := http.NewRequest(HTTPMethod, URLPath, nil)
	if err != nil {
		return nil, err
	}

	headerEnrichment(req, userToken, headerOpt)
	resp, err := requestClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		var requestError APIRequestError
		err := json.Unmarshal(response, &requestError)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("There was an error authenticating you: %s", requestError.APIError.ErrorCode)
	}

	return response, nil
}
