package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:     "chamber",
		Short:   "Chamber is a CLI tool for generating boilerplate code",
		Version: "0.0.1",
	}
)

type CommandArgs struct {
	ConfigFilePath string
	Verbose        bool
}

var config CommandArgs

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	home, _ := os.UserHomeDir()

	rootCmd.PersistentFlags().StringVarP(&config.ConfigFilePath, "config", "c", path.Join(home, ".chambercfg"), "Path to Config File")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose")
}

func initConfig() {

}
