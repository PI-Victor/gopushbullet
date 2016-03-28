package client

import (
	"github.com/PI-Victor/gopushbullet/pkg/log"
	"github.com/PI-Victor/gopushbullet/pkg/util"
)

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

// ListDevices list devices associated to the account
func ListDevices() {}

// PushNote push a message/notification to a specific device or to all
func PushNote() {}

// PushSMS sends an sms message on behalf of a specified device to a specified
// number
func PushSMS() {}
