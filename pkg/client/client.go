package client

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	configFileName = "gopush.json"
	configDir      = ".gopush"
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
	confDir, confFile, err := configSetup()
	if err != nil {
		panic(err)
	}

	return &Configuration{
		configDir:  confDir,
		configFile: confFile,
	}
}

// WriteConfig flushes the jsonified data to the config file
func (c *Configuration) WriteConfig(user interface{}) error {
	fileHandler, err := os.Create(c.configFile)
	defer fileHandler.Close()
	if err != nil {
		fmt.Println("i failed", err)
		return err
	}

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

// PurgeConfig - purges current config file regardless if it has details about
// a user or not. Flushes details but doesn't delete the file
func (c *Configuration) PurgeConfig() {
	var emptyBytes []byte

	fileHandler, err := os.Create(c.configFile)
	if err != nil {
		panic(err)
	}

	defer fileHandler.Close()
	fileHandler.Write(emptyBytes)
}

func (c *Configuration) readConfig() interface{} {

	return ""
}

// creates current setup dir and config file
func configSetup() (configDirPath string, configFilePath string, err error) {
	homePath := os.Getenv("HOME")
	if homePath == "" {
		return "", "", fmt.Errorf("Couldn't read the $HOME variable, no place to store app config")
	}

	configDirPath = path.Join(homePath, configDir)
	if _, err := os.Stat(configDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(configDirPath, 0700)
		if err != nil {
			return "", "", err
		}
	}

	configFilePath = path.Join(configDirPath, configFileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFileHandler, err := os.Create(configFilePath)
		defer configFileHandler.Close()
		if err != nil {
			return "", "", err
		}
	}
	return
}
