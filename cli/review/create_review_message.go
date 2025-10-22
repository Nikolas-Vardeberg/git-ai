package review

import "gitai/groq"

func CreateReviewMessage(gitDiff string) (string, error) {
	reviewMessage, err := groq.CreateCommitReviewWithGroq(gitDiff)

	if err != nil {
		return "", err
	}

	return reviewMessage, nil
}