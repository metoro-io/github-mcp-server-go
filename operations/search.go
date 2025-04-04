package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// SearchCodeOptions defines options for searching code
type SearchCodeOptions struct {
	Query   string `json:"query" jsonschema:"description=The search query string. Format follows GitHub's code search syntax. Example: filename:.go extension:go"`
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
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
	Query   string `json:"query" jsonschema:"description=The search query string. Format follows GitHub's issue search syntax. Example: is:issue is:open label:bug"`
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
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
	Query   string `json:"query" jsonschema:"description=The search query string. Format follows GitHub's user search syntax. Example: type:user language:go location:japan"`
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
}

// Validate validates the SearchUsersOptions
func (o *SearchUsersOptions) Validate() error {
	if o.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

// SearchCode searches for code across GitHub repositories
func SearchCode(options *SearchCodeOptions, apiReqs *common.APIRequirements) (*common.GitHubSearchCodeResponse, error) {
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

	resp, err := common.GitHubRequest(fullURL, "GET", nil, apiReqs)
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
func SearchIssues(options *SearchIssuesOptions, apiReqs *common.APIRequirements) (*common.GitHubSearchIssuesResponse, error) {
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

	resp, err := common.GitHubRequest(fullURL, "GET", nil, apiReqs)
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
func SearchUsers(options *SearchUsersOptions, apiReqs *common.APIRequirements) (*common.GitHubSearchUsersResponse, error) {
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

	resp, err := common.GitHubRequest(fullURL, "GET", nil, apiReqs)
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
