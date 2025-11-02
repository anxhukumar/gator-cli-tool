package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Reads the .gatorconfig json file from the home directory and returns the
// output as a Config struct.
func Read() (Config, error) {
	configPath := getConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		err := fmt.Errorf("error in reading config file: %w", err)
		return Config{}, err
	}

	var res Config
	if err := json.Unmarshal(data, &res); err != nil {
		err := fmt.Errorf("error decoding config json: %w", err)
		return Config{}, err
	}

	return res, nil
}
