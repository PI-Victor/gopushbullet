package client

import (
	"github.com/PI-Victor/gunner/pkg/log"
	"os"
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

// NewUser is a helper function that returns a new empty user instance
func NewUser() *User {
	return &User{}
}

// Authenticate tries to validate the user
func (u *User) Authenticate() {
	if err := u.validateUserToken(); err != nil {
		log.Fatal("An error occured while trying to validate specified token", err)
		os.Exit(0)
	}
}

func (u *User) validateUserToken() error {
	return nil
}

func (u *User) storeUserToken() error {
	return nil
}
