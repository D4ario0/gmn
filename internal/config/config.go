package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

const (
	CONFIG_DIR     = "gmn/config.toml"
	DEFAULT_CONFIG = `
model = "%s"
profile = "default"

[profiles.default]
instructions = []
format = "markdown"`
)

type (
	profiles struct {
		Instructions    []string
		Response_Format string
	}

	AppConfig struct {
		Model    string
		Profile  string
		Profiles map[string]profiles
	}
)

const (
	GEMINI_2_0_FLASH = "gemini-2.0-flash"
	GEMINI_2_5_FLASH = "gemini-2.5-flash"
	JSON_FORMAT      = "json"
	MARKDOWN_FORMAT  = "markdown"
)

func LoadConfig() (*AppConfig, error) {
	configFile, err := findConfigFile()
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	return parseTomlFile(content)
}

func parseTomlFile(fileContent []byte) (*AppConfig, error) {
	var cfg AppConfig

	if err := toml.Unmarshal(fileContent, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *AppConfig) GetProfileConfig() string {
	formatPrompt := "Answer using %s format"
	selectedProfile := cfg.Profiles[cfg.Profile]

	// Fail silently and use markdown
	// only 2 supported output types
	if selectedProfile.Response_Format != JSON_FORMAT {
		selectedProfile.Response_Format = MARKDOWN_FORMAT
	}

	selectedProfile.Instructions = append(selectedProfile.Instructions, fmt.Sprintf(formatPrompt, selectedProfile.Response_Format))
	return strings.Join(selectedProfile.Instructions, ".")
}

func findConfigFile() (*os.File, error) {
	configDir, err := xdg.SearchConfigFile(CONFIG_DIR)
	if err != nil {
		// Create the parent dir if it does not exists
		configDir, err = xdg.ConfigFile(CONFIG_DIR)
		if err != nil {
			log.Fatalf("%s", err)
		}

		// Create a config file
		f, err := os.Create(configDir)
		if err != nil {
			log.Fatalf("%s", err)
		}

		// Write a DEFAULT CONFIG
		_, err = fmt.Fprintf(f, DEFAULT_CONFIG, GEMINI_2_0_FLASH)
		return f, err
	}
	return os.Open(configDir)
}
