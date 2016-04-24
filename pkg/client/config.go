package client

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PI-Victor/gunner/pkg/log"
	"github.com/PI-Victor/gunner/pkg/util"
)

// Configuration holds information about the current setup of this CLI
// application
type Configuration struct {
	// full path to the config directory
	configDir string
	// full path to the config file
	configFile string
}

// NewConfig returns a new instance of the client config
func NewConfig() *Configuration {
	confDir, confFile, err := util.CreateDirectories()
	if err != nil {
		log.Fatal("", err)
	}

	return &Configuration{
		configDir:  confDir,
		configFile: confFile,
	}
}

// WriteConfig flushes the jsonified data to the config file
func (c *Configuration) WriteConfig(user interface{}) error {
	fileHandler, err := os.Create(c.configFile)
	if err != nil {
		return err
	}
	defer fileHandler.Close()

	// prettify the encoding so that it's human readable
	encodedUserDetails, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return err
	}

	_, err = fileHandler.Write(encodedUserDetails)
	if err != nil {
		return err
	}
	return nil
}

// ReadConfig returns the user details stored on disk in json encoding
func (c *Configuration) ReadConfig() (configDetails []byte, err error) {
	if _, err = os.Stat(c.configFile); os.IsNotExist(err) {
		return nil, err
	}

	configDetails, err = ioutil.ReadFile(c.configFile)
	if err != nil {
		return nil, err
	}

	return configDetails, nil
}

// Logout logs a user out by purging all the stored config files and data from
// disk
func (c *Configuration) Logout() {
	err := util.PurgeArtifacts(c.configDir)
	if err != nil {
		log.Fatal("An error occured while logging out: ", err)
	}

	log.Info("Your user details have been successfully removed")
}
