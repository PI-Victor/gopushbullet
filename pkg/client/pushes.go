package client

import (
	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// PushURL is the structure for pushes
type PushURL struct {
	Active        bool    `json:"active"`
	Body          string  `json:"body"`
	Created       float64 `json:"created"`
	Direction     string  `json:"direction"`
	Dismissed     string  `json:"dismissed"`
	Ident         string  `json:"ident"`
	Modified      string  `json:"modified"`
	ReceiverEmail string  `json:"receiver_email_normalized"`
	ReceiverID    string  `json:"receiver_iden"`
	SenderEmail   string  `json:"sender_email_normalized"`
	SenderID      string  `json:"sender_iden"`
	SenderName    string  `json:"sender_name"`
	Title         string  `json:"title"`
	Type          string  `json:"type"`
}

// ListPushes retrieves pushes from an account and stores them on disk
func ListPushes() {
	newUserDetails := UserDetails{}
	newConfig := NewConfig()
	newConfig.ReadConfig(newUserDetails)
	log.Debug("printing", newUserDetails)

	headerRequestOpt := map[string]string{
		"deleted": "off",
	}

	util.ProcessAPIRequest("GET", util.PushesAPIURL, newUserDetails.Token, headerRequestOpt)
}
