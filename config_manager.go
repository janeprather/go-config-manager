package configManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadConfig will read configuration file specified
func LoadConfig(config interface{}, fileName string) error {
	errPattern := "LoadConfig: %s"
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf(errPattern, err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return fmt.Errorf(errPattern, err.Error())
	}
	return nil
}

// SaveNewConfig will save the config to filename specified only if it doesn't exist
func SaveNewConfig(config interface{}, fileName string) error {
	errPattern := "SaveNewConfig: %s"
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return fmt.Errorf(errPattern, fileName+": file already exists")
	}
	return SaveConfig(config, fileName)
}

// SaveConfig will save config to filename specified
func SaveConfig(config interface{}, fileName string) error {
	errPattern := "SaveConfig: %s"
	jsonBytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return fmt.Errorf(errPattern, err.Error())
	}
	err = ioutil.WriteFile(fileName, jsonBytes, 0600)
	if err != nil {
		return fmt.Errorf(errPattern, err.Error())
	}
	return nil
}
