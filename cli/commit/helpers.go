package commit

import (
	"fmt"
	"gitai/git"
	"gitai/ui"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
)

func HandleCommitFlow(commitMessage, fullPrompt string) {
	HandleCommitFlowWithHistory(commitMessage, fullPrompt, []string{})
}

func HandleCommitFlowWithHistory(commitMessage, fullPrompt string, previousMessages []string) {

	ui.Box(ui.BoxOptions{Title: "Commit message", Message: commitMessage})

	choice := choicePrompt()

	switch choice {
	case "commit":
		executeCommit(commitMessage, false)
	case "commit-push":
		executeCommit(commitMessage, true)
	case "exit":
		ui.RenderTitle("Bye!")
		fmt.Println()
		os.Exit(0)
	}
}

func choicePrompt() string {
	var choice string

	err := huh.NewSelect[string]().
		Title("What would you like to do next?").
		Description("Select an option using arrow keys or j,k and press Enter").
		Options(
			huh.NewOption("Commit this message", "commit"),
			huh.NewOption("Commit and push", "commit-push"),
			huh.NewOption("Exit", "exit"),
		).
		Value(&choice).
		Height(9).
		WithTheme(ui.GetHuhPrimaryTheme()).
		Run()

	if err != nil {
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Error running prompt: %v", err), Variant: ui.Error})
		os.Exit(1)
	}

	return choice
}

func openInEditor(message string) (string, error) {
	editor := git.GetGitEditor()

	tmpFile, err := os.CreateTemp("", "diny-commit-*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(message); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	editorArgs := strings.Fields(editor)
	editorCmd := editorArgs[0]
	args := append(editorArgs[1:], tmpFile.Name())

	cmd := exec.Command(editorCmd, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("editor exited with error: %v", err)
	}

	editedContent, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to read edited file: %v", err)
	}

	return strings.TrimSpace(string(editedContent)), nil
}

func executeCommit(commitMessage string, push bool) {
	var output []byte
	var err error

	spinnerErr := ui.WithSpinner("Committing...", func() error {
		commitCmd := exec.Command("git", "commit", "-m", commitMessage)
		output, err = commitCmd.CombinedOutput()
		return err
	})

	if spinnerErr != nil {
		if len(output) > 0 {
			fmt.Fprint(os.Stderr, string(output))
		}
		ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Commit failed: %v", spinnerErr), Variant: ui.Error})
		os.Exit(1)
	}
	ui.Box(ui.BoxOptions{Message: "Commited!", Variant: ui.Success})

	if push {
		var pushOutput []byte
		var pushErr error

		pushSpinnerErr := ui.WithSpinner("Pushing...", func() error {
			pushCmd := exec.Command("git", "push")
			pushOutput, pushErr = pushCmd.CombinedOutput()
			return pushErr
		})

		if pushSpinnerErr != nil {
			if len(pushOutput) > 0 {
				fmt.Fprint(os.Stderr, string(pushOutput))
			}
			ui.Box(ui.BoxOptions{Message: fmt.Sprintf("Push failed: %v", pushSpinnerErr), Variant: ui.Error})
			os.Exit(1)
		}
		ui.Box(ui.BoxOptions{Message: "Pushed!", Variant: ui.Success})
	}

	fmt.Println()
}