package cmd

import "github.com/spf13/cobra"

type generateFlags struct {
}

func GetGenerateCmd() *cobra.Command {
	flags := generateFlags{}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate deployment assets.",
		Long:  "Generate deployment assets. This command takes a Helm template and creates a fully-rendered Helm chart for deploying the application on a target platform, such as OCP.",
		Run: func(cmd *cobra.Command, args []string) {
			generateHandler(cmd, args, flags)
		},
	}

	return generateCmd
}

func generateHandler(cmd *cobra.Command, args []string, flags generateFlags) {
	cmd.Help()
}
