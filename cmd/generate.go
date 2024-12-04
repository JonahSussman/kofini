package cmd

import "github.com/spf13/cobra"

func GetGenerateCmd() *cobra.Command {
	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate deployment assets.",
		Long:  "Generate deployment assets. This command takes a Helm template and creates a fully-rendered Helm chart for deploying the application on a target platform, such as OCP.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return generateCmd
}
