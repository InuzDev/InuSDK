/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "InuSDK",
	Short: "A Multiplatform Software Development Kit Installer",
	Long: `InuSDK is a multiplatform Software Development Kit installer

	Made for developers by a developer, this CLI provides the comfort of simplicity but also
	the freedom to configure your SDKs as you like. Providing full detail of each installation.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (Default: $HOME/.InuSDK/config.yaml")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.InuSDK")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
