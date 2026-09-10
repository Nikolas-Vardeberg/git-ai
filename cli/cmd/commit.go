package cmd

import (
	"gitai/commit"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate clean, conventional commit messages from staged changes",
	Long: `Reads your staged changes and proposes concise, conventional-friendly
commit messages. Use it to keep a tidy, consistent history—interactively or in
scripts.

Examples:
  gitai commit                           # Interactive mode with options
  gitai commit --print                   # Print message to stdout only
  gitai commit --print | git commit -F - # Generate and commit directly
  gitai commit --print | pbcopy          # Copy to clipboard (macOS)
  gitai commit --print | xclip -sel clip # Copy to clipboard (Linux)
  gitai commit --print | clip            # Copy to clipboard (Windows)`,
	Run: func(cmd *cobra.Command, args []string) {
		commit.Main(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().BoolP("print", "p", false, "Print commit message to stdout (no interactive UI)")
}