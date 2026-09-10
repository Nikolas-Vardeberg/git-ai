package pr

import (
	"gitai/ui"
)

func HandlePRFlow(prDescription, fullPrompt string) {
	HandlePRFlowWithHistory(prDescription, fullPrompt, []string{})
}

func HandlePRFlowWithHistory(prDescription, fullPrompt string, previousMessages []string) {
	ui.Box(ui.BoxOptions{Title: "PR Description", Message: prDescription})
}
