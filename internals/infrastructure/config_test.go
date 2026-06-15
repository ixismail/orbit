package infrastructure

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// setupMockEnv safely tricks os.UserConfigDir() into using a temporary test folder
func setupMockEnv(t *testing.T) string {
	// Create a safe, temporary folder that gets auto-deleted after the test
	tempDir := t.TempDir()

	// Mock the environment variable based on the OS
	switch runtime.GOOS {
	case "windows":
		t.Setenv("AppData", tempDir)
	case "darwin": // Mac
		t.Setenv("HOME", tempDir) // Mac uses $HOME/Library/Application Support
	default: // Linux
		t.Setenv("XDG_CONFIG_HOME", tempDir)
	}

	return tempDir
}

func TestConfigureOrbit(t *testing.T) {
	// Setup isolated environment
	tempDir := setupMockEnv(t)

	// Run the function we want to test
	ConfigureOrbit()

	// Verify the file was actually created in our temp directory
	var expectedConfigPath string
	if runtime.GOOS == "darwin" {
		expectedConfigPath = filepath.Join(tempDir, "Library", "Application Support", "orbit", "config.yaml")
	} else {
		expectedConfigPath = filepath.Join(tempDir, "orbit", "config.yaml")
	}

	if _, err := os.Stat(expectedConfigPath); os.IsNotExist(err) {
		t.Errorf("Expected config file to be created at %s, but it was not found", expectedConfigPath)
	}
}

func TestLoadConfig(t *testing.T) {
	// Setup isolated environment
	setupMockEnv(t)

	// We must run ConfigureOrbit first to ensure the file exists to be read!
	ConfigureOrbit()

	// Run the function we want to test
	config := LoadConfig()

	// Assertions
	if config == nil {
		t.Fatal("Expected config to be loaded, but got nil")
	}

	// Verify that the default global settings were loaded properly
	if config.Global.CorsBypass != true {
		t.Errorf("Expected default CorsBypass to be true, got %v", config.Global.CorsBypass)
	}

	// Verify that the default routes were loaded
	if _, exists := config.Routes["web.app.loc"]; !exists {
		t.Error("Expected default route 'web.app.loc' to be loaded into the map")
	}
}
