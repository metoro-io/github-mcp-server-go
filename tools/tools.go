package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/metoro-io/github-mcp-server-go/common"
	"github.com/metoro-io/github-mcp-server-go/operations"
	mcpgolang "github.com/metoro-io/mcp-golang"
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
func SearchRepositoriesHandler(ctx context.Context, args operations.SearchRepositoriesOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.SearchRepositories(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// CreateRepositoryHandler handles create_repository requests
func CreateRepositoryHandler(ctx context.Context, args operations.CreateRepositoryOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.CreateRepository(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// ForkRepositoryHandler handles fork_repository requests
func ForkRepositoryHandler(ctx context.Context, args operations.ForkRepositoryOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.ForkRepository(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// CreateBranchHandler handles create_branch requests
func CreateBranchHandler(ctx context.Context, args operations.CreateBranchOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.CreateBranchFromRef(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// GetFileContentsHandler handles get_file_contents requests
func GetFileContentsHandler(ctx context.Context, args operations.GetFileContentsOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.GetFileContents(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// CreateOrUpdateFileHandler handles create_or_update_file requests
func CreateOrUpdateFileHandler(ctx context.Context, args operations.CreateOrUpdateFileOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.CreateOrUpdateFile(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// PushFilesHandler handles push_files requests
func PushFilesHandler(ctx context.Context, args operations.PushFilesOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.PushFiles(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// CreateIssueHandler handles create_issue requests
func CreateIssueHandler(ctx context.Context, args operations.CreateIssueOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.CreateIssue(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// GetIssueHandler handles get_issue requests
func GetIssueHandler(ctx context.Context, args operations.GetIssueOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.GetIssue(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// ListIssuesHandler handles list_issues requests
func ListIssuesHandler(ctx context.Context, args operations.ListIssuesOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.ListIssues(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// UpdateIssueHandler handles update_issue requests
func UpdateIssueHandler(ctx context.Context, args operations.UpdateIssueOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.UpdateIssue(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// AddIssueCommentHandler handles add_issue_comment requests
func AddIssueCommentHandler(ctx context.Context, args operations.IssueCommentOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.AddIssueComment(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// ListCommitsHandler handles list_commits requests
func ListCommitsHandler(ctx context.Context, args operations.ListCommitsOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.ListCommits(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// SearchCodeHandler handles search_code requests
func SearchCodeHandler(ctx context.Context, args operations.SearchCodeOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.SearchCode(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// SearchIssuesHandler handles search_issues requests
func SearchIssuesHandler(ctx context.Context, args operations.SearchIssuesOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.SearchIssues(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// SearchUsersHandler handles search_users requests
func SearchUsersHandler(ctx context.Context, args operations.SearchUsersOptions) (*mcpgolang.ToolResponse, error) {
	apiReqs := common.GetGitHubAPIRequirementsFromContext(ctx)

	result, err := operations.SearchUsers(&args, apiReqs)
	if err != nil {
		return nil, formatError(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}

	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(jsonData))), nil
}

// formatError formats errors for response
func formatError(err error) error {
	if common.IsGitHubError(err) {
		return fmt.Errorf(common.FormatGitHubError(err))
	}
	return err
}
