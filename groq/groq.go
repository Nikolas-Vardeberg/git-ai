package groq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitAi/config"
	"gitAi/git"
	"gitAi/server"
	"gitAi/version"
	"io"
	"net/http"
	"runtime"
	"time"
)

type CommitRequest struct {
	GitDiff    string             `json:"gitDiff"`
	Version    string             `json:"version"`
	Name       string             `json:"name"`
	RepoName   string             `json:"repoName"`
	UserConfig *config.UserConfig `json:"userConfig"`
	System         string             `json:"system,omitempty"`
}

type commitData struct {
	CommitMessage string `json:"commitMessage"`
}

type commitResp struct {
	Error *string     `json:"error,omitempty"`
	Data  *commitData `json:"data,omitempty"`
}

func CreateCommitMessageWithGroq(gitDiff string, userConfig *config.UserConfig) (string, error) {
	gitName := git.GetGitName()
	repoName := git.GetRepoName()

	payload := CommitRequest{
		UserConfig: userConfig,
		Version:    version.Get(),
		GitDiff:    gitDiff,
		Name:       gitName,
		RepoName:   repoName,
		System:         getOS(),
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost,
		server.ServerConfig.BaseURL+"/api/commit",
		bytes.NewReader(buf),
	)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var out commitResp

	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if out.Error != nil {
		return "", fmt.Errorf("%s", *out.Error)
	}

	if out.Data == nil {
		return "", fmt.Errorf("no data in response")
	}

	if out.Data.CommitMessage == "" {
		return "", fmt.Errorf("empty commit message from server")
	}

	return out.Data.CommitMessage, nil
}

func getOS() string {
	return runtime.GOOS
}