package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTomlFileParsing(t *testing.T) {
	t.Run("Valid Toml Syntax", func(t *testing.T) {
		configFileContent := `
			model = "gemini"
			profile = "user"

			[profiles.user]
			instructions = ["first command"]
			responseFormat = "markdown"
		`
		got, err := parseTomlFile([]byte(configFileContent))
		want := &AppConfig{
			Model:   "gemini",
			Profile: "user",
			Profiles: map[string]profiles{
				"user": {
					Instructions:   []string{"first command"},
					ResponseFormat: "markdown",
				},
			},
		}

		assert.Nil(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("Invalid TOML file", func(t *testing.T) {
		configFileContent := `
			model = gemini
			profile = "user"

			[profiles.guest]
			instructions = "first command"
			response= "markdown"
		`
		conf, err := parseTomlFile([]byte(configFileContent))

		assert.Nil(t, conf)
		assert.Error(t, err)

	})
}
