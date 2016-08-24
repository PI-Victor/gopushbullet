package client

import (
	"encoding/json"

	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// User holds information returned from the API
type User struct {
	Active        bool    `json:"active"`
	ID            string  `json:"iden"`
	DateCreated   float64 `json:"created"`
	DateModified  float64 `json:"modified"`
	Email         string  `json:"email_normalized"`
	Name          string  `json:"name"`
	MaxUploadSize float64 `json:"max_upload_size"`
	Token         string  `json:"token"`
}

// NewUser returns a new empty user instance
func NewUser() *User {
	return &User{}
}

// Authenticate validates the user Access Token
func (u *User) Authenticate() {
	err := u.validateUserToken(u.Token)
	if err != nil {
		log.Fatal("", err)
	}
}

// validateUserToken validates the current access token
func (u *User) validateUserToken(userToken string) error {
	// replace this with the header wrapper
	headerOptions := make(map[string]string)
	apiResponse, err := util.ProcessAPIRequest("GET", util.UsersAPIURL, userToken, headerOptions)
	if err != nil {
		return err
	}
	err = json.Unmarshal(apiResponse, &u)
	if err != nil {
		return err
	}
	log.Info("Token validated!\nLogged in as: %s \n", u.Name)
	if err := u.storeUserToken(); err != nil {
		return err
	}

	return nil
}

// stores the user token in a temporary hidden folder in $HOME
func (u *User) storeUserToken() error {
	newConfig := NewConfig()
	err := newConfig.WriteConfig(u)
	if err != nil {
		return err
	}
	return nil
}
