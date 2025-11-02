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
func (c *Config) SetUser(userName string) {
	configStruct := Read()
	configStruct.Current_user_name = userName

	// convert struct to slice of bytes(json)
	data, err := json.Marshal(configStruct)
	if err != nil {
		fmt.Println("Cannot convert config struct to json ( config.SetUser(--userName--) )")
		return
	}

	// send slice of bytes to the local config file
	configPath := getConfigPath()
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Println("Cannot write file to config file ( config.SetUser -> os.WriteFile )", err)
		return
	}
}
