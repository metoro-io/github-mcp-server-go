package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// VERSION is the server version
	VERSION = "1.0.0"
	// USER_AGENT is the user agent sent with requests
	USER_AGENT = "modelcontextprotocol/servers/github-go/v" + VERSION
	// GITHUB_TOKEN_ENV_VAR is the environment variable name for the GitHub token
	GITHUB_TOKEN_ENV_VAR = "GITHUB_PERSONAL_ACCESS_TOKEN"
)

// APIRequirements contains the authentication information for GitHub API
type APIRequirements struct {
	Token string
}

// GetGitHubAPIRequirementsFromContext extracts GitHub authentication information from the request context
func GetGitHubAPIRequirementsFromContext(ctx context.Context) *APIRequirements {
	// Try to get the gin context from the context
	c := ctx.Value("ginContext")
	if c == nil {
		// Fall back to the previous method for backward compatibility
		reqVal := ctx.Value("http_request")
		if reqVal == nil {
			return nil
		}

		httpReq, ok := reqVal.(*http.Request)
		if !ok {
			return nil
		}

		// Extract authorization header
		authHeader := httpReq.Header.Get("Authorization")
		if authHeader != "" {
			// The header might be in the format "Bearer <token>" or just "<token>"
			token := authHeader
			if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
				token = authHeader[7:] // Remove "Bearer " prefix
			}

			return &APIRequirements{
				Token: token,
			}
		}

		return nil
	}

	ginContext, ok := c.(*gin.Context)
	if !ok {
		return nil
	}

	if ginContext.Request.Header.Get("Authorization") != "" {
		// The header might be in the format "Bearer <token>" or just "<token>"
		token := ginContext.Request.Header.Get("Authorization")
		if strings.HasPrefix(strings.ToLower(token), "bearer ") {
			token = token[7:] // Remove "Bearer " prefix
		}

		return &APIRequirements{
			Token: token,
		}
	}

	return nil
}

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
func GitHubRequest(urlStr string, method string, body interface{}, apiReqs *APIRequirements) (interface{}, error) {
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

	// Use token from provided APIRequirements if available, otherwise fall back to environment variable
	var token string
	if apiReqs != nil && apiReqs.Token != "" {
		token = apiReqs.Token
	} else {
		token = os.Getenv(GITHUB_TOKEN_ENV_VAR)
	}

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

	// GitHub usernames and organization names can contain alphanumeric characters and hyphens
	// They must start with an alphanumeric character and can't have consecutive hyphens
	// Max length is 39 characters
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]([a-zA-Z0-9]|-)*[a-zA-Z0-9]$`, sanitized)
	if !match && len(sanitized) > 1 {
		return "", fmt.Errorf("owner name must start and end with a letter or number, can contain hyphens (but not consecutive ones), and can be up to 39 characters")
	} else if !match {
		return "", fmt.Errorf("owner name must be alphanumeric")
	}

	if len(sanitized) > 39 {
		return "", fmt.Errorf("owner name is too long (max 39 characters)")
	}

	return sanitized, nil
}

// CheckBranchExists checks if a branch exists in a repository
func CheckBranchExists(owner, repo, branch string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches/%s", owner, repo, branch)
	_, err := GitHubRequest(url, "GET", nil, nil)
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
	_, err := GitHubRequest(url, "GET", nil, nil)
	if err != nil {
		if _, ok := err.(*GitHubResourceNotFoundError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
