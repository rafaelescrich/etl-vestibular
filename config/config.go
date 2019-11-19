package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Cfg Holds all config information
var Cfg *Config

// Database holds all information about the database connection
type database struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

// Config has all information about configuration
type Config struct {
	Database database
}

// Load bootstrap configuration
func Load() error {

	// start config with default path
	binary, err := os.Executable()
	if err != nil {
		return err
	}
	configDirectory := filepath.Dir(binary) + "/config"

	if _, err := toml.DecodeFile(configDirectory+"/config.toml", &Cfg); err != nil {
		return err
	}

	return err
}
