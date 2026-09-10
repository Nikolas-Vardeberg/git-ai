package cmd

import (
	"gitai/pr"

	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Generate pull request descriptions from your staged changes",
	Long:  `Analyze your staged changes and generate an AI-powered pull request description.`,
	Run: func(cmd *cobra.Command, args []string) {
		pr.Main(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	prCmd.Flags().BoolP("print", "p", false, "Print PR description to stdout (no interactive UI)")
}
