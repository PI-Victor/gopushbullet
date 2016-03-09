package client

import (
	"fmt"
	"os"
	"path"
)

const configFileName = "gopush.json"
const configDir = ".gopush"

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

// ReadConfig reads the configuration file and returns needed data for
// authtentication
func (c Configuration) readConfig() (fileHandler *os.File, err error) {
	fileHandler, err = os.Open(c.configFile)
	if err != nil {
		return nil, err
	}
	return fileHandler, nil
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
