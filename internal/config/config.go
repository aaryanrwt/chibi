package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig initializes Viper for configuration management.
func InitConfig(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".chibi" (without extension).
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".chibi")
	}

	viper.SetEnvPrefix("CHIBI")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	// We don't return an error if config file is not found, as flags/env are sufficient
	return nil
}
