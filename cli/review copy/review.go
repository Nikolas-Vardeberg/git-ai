package review

import (
	"fmt"
	"gitai/git"
	"gitai/ui"
	"os"

	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string) {
	printMode, _ := cmd.Flags().GetBool("print")

	diff := getCommitData(printMode)

	if printMode {
		reviewMessage, err := CreateReviewMessage(diff)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating review message: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(reviewMessage)
		return
	}

	var reviewMessage string
	err := ui.WithSpinner("Generating your review message...", func() error {
		var genErr error
		reviewMessage, genErr = CreateReviewMessage(diff)
		return genErr
	})

	if err != nil {
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("%v", err), Variant: ui.Error})
		os.Exit(1)
	}

	HandleReviewFlow(reviewMessage, diff)
}

func getCommitData(isQuietMode bool) (string) {
	gitDiff, err := git.GetGitDiff()

	if err != nil {
		if isQuietMode {
			fmt.Fprintf(os.Stderr, "Failed to get git diff: %v\n", err)
		} else {
			ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Failed to get git diff: %v", err), Variant: ui.Error})
		}
		os.Exit(1)
	}

	if len(gitDiff) == 0 {
		if isQuietMode {
			fmt.Fprintf(os.Stderr, "No staged changes found. Stage files first with `git add`.\n")
		} else {
			ui.Box(ui.BoxOptions{Message: "No staged changes found. Stage files first with `git add`.", Variant: ui.Warning})
		}
		os.Exit(0)
	}

	return gitDiff
}