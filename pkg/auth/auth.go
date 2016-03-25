package auth

import (
	"encoding/json"
	"fmt"

	"github.com/PI-Victor/gopushbullet/pkg/client"
	"github.com/PI-Victor/gopushbullet/pkg/util"
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

// Authenticate validates the user Access Token
func Authenticate(userToken string) {
	err := validateUserToken(userToken)
	if err != nil {
		fmt.Println(err)
	}
}

// validateUserToken validates the current access token
func validateUserToken(userToken string) error {
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
	fmt.Printf("Token validated! Logged in as: %s \n", user.Name)

	if err := storeUserToken(user); err != nil {
		return err
	}

	return nil
}

// stores the user token in a temporary hidden folder in $HOME
func storeUserToken(user UserDetails) error {
	newConfig := client.NewConfig()
	err := newConfig.WriteConfig(user)
	if err != nil {
		return err
	}
	return nil
}
