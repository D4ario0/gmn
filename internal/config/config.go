package config

import (
	"os"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type (
	profile struct {
		instructions   []string
		responseFormat string
	}

	AppConfig struct {
		Model    string
		profile  string
		profiles map[string]profile
	}
)

func LoadConfig() (*AppConfig, error) {
	configFilePath, err := xdg.ConfigFile("gmn/gmn.toml")
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	return parseTomlFile(f)
}

func parseTomlFile(fileContent []byte) (*AppConfig, error) {
	var cfg AppConfig

	if err := toml.Unmarshal(fileContent, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *AppConfig) GetProfileConfig() string {
	return strings.Join(cfg.profiles[cfg.profile].instructions, "")
}

func ValidateConfig() {}
