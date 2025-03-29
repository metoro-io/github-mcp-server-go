package common

import (
	"fmt"
	"time"
)

// GitHubError is the base error type for GitHub API errors
type GitHubError struct {
	Message  string
	Status   int
	Response interface{}
}

func (e *GitHubError) Error() string {
	return fmt.Sprintf("GitHub API Error: %s", e.Message)
}

// GitHubValidationError represents validation errors from GitHub API
type GitHubValidationError struct {
	GitHubError
}

func (e *GitHubValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", e.Message)
}

// GitHubResourceNotFoundError represents 404 not found errors
type GitHubResourceNotFoundError struct {
	GitHubError
}

func (e *GitHubResourceNotFoundError) Error() string {
	return fmt.Sprintf("Not Found: %s", e.Message)
}

// GitHubAuthenticationError represents authentication failures
type GitHubAuthenticationError struct {
	GitHubError
}

func (e *GitHubAuthenticationError) Error() string {
	return fmt.Sprintf("Authentication Failed: %s", e.Message)
}

// GitHubPermissionError represents permission/authorization errors
type GitHubPermissionError struct {
	GitHubError
}

func (e *GitHubPermissionError) Error() string {
	return fmt.Sprintf("Permission Denied: %s", e.Message)
}

// GitHubRateLimitError represents rate limit exceeded errors
type GitHubRateLimitError struct {
	GitHubError
	ResetAt time.Time
}

func (e *GitHubRateLimitError) Error() string {
	return fmt.Sprintf("Rate Limit Exceeded: %s\nResets at: %s", e.Message, e.ResetAt.Format(time.RFC3339))
}

// GitHubConflictError represents conflict errors
type GitHubConflictError struct {
	GitHubError
}

func (e *GitHubConflictError) Error() string {
	return fmt.Sprintf("Conflict: %s", e.Message)
}

// IsGitHubError checks if an error is a GitHub API error
func IsGitHubError(err error) bool {
	_, ok := err.(*GitHubError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubValidationError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubResourceNotFoundError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubAuthenticationError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubPermissionError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubRateLimitError)
	if ok {
		return true
	}

	_, ok = err.(*GitHubConflictError)
	return ok
}

// CreateGitHubError creates the appropriate GitHub error based on status code
func CreateGitHubError(status int, response interface{}) error {
	respMap, ok := response.(map[string]interface{})
	var message string
	if ok {
		if msg, ok := respMap["message"].(string); ok {
			message = msg
		}
	}

	if message == "" {
		message = "GitHub API error"
	}

	switch status {
	case 401:
		return &GitHubAuthenticationError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
		}
	case 403:
		return &GitHubPermissionError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
		}
	case 404:
		return &GitHubResourceNotFoundError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
		}
	case 409:
		return &GitHubConflictError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
		}
	case 422:
		return &GitHubValidationError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
		}
	case 429:
		resetAt := time.Now().Add(1 * time.Minute)
		if ok {
			if reset, ok := respMap["reset_at"].(string); ok {
				if t, err := time.Parse(time.RFC3339, reset); err == nil {
					resetAt = t
				}
			}
		}
		return &GitHubRateLimitError{
			GitHubError: GitHubError{
				Message:  message,
				Status:   status,
				Response: response,
			},
			ResetAt: resetAt,
		}
	default:
		return &GitHubError{
			Message:  message,
			Status:   status,
			Response: response,
		}
	}
}

// FormatGitHubError formats a GitHub error for display
func FormatGitHubError(err error) string {
	switch e := err.(type) {
	case *GitHubValidationError:
		details := ""
		if e.Response != nil {
			details = fmt.Sprintf("\nDetails: %v", e.Response)
		}
		return fmt.Sprintf("Validation Error: %s%s", e.Message, details)
	case *GitHubResourceNotFoundError:
		return fmt.Sprintf("Not Found: %s", e.Message)
	case *GitHubAuthenticationError:
		return fmt.Sprintf("Authentication Failed: %s", e.Message)
	case *GitHubPermissionError:
		return fmt.Sprintf("Permission Denied: %s", e.Message)
	case *GitHubRateLimitError:
		return fmt.Sprintf("Rate Limit Exceeded: %s\nResets at: %s", e.Message, e.ResetAt.Format(time.RFC3339))
	case *GitHubConflictError:
		return fmt.Sprintf("Conflict: %s", e.Message)
	case *GitHubError:
		return fmt.Sprintf("GitHub API Error: %s", e.Message)
	default:
		return err.Error()
	}
}
