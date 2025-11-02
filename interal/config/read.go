package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Reads the .gatorconfig json file from the home directory and returns the
// output as a Config struct.
func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error reading home directory (os.UserHomeDir)", err)
		return Config{}
	}

	configPath := filepath.Join(homeDir, ".gatorconfig.json")

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
