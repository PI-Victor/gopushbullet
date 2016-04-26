package client

import (
	"encoding/json"

	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// Devices has the information about the devices that are attached to your
// account.

// DeviceList is a list of retrieved devices.
type DeviceList struct {
	Devices []Device `json:"devices"`
}

// Device holds information about a single registered device.
type Device struct {
	Active     bool    `json:"active"`
	AppVersion float64 `json:"app_version"`
	Created    float64 `json:"created"`
	ID         string  `json:"iden"`
	Vendor     string  `json:"manufacturer"`
	Model      string  `json:"model"`
	Modified   float64 `json:"modified"`
	Nickname   string  `json:"nickname"`
	Token      string  `json:"push_token"`
}

// IDEA: the authetication validation function needs to be extracted to a
// general wrapper for all outgoing requests. Same for pkg/client/pushes.go

// ListDevices list devices associated to the account.
func ListDevices() {
	userDetails := newUserDetails()
	newConfig := NewConfig()
	configDetails, err := newConfig.ReadConfig()
	if err != nil {
		log.Fatal("", err)
	}
	err = json.Unmarshal(configDetails, &userDetails)
	if err != nil {
		log.Fatal("", err)
	}

	headerRequestOpt := make(map[string]string)
	response, err := util.ProcessAPIRequest("GET", util.DevicesAPIURL, userDetails.Token, headerRequestOpt)
	if err != nil {
		log.Fatal("", err)
	}

	deviceList := &DeviceList{}
	err = json.Unmarshal(response, &deviceList)
	if err != nil {
		log.Fatal("", err)
	}

	for _, device := range deviceList.Devices {
		log.Info("The devices: %#v", device)
	}

}
