package review

import (
	"gitai/ui"
)

func HandleReviewFlow(reviewMessage, fullPrompt string) {
	HandleReviewFlowWithHistory(reviewMessage, fullPrompt, []string{})
}

func HandleReviewFlowWithHistory(reviewMessage, fullPrompt string, previousMessages []string) {
	ui.Box(ui.BoxOptions{Title: "Review message", Message: reviewMessage})
}