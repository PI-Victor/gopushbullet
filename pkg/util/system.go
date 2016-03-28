package util

import (
	"fmt"
	"os"
	"path"
)

const (
	configFileName = "user.json"
	configDir      = ".gopush"
)

// CreateDirectories creates the configuration directories where the
// application stores configuration and data
func CreateDirectories() (configDirPath string, configFilePath string, err error) {
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
		if err != nil {
			return "", "", err
		}
		defer configFileHandler.Close()
	}
	return
}

// PurgeArtifacts removes user stored configuration and data from disk
func PurgeArtifacts(configDir string) error {
	err := os.RemoveAll(configDir)
	if err != nil {
		return err
	}
	return nil
}
