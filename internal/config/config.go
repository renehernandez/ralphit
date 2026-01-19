package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetDefaults() {
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(home)
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".ralphit")
	viper.SetEnvPrefix("RALPHIT")
}

func ConfigFilePath() string {
	if viper.ConfigFileUsed() != "" {
		return viper.ConfigFileUsed()
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ".ralphit.yaml"
	}

	return filepath.Join(home, ".ralphit.yaml")
}
