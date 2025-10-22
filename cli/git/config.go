package git

import (
	"os"
	"os/exec"
	"strings"
)

func GetGitEditor() string {
	if gitEditor := os.Getenv("GIT_EDITOR"); gitEditor != "" {
		return gitEditor
	}

	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}

	cmd := exec.Command("git", "config", "--get", "core.editor")
	if output, err := cmd.Output(); err == nil {
		if editor := strings.TrimSpace(string(output)); editor != "" {
			return editor
		}
	}

	return "vi"
}