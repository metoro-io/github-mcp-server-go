package tools

import (
	"encoding/json"
	"fmt"

	"github.com/metoro-k8s/github-mcp-server-go/common"
	"github.com/metoro-k8s/github-mcp-server-go/operations"
)

// GitHubTool represents a GitHub tool handler
type GitHubTool struct {
	Name        string
	Description string
	Handler     interface{}
}

// GitHubToolsList is the list of tools available for GitHub operations
var GitHubToolsList = []GitHubTool{
	{
		Name:        "search_repositories",
		Description: "Search for GitHub repositories",
		Handler:     SearchRepositoriesHandler,
	},
	{
		Name:        "create_repository",
		Description: "Create a new GitHub repository in your account",
		Handler:     CreateRepositoryHandler,
	},
	{
		Name:        "fork_repository",
		Description: "Fork a GitHub repository to your account or specified organization",
		Handler:     ForkRepositoryHandler,
	},
	{
		Name:        "create_branch",
		Description: "Create a new branch in a GitHub repository",
		Handler:     CreateBranchHandler,
	},
	{
		Name:        "create_or_update_file",
		Description: "Create or update a single file in a GitHub repository",
		Handler:     CreateOrUpdateFileHandler,
	},
	{
		Name:        "get_file_contents",
		Description: "Get the contents of a file or directory from a GitHub repository",
		Handler:     GetFileContentsHandler,
	},
	{
		Name:        "push_files",
		Description: "Push multiple files to a GitHub repository in a single commit",
		Handler:     PushFilesHandler,
	},
	{
		Name:        "create_issue",
		Description: "Create a new issue in a GitHub repository",
		Handler:     CreateIssueHandler,
	},
	{
		Name:        "get_issue",
		Description: "Get details of a specific issue in a GitHub repository",
		Handler:     GetIssueHandler,
	},
	{
		Name:        "list_issues",
		Description: "List issues in a GitHub repository with filtering options",
		Handler:     ListIssuesHandler,
	},
	{
		Name:        "update_issue",
		Description: "Update an existing issue in a GitHub repository",
		Handler:     UpdateIssueHandler,
	},
	{
		Name:        "add_issue_comment",
		Description: "Add a comment to an existing issue",
		Handler:     AddIssueCommentHandler,
	},
	{
		Name:        "list_commits",
		Description: "Get list of commits of a branch in a GitHub repository",
		Handler:     ListCommitsHandler,
	},
	{
		Name:        "search_code",
		Description: "Search for code across GitHub repositories",
		Handler:     SearchCodeHandler,
	},
	{
		Name:        "search_issues",
		Description: "Search for issues and pull requests across GitHub repositories",
		Handler:     SearchIssuesHandler,
	},
	{
		Name:        "search_users",
		Description: "Search for users on GitHub",
		Handler:     SearchUsersHandler,
	},
}

// SearchRepositoriesHandler handles search_repositories requests
func SearchRepositoriesHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.SearchRepositoriesOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.SearchRepositories(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// CreateRepositoryHandler handles create_repository requests
func CreateRepositoryHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.CreateRepositoryOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.CreateRepository(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// ForkRepositoryHandler handles fork_repository requests
func ForkRepositoryHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.ForkRepositoryOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.ForkRepository(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// CreateBranchHandler handles create_branch requests
func CreateBranchHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.CreateBranchOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.CreateBranchFromRef(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// GetFileContentsHandler handles get_file_contents requests
func GetFileContentsHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.GetFileContentsOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.GetFileContents(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// CreateOrUpdateFileHandler handles create_or_update_file requests
func CreateOrUpdateFileHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.CreateOrUpdateFileOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.CreateOrUpdateFile(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// PushFilesHandler handles push_files requests
func PushFilesHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.PushFilesOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.PushFiles(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// CreateIssueHandler handles create_issue requests
func CreateIssueHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.CreateIssueOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.CreateIssue(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// GetIssueHandler handles get_issue requests
func GetIssueHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.GetIssueOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.GetIssue(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// ListIssuesHandler handles list_issues requests
func ListIssuesHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.ListIssuesOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.ListIssues(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// UpdateIssueHandler handles update_issue requests
func UpdateIssueHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.UpdateIssueOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.UpdateIssue(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// AddIssueCommentHandler handles add_issue_comment requests
func AddIssueCommentHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.IssueCommentOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.AddIssueComment(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// ListCommitsHandler handles list_commits requests
func ListCommitsHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.ListCommitsOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.ListCommits(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// SearchCodeHandler handles search_code requests
func SearchCodeHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.SearchCodeOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.SearchCode(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// SearchIssuesHandler handles search_issues requests
func SearchIssuesHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.SearchIssuesOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.SearchIssues(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// SearchUsersHandler handles search_users requests
func SearchUsersHandler(rawArgs json.RawMessage) (interface{}, error) {
	var options operations.SearchUsersOptions
	if err := json.Unmarshal(rawArgs, &options); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	result, err := operations.SearchUsers(&options)
	if err != nil {
		return nil, formatError(err)
	}

	return result, nil
}

// formatError formats errors for response
func formatError(err error) error {
	if common.IsGitHubError(err) {
		return fmt.Errorf(common.FormatGitHubError(err))
	}
	return err
}
