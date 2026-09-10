package pr

import "gitai/groq"

func CreatePRDescription(gitDiff string) (string, error) {
	prDescription, err := groq.CreatePrDescriptionWithGroq(gitDiff)

	if err != nil {
		return "", err
	}

	return prDescription, nil
}
