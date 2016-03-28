package client

import (
	"encoding/json"

	"github.com/PI-Victor/gopushbullet/pkg/log"
	"github.com/PI-Victor/gopushbullet/pkg/util"
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

// Authenticate validates the user Access Token
func Authenticate(userToken string) {
	err := validateUserToken(userToken)
	if err != nil {
		log.Critical(err)
	}
}

// validateUserToken validates the current access token
func validateUserToken(userToken string) error {
	// replace this with the header wrapper
	headerOptions := make(map[string]string)
	apiResponse, err := util.ProcessAPIRequest("GET", util.UsersAPIURL, userToken, headerOptions)
	if err != nil {
		return err
	}

	var user UserDetails
	err = json.Unmarshal(apiResponse, &user)
	if err != nil {
		return err
	}
	user.Token = userToken
	log.Info("Token validated!\nLogged in as: %s \n", user.Name)

	if err := storeUserToken(user); err != nil {
		return err
	}

	return nil
}

// stores the user token in a temporary hidden folder in $HOME
func storeUserToken(user UserDetails) error {
	newConfig := NewConfig()
	err := newConfig.WriteConfig(user)
	if err != nil {
		return err
	}
	return nil
}
