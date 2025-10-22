package commit

import (
	"gitAi/groq"
)

func CreateCommitMessage(gitDiff string) (string, error) {
	commitMessage, err := groq.CreateCommitMessageWithGroq(gitDiff)

	if err != nil {
		return "", err
	}

	return commitMessage, nil
}