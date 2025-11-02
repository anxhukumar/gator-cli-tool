package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Reads the .gatorconfig json file from the home directory and returns the
// output as a Config struct.
func Read() Config {
	configPath := getConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error in reading file (os.ReadFile)", err)
		return Config{}
	}

	var res Config
	if err := json.Unmarshal(data, &res); err != nil {
		fmt.Println("Error decoding config json", err)
		return Config{}
	}

	return res
}
