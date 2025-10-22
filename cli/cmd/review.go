package cmd

import (
	"gitai/review"

	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use: "review",
	Short: "Generate code review comments from your commits",
	Long: `Analyze your recent commits and generate AI-powered code review comments.`,
	Run: func(cmd *cobra.Command, args []string) {
		review.Main(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
	
	reviewCmd.Flags().BoolP("print", "p", false, "Print review message to stdout (no interactive UI)")
}