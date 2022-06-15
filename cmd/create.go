package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand()
}

var createCmd = &cobra.Command{
	Use:   "create [template] (options)",
	Short: "Create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
