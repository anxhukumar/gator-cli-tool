package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

// helper function to get config path
func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error reading home directory (os.UserHomeDir)", err)
		return ""
	}

	configPath := filepath.Join(homeDir, configFileName)
	return configPath
}
