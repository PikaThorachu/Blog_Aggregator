package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath() // Get file path for config file
	if err != nil {
		return Config{}, fmt.Errorf("failed to get config file path: %v", err)
	}

	data, err := os.ReadFile(configFilePath) // Unmarshal config file
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to write config file: %v", err)
	}
	return config, nil // Return instance of Config{}
}

// Non-exported helper functions
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home dir path: %v", err)
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath() // Get file path for config file
	if err != nil {
		return fmt.Errorf("failed to get config file path: %v", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ") // Marshal edited config file
	if err != nil {
		return fmt.Errorf("failed to marshal config data: %v", err)
	}

	err = os.WriteFile(configFilePath, data, 0644) // Write marshalled config file to config file path
	if err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

const configFileName = ".gatorconfig.json"
