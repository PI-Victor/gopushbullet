package client

import (
	"encoding/json"

	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// PushesList contains the data associated with all the retrieved pushes
type PushesList struct {
	Pushes []PushURL `json:"pushes"`
	Cursor string    `json:"cursor"`
}

// PushURL is the structure for pushes
type PushURL struct {
	Active        bool    `json:"active"`
	Body          string  `json:"body"`
	URL           string  `json:"url"`
	Created       float64 `json:"created"`
	Direction     string  `json:"direction"`
	Dismissed     bool    `json:"dismissed"`
	Ident         string  `json:"iden"`
	Modified      float64 `json:"modified"`
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
	userDetails := UserDetails{}
	newConfig := NewConfig()
	configDetails, err := newConfig.ReadConfig()
	if err != nil {
		log.Fatal("", err)
	}
	err = json.Unmarshal(configDetails, &userDetails)
	if err != nil {
		log.Fatal("", err)
	}
	// create a bogus request
	headerRequestOpt := make(map[string]string)

	response, err := util.ProcessAPIRequest("GET", util.PushesAPIURL, userDetails.Token, headerRequestOpt)

	if err != nil {
		log.Fatal("", err)
	}

	pushes := &PushesList{}
	err = json.Unmarshal(response, &pushes)
	if err != nil {
		log.Fatal("", err)
	}

	if pushes.Cursor != "" {
		log.Debug("Cursor is not empty", pushes.Cursor)
	}

	for _, i := range pushes.Pushes {
		if i.Type == "" {
			log.Debug("This is not an URL: %#v", i)
			continue
		}
		log.Info("%+v", i.URL)
	}
}

func pushPrinter() {

}
