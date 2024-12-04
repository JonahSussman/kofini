package cmd

import (
	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "kofini",
		Short: "kofinin is a CLI for app discovery and transformation.",
		Long:  "kofini is a CLI for app discovery and transformation. It outputs platform configuration files and generates Helm charts for seamless deployment",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(GetGenerateCmd())
	rootCmd.AddCommand(GetDiscoverCmd())

	return rootCmd
}
