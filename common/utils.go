package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

const (
	// VERSION is the server version
	VERSION = "1.0.0"
	// USER_AGENT is the user agent sent with requests
	USER_AGENT = "modelcontextprotocol/servers/github-go/v" + VERSION
)

// BuildURL builds a URL with query parameters
func BuildURL(baseURL string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range params {
		if v != "" {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GitHubRequest sends an HTTP request to the GitHub API
func GitHubRequest(urlStr string, method string, body interface{}) (interface{}, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(method, urlStr, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &result); err != nil {
			return string(responseBody), nil
		}
	}

	if resp.StatusCode >= 400 {
		return nil, CreateGitHubError(resp.StatusCode, result)
	}

	return result, nil
}

// ValidateBranchName validates a branch name according to Git rules
func ValidateBranchName(branch string) (string, error) {
	sanitized := branch
	if sanitized == "" {
		return "", fmt.Errorf("branch name cannot be empty")
	}
	if regexp.MustCompile(`\.\.`).MatchString(sanitized) {
		return "", fmt.Errorf("branch name cannot contain '..'")
	}
	if regexp.MustCompile(`[\s~^:?*[\\\]]`).MatchString(sanitized) {
		return "", fmt.Errorf("branch name contains invalid characters")
	}
	if sanitized[0] == '/' || sanitized[len(sanitized)-1] == '/' {
		return "", fmt.Errorf("branch name cannot start or end with '/'")
	}
	if len(sanitized) >= 5 && sanitized[len(sanitized)-5:] == ".lock" {
		return "", fmt.Errorf("branch name cannot end with '.lock'")
	}
	return sanitized, nil
}

// ValidateRepositoryName validates a repository name
func ValidateRepositoryName(name string) (string, error) {
	sanitized := name
	if sanitized == "" {
		return "", fmt.Errorf("repository name cannot be empty")
	}

	match, _ := regexp.MatchString(`^[a-zA-Z0-9_.-]+$`, sanitized)
	if !match {
		return "", fmt.Errorf("repository name can only contain letters, numbers, hyphens, periods, and underscores")
	}

	if sanitized[0] == '.' || sanitized[len(sanitized)-1] == '.' {
		return "", fmt.Errorf("repository name cannot start or end with a period")
	}

	return sanitized, nil
}

// ValidateOwnerName validates an owner name (user or organization)
func ValidateOwnerName(owner string) (string, error) {
	sanitized := owner
	if sanitized == "" {
		return "", fmt.Errorf("owner name cannot be empty")
	}

	match, _ := regexp.MatchString(`^[a-zA-Z0-9](?:[a-zA-Z0-9]|-(?=[a-zA-Z0-9])){0,38}$`, sanitized)
	if !match {
		return "", fmt.Errorf("owner name must start with a letter or number and can contain up to 39 characters")
	}

	return sanitized, nil
}

// CheckBranchExists checks if a branch exists in a repository
func CheckBranchExists(owner, repo, branch string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches/%s", owner, repo, branch)
	_, err := GitHubRequest(url, "GET", nil)
	if err != nil {
		if _, ok := err.(*GitHubResourceNotFoundError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CheckUserExists checks if a GitHub user exists
func CheckUserExists(username string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	_, err := GitHubRequest(url, "GET", nil)
	if err != nil {
		if _, ok := err.(*GitHubResourceNotFoundError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
