package client

import (
	"encoding/json"

	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// UserDetails holds information returned from the API
type UserDetails struct {
	Active        bool    `json:"active"`
	ID            string  `json:"iden"`
	DateCreated   float64 `json:"created"`
	DateModified  float64 `json:"modified"`
	Email         string  `json:"email_normalized"`
	Name          string  `json:"name"`
	MaxUploadSize float64 `json:"max_upload_size"`
	Token         string  `json:"token"`
}

func newUserDetails() *UserDetails {
	return &UserDetails{}
}

// Authenticate validates the user Access Token
func Authenticate(userToken string) {
	user, err := validateUserToken(userToken)
	if err != nil {
		log.Fatal("", err)
	}
	log.Info("Token validated! Welcome %s", user.Name)
}

// validateUserToken validates the current access token
func validateUserToken(userToken string) (*UserDetails, error) {
	// replace this with the header wrapper
	headerOptions := make(map[string]string)
	apiResponse, err := util.ProcessAPIRequest("GET", util.UsersAPIURL, userToken, headerOptions)
	if err != nil {
		return nil, err
	}

	user := newUserDetails()
	err = json.Unmarshal(apiResponse, &user)
	if err != nil {
		return nil, err
	}
	user.Token = userToken

	if err := storeUserToken(user); err != nil {
		return nil, err
	}

	return user, nil
}

// stores the user token in a temporary hidden folder in $HOME
func storeUserToken(user *UserDetails) error {
	newConfig := NewConfig()
	err := newConfig.WriteConfig(user)
	if err != nil {
		return err
	}
	return nil
}
