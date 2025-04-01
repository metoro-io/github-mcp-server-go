package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// CreateIssueOptions defines options for creating an issue
type CreateIssueOptions struct {
	Owner     string   `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo      string   `json:"repo" jsonschema:"description=The name of the repository where the issue will be created"`
	Title     string   `json:"title" jsonschema:"description=The title of the issue"`
	Body      string   `json:"body" jsonschema:"description=The body content of the issue"`
	Assignees []string `json:"assignees,omitempty" jsonschema:"description=A list of usernames to assign to this issue"`
	Labels    []string `json:"labels,omitempty" jsonschema:"description=A list of label names to add to this issue"`
}

// Validate validates the CreateIssueOptions
func (o *CreateIssueOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}

// GetIssueOptions defines options for getting an issue
type GetIssueOptions struct {
	Owner  string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo   string `json:"repo" jsonschema:"description=The name of the repository containing the issue"`
	Number int    `json:"number" jsonschema:"description=The issue number to retrieve"`
}

// Validate validates the GetIssueOptions
func (o *GetIssueOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Number <= 0 {
		return fmt.Errorf("issue number must be a positive integer")
	}
	return nil
}

// ListIssuesOptions defines options for listing issues
type ListIssuesOptions struct {
	Owner     string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo      string `json:"repo" jsonschema:"description=The name of the repository to list issues from"`
	State     string `json:"state,omitempty" jsonschema:"description=Filter issues by state. Can be one of: open closed all. Default: open"`
	Sort      string `json:"sort,omitempty" jsonschema:"description=What to sort results by. Can be one of: created updated comments. Default: created"`
	Direction string `json:"direction,omitempty" jsonschema:"description=The direction of the sort. Can be one of: asc desc. Default: desc"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
}

// Validate validates the ListIssuesOptions
func (o *ListIssuesOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.State != "" && o.State != "open" && o.State != "closed" && o.State != "all" {
		return fmt.Errorf("state must be one of: open, closed, all")
	}
	if o.Sort != "" && o.Sort != "created" && o.Sort != "updated" && o.Sort != "comments" {
		return fmt.Errorf("sort must be one of: created, updated, comments")
	}
	if o.Direction != "" && o.Direction != "asc" && o.Direction != "desc" {
		return fmt.Errorf("direction must be one of: asc, desc")
	}
	return nil
}

// UpdateIssueOptions defines options for updating an issue
type UpdateIssueOptions struct {
	Owner     string   `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo      string   `json:"repo" jsonschema:"description=The name of the repository containing the issue to update"`
	Number    int      `json:"number" jsonschema:"description=The issue number to update"`
	Title     string   `json:"title,omitempty" jsonschema:"description=The new title of the issue"`
	Body      string   `json:"body,omitempty" jsonschema:"description=The new body content of the issue"`
	State     string   `json:"state,omitempty" jsonschema:"description=The state of the issue. Can be one of: open closed"`
	Assignees []string `json:"assignees,omitempty" jsonschema:"description=A list of usernames to assign to this issue. Pass an empty array to clear all assignees"`
	Labels    []string `json:"labels,omitempty" jsonschema:"description=A list of label names to add to this issue. Pass an empty array to clear all labels"`
}

// Validate validates the UpdateIssueOptions
func (o *UpdateIssueOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Number <= 0 {
		return fmt.Errorf("issue number must be a positive integer")
	}
	if o.State != "" && o.State != "open" && o.State != "closed" {
		return fmt.Errorf("state must be one of: open, closed")
	}
	return nil
}

// IssueCommentOptions defines options for adding a comment to an issue
type IssueCommentOptions struct {
	Owner  string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo   string `json:"repo" jsonschema:"description=The name of the repository containing the issue to comment on"`
	Number int    `json:"number" jsonschema:"description=The issue number to add a comment to"`
	Body   string `json:"body" jsonschema:"description=The contents of the comment to add"`
}

// Validate validates the IssueCommentOptions
func (o *IssueCommentOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Number <= 0 {
		return fmt.Errorf("issue number must be a positive integer")
	}
	if o.Body == "" {
		return fmt.Errorf("comment body is required")
	}
	return nil
}

// CreateIssue creates a new issue in a GitHub repository
func CreateIssue(options *CreateIssueOptions, apiReqs *common.APIRequirements) (*common.GitHubIssue, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", options.Owner, options.Repo)

	requestBody := map[string]interface{}{
		"title": options.Title,
		"body":  options.Body,
	}

	if len(options.Assignees) > 0 {
		requestBody["assignees"] = options.Assignees
	}

	if len(options.Labels) > 0 {
		requestBody["labels"] = options.Labels
	}

	resp, err := common.GitHubRequest(url, "POST", requestBody, apiReqs)
	if err != nil {
		return nil, err
	}

	var issue common.GitHubIssue
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

// GetIssue gets details of a specific issue in a GitHub repository
func GetIssue(options *GetIssueOptions, apiReqs *common.APIRequirements) (*common.GitHubIssue, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d",
		options.Owner, options.Repo, options.Number)

	resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
	if err != nil {
		return nil, err
	}

	var issue common.GitHubIssue
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

// ListIssues lists issues in a GitHub repository
func ListIssues(options *ListIssuesOptions, apiReqs *common.APIRequirements) ([]common.GitHubIssue, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", options.Owner, options.Repo)

	params := make(map[string]string)
	if options.State != "" {
		params["state"] = options.State
	}
	if options.Sort != "" {
		params["sort"] = options.Sort
	}
	if options.Direction != "" {
		params["direction"] = options.Direction
	}
	if options.Page > 0 {
		params["page"] = strconv.Itoa(options.Page)
	}
	if options.PerPage > 0 {
		params["per_page"] = strconv.Itoa(options.PerPage)
	}

	if len(params) > 0 {
		var err error
		url, err = common.BuildURL(url, params)
		if err != nil {
			return nil, err
		}
	}

	resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
	if err != nil {
		return nil, err
	}

	var issues []common.GitHubIssue
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// UpdateIssue updates an existing issue in a GitHub repository
func UpdateIssue(options *UpdateIssueOptions, apiReqs *common.APIRequirements) (*common.GitHubIssue, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d",
		options.Owner, options.Repo, options.Number)

	requestBody := make(map[string]interface{})
	if options.Title != "" {
		requestBody["title"] = options.Title
	}
	if options.Body != "" {
		requestBody["body"] = options.Body
	}
	if options.State != "" {
		requestBody["state"] = options.State
	}
	if options.Assignees != nil {
		requestBody["assignees"] = options.Assignees
	}
	if options.Labels != nil {
		requestBody["labels"] = options.Labels
	}

	resp, err := common.GitHubRequest(url, "PATCH", requestBody, apiReqs)
	if err != nil {
		return nil, err
	}

	var issue common.GitHubIssue
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

// AddIssueComment adds a comment to an existing issue
func AddIssueComment(options *IssueCommentOptions, apiReqs *common.APIRequirements) (interface{}, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/comments",
		options.Owner, options.Repo, options.Number)

	requestBody := map[string]string{
		"body": options.Body,
	}

	resp, err := common.GitHubRequest(url, "POST", requestBody, apiReqs)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
