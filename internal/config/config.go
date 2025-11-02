package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

// methods
func (c *Config) SetUser(userName string) error {
	configStruct, err := Read()
	if err != nil {
		return err
	}
	configStruct.Current_user_name = userName

	// convert struct to slice of bytes(json)
	data, err := json.Marshal(configStruct)
	if err != nil {
		err := fmt.Errorf("cannot convert config struct to json ( config.setuser(--username--) )")
		return err
	}

	// send slice of bytes to the local config file
	configPath := getConfigPath()
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		err := fmt.Errorf("cannot write file to config file ( config.setuser -> os.writeFile )")
		return err
	}
	return nil
}
