package cmd

import (
	"fmt"
	"os"

	"github.com/ralphit/ralphit/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ralphit",
	Short: "Manage iterative AI coding sessions with fresh context",
	Long: `ralphit implements the "Ralph Wiggum" technique for iteratively building
software across multiple AI coding assistant sessions (Claude Code, Codex,
OpenCode, etc.) with fresh context.

It helps developers maintain continuity and structure when working with
AI assistants across different sessions.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ralphit.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		config.SetDefaults()
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
