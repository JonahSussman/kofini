package cmd

import "github.com/spf13/cobra"

type discoverFlags struct {
	platform string
}

func GetDiscoverCmd() *cobra.Command {
	flags := discoverFlags{}
	discoverCmd := &cobra.Command{
		Use:   "discover",
		Short: "Discover application configurations.",
		Long:  "Discover application configurations. This command retrieves application configuration data from a specified platform, such as PCF, and outputs it as a YAML file. This file forms the basis for further transformations and deployments",
		Run: func(cmd *cobra.Command, args []string) {
			discoverHandler(cmd, args, flags)
		},
	}

	discoverCmd.Flags().StringVarP(&flags.platform, "platform", "p", "", "The platform to discover configurations from.")

	return discoverCmd
}

func discoverHandler(cmd *cobra.Command, args []string, flags discoverFlags) {
	cmd.Help()
}
