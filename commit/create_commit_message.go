package commit

import (
	"gitAi/config"
	"gitAi/groq"
)

func CreateCommitMessage(gitDiff string, userConfig *config.UserConfig) (string, error) {
	commitMessage, err := groq.CreateCommitMessageWithGroq(gitDiff, userConfig)

	if err != nil {
		return "", err
	}

	return commitMessage, nil
}