package timeline

import (
	"fmt"
	"gitai/git"
	"gitai/ui"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func Main() {
	choice := timelinePrompt("Choose timeline for commit analysis:")

	var timelineCommits []string
	var dateRange string
	var err error

	switch choice {
	case "today":
		timelineCommits, err = git.GetCommitsToday()
		dateRange = "today"
	default:
		ui.Box(ui.BoxOptions{Message: "Invalid choice.", Variant: ui.Error})
		return
	}

	if err != nil {
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Failed to get timeline commits: %v", err), Variant: ui.Error})
		os.Exit(1)
	}

	if len(timelineCommits) == 0 {
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("No commits found for the selected period (%s).", dateRange), Variant: ui.Warning})
		return
	}

	commitList := ""
	for i, commit := range timelineCommits {
		commitList += fmt.Sprintf("%d. %s\n", i+1, commit)
	}
	ui.Box(ui.BoxOptions{
		Title:   fmt.Sprintf("Found %d commits from %s", len(timelineCommits), dateRange),
		Message: strings.TrimSpace(commitList),
	})
}

func timelinePrompt(message string) string {
	var choice string

	err := huh.NewSelect[string]().
		Title(message).
		Description("Select an option using arrow keys or j,k and press Enter").
		Options(
			huh.NewOption("Today", "today"),
		).
		Value(&choice).
		Height(5).
		WithTheme(ui.GetHuhPrimaryTheme()).
		Run()

	if err != nil {
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Error running prompt: %v", err), Variant: ui.Error})
		os.Exit(1)
	}

	return choice
}
