package infrastructure

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/ixismail/orbit/internals/theme"
)

// Embedding the oribit default config file into the binary
//go:embed default_orbit_config.yaml
var DefaultConfigurations []byte

// ConfigureOrbit sets up the configuration directory and file with the
// default configurations for Orbit if the directory or file is not present.
func ConfigureOrbit() {
	// Find the config directory based on the Operating System
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf(theme.Error.Render("Error finding config directory: %v\n"), err)
	}

	// Orbit config directory and config file path
	configDirPath := filepath.Join(configDir, "orbit")
	configFilePath := filepath.Join(configDirPath, "config.yaml")

	// Check if the config file already exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {

		// Config dir doesn't exist, create it with default config file
		if err := os.MkdirAll(configDirPath, 0755); err != nil {
			log.Fatalf(theme.Error.Render("Failed to create orbit configuration directory: %v\n"), err)
		}

		if err := os.WriteFile(configFilePath, DefaultConfigurations, 0644); err != nil {
			log.Fatalf(theme.Error.Render("Failed to write default config: %v\n"), err)
		}
	}
}