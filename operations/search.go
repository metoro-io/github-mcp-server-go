package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-k8s/github-mcp-server-go/common"
)

// SearchCodeOptions defines options for searching code
type SearchCodeOptions struct {
	Query   string `json:"query"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
}

// Validate validates the SearchCodeOptions
func (o *SearchCodeOptions) Validate() error {
	if o.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

// SearchIssuesOptions defines options for searching issues and pull requests
type SearchIssuesOptions struct {
	Query   string `json:"query"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
}

// Validate validates the SearchIssuesOptions
func (o *SearchIssuesOptions) Validate() error {
	if o.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

// SearchUsersOptions defines options for searching users
type SearchUsersOptions struct {
	Query   string `json:"query"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
}

// Validate validates the SearchUsersOptions
func (o *SearchUsersOptions) Validate() error {
	if o.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

// SearchCode searches for code across GitHub repositories
func SearchCode(options *SearchCodeOptions) (*common.GitHubSearchCodeResponse, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := "https://api.github.com/search/code"

	params := map[string]string{
		"q": options.Query,
	}

	if options.Page > 0 {
		params["page"] = strconv.Itoa(options.Page)
	}

	if options.PerPage > 0 {
		params["per_page"] = strconv.Itoa(options.PerPage)
	}

	fullURL, err := common.BuildURL(url, params)
	if err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest(fullURL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var searchResp common.GitHubSearchCodeResponse
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// SearchIssues searches for issues and pull requests across GitHub repositories
func SearchIssues(options *SearchIssuesOptions) (*common.GitHubSearchIssuesResponse, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := "https://api.github.com/search/issues"

	params := map[string]string{
		"q": options.Query,
	}

	if options.Page > 0 {
		params["page"] = strconv.Itoa(options.Page)
	}

	if options.PerPage > 0 {
		params["per_page"] = strconv.Itoa(options.PerPage)
	}

	fullURL, err := common.BuildURL(url, params)
	if err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest(fullURL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var searchResp common.GitHubSearchIssuesResponse
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// SearchUsers searches for users on GitHub
func SearchUsers(options *SearchUsersOptions) (*common.GitHubSearchUsersResponse, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := "https://api.github.com/search/users"

	params := map[string]string{
		"q": options.Query,
	}

	if options.Page > 0 {
		params["page"] = strconv.Itoa(options.Page)
	}

	if options.PerPage > 0 {
		params["per_page"] = strconv.Itoa(options.PerPage)
	}

	fullURL, err := common.BuildURL(url, params)
	if err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest(fullURL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var searchResp common.GitHubSearchUsersResponse
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}
