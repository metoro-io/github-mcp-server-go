package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// CreateIssueOptions defines options for creating an issue
type CreateIssueOptions struct {
	Owner     string   `json:"owner"`
	Repo      string   `json:"repo"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
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
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Number int    `json:"number"`
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
	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	State     string `json:"state,omitempty"`     // open, closed, all
	Sort      string `json:"sort,omitempty"`      // created, updated, comments
	Direction string `json:"direction,omitempty"` // asc, desc
	Page      int    `json:"page,omitempty"`
	PerPage   int    `json:"per_page,omitempty"`
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
	Owner     string   `json:"owner"`
	Repo      string   `json:"repo"`
	Number    int      `json:"number"`
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	State     string   `json:"state,omitempty"` // open, closed
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
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
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Number int    `json:"number"`
	Body   string `json:"body"`
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
