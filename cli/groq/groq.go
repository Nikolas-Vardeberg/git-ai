package groq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitai/server"
	"io"
	"net/http"
	"time"
)

type CommitRequest struct {
	GitDiff    string             `json:"gitDiff"`
}

type reviewData struct {
	ReviewMessage string `json:"reviewMessage"`
}


type commitData struct {
	CommitMessage string `json:"commitMessage"`
}

type commitResp struct {
	Error *string     `json:"error,omitempty"`
	Data  *commitData `json:"data,omitempty"`
}

type reviewResp struct {
	Error *string    `json:"error,omitempty"`
	Data  *reviewData `json:"data,omitempty"`
}

type prDescriptionData struct {
	PrDescription string `json:"prDescription"`
	PRDescription string `json:"PRDescription"`
	Description   string `json:"description"`
	CommitMessage string `json:"commitMessage"`
}

type prDescriptionResp struct {
	Error *string            `json:"error,omitempty"`
	Data  *prDescriptionData `json:"data,omitempty"`
}

func CreateCommitMessageWithGroq(gitDiff string) (string, error) {

	payload := CommitRequest{
		GitDiff:    gitDiff,
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

func CreateCommitReviewWithGroq(gitDiff string) (string, error) {
	payload := CommitRequest{
		GitDiff:    gitDiff,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost,
		server.ServerConfig.BaseURL+"/api/review",
		bytes.NewReader(buf),
	)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var out reviewResp

	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if out.Error != nil {
		return "", fmt.Errorf("%s", *out.Error)
	}

	if out.Data == nil {
		return "", fmt.Errorf("no data in response")
	}

	if out.Data.ReviewMessage == "" {
		return "", fmt.Errorf("empty review message from server")
	}

	return out.Data.ReviewMessage, nil
}

func CreatePrDescriptionWithGroq(gitDiff string) (string, error) {
	payload := CommitRequest{
		GitDiff: gitDiff,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost,
		server.ServerConfig.BaseURL+"/api/pr-description",
		bytes.NewReader(buf),
	)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var out prDescriptionResp

	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if out.Error != nil {
		return "", fmt.Errorf("%s", *out.Error)
	}

	if out.Data == nil {
		return "", fmt.Errorf("no data in response")
	}

	msg := out.Data.PrDescription
	if msg == "" {
		msg = out.Data.PRDescription
	}
	if msg == "" {
		msg = out.Data.Description
	}
	if msg == "" {
		msg = out.Data.CommitMessage
	}

	if msg == "" {
		return "", fmt.Errorf("empty PR description from server")
	}

	return msg, nil
}

func CreatePRDescriptionWithGroq(gitDiff string) (string, error) {
	return CreatePrDescriptionWithGroq(gitDiff)
}
