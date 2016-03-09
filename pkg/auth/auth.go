package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	"path"

	"github.com/PI-Victor/gopushbullet/pkg/client"
)

const (
	// HeaderMIMEJsonType is the MIME Type to be used in requests to the API
	HeaderMIMEJsonType  = "application/json"
	transportProtocol   = "https://"
	pushbulletAPIURL    = "api.pushbullet.com"
	apiVersion          = "v2"
	usersAuthPath       = "users/me"
	okResponse          = "200 OK"
	unathorizedResponse = "401 Unauthorized"
)

// UserDetails holds information returned from the API
type UserDetails struct {
	Active        bool    `json:"active"`
	ID            string  `json:"iden"`
	DateCreated   float64 `json:"created"`
	DateModified  float64 `json:"modified"`
	Email         string  `json:"email"`
	Name          string  `json:"name"`
	MaxUploadSize float64 `json:"max_upload_size"`
	Token         string  `json:"token"`
}

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

// PushbulletAPIURL is the current version of the API URL
var PushbulletAPIURL = transportProtocol + path.Join(pushbulletAPIURL, apiVersion, usersAuthPath)

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
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	// this doesn't leave room for other response codes, but it will have to do
	// for now
	if resp.Status == unathorizedResponse {
		var requestError APIRequestError
		if err != nil {
			return err
		}
		err = json.Unmarshal(response, &requestError)
		if err != nil {
			return err
		}
		return fmt.Errorf("There was an error authenticating you: %s", requestError.APIError.ErrorCode)
	}

	var user UserDetails
	err = json.Unmarshal(response, &user)
	if err != nil {
		return err
	}
	user.Token = userToken
	fmt.Printf("Token validated! Logged in as: %+v \n", user)

	storeUserToken(userToken)

	return nil
}

// stores the user token in a temporary folder hidden folder in $HOME
func storeUserToken(userToken string) {
	newConfig := client.NewConfig()
	fmt.Println(newConfig)
}
