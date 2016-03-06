package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	transportProtocol = "https://"
	pushbulletAPIURL  = "api.pushbullet.com"
	apiVersion        = "v2"
	usersAuthPath     = "users/me"
)

// UserDetails holds information returned from the API
type UserDetails struct {
	created string
	email   string
}

//{"error":{"code":"invalid_access_token","type":"invalid_request","message":"Access token is missing or invalid.","cat":"(=^･ω･^)y＝"},"error_code":"invalid_access_token"}

// APIRequestError holds an error returned from the API
type APIRequestError struct {
	error     apiErrorCode
	errorCode string
}

type apiErrorCode struct {
	errorCode string
	erroType  string
	message   string
	cat       string
}

// PushbulletAPIURL is the current version of the API URL
var PushbulletAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, usersAuthPath)

// HeaderMIMEJsonType is the MIME Type
const HeaderMIMEJsonType = "application/json"

// Authenticate validates the user Access Token
func Authenticate(userToken string) {
	err := validateUserToken(userToken)
	if err != nil {
		fmt.Println(err)
	}
}

// createNewHttpClient creates a new request client
func createNewHTTPClient() *http.Client {
	return &http.Client{}
}

// validateUserToken validates the current access token
func validateUserToken(userToken string) error {
	requestClient := createNewHTTPClient()
	req, err := http.NewRequest("GET", PushbulletAPIURL, nil)
	req.Header.Add("application/type", HeaderMIMEJsonType)
	req.Header.Add("Access-Token", userToken)

	if err != nil {
		return err
	}

	resp, err := requestClient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	errorCode := &apiErrorCode{}
	err = json.Unmarshal(respBody, errorCode)
	if err != nil {
		return err
	}
	fmt.Println(errorCode)
	return nil
}

// stores the user token in a temporary folder hidden folder in $HOME
func storeUserToken() {}
