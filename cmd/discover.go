package cmd

import "github.com/spf13/cobra"

func GetDiscoverCmd() *cobra.Command {
	var discoverCmd = &cobra.Command{
		Use:   "discover",
		Short: "Discover application configurations.",
		Long:  "Discover application configurations. This command retrieves application configuration data from a specified platform, such as PCF, and outputs it as a YAML file. This file forms the basis for further transformations and deployments",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return discoverCmd
}
