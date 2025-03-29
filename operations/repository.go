package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// CreateRepositoryOptions defines the options for creating a repository
type CreateRepositoryOptions struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
	AutoInit    bool   `json:"auto_init,omitempty"`
}

// Validate validates the CreateRepositoryOptions
func (o *CreateRepositoryOptions) Validate() error {
	_, err := common.ValidateRepositoryName(o.Name)
	return err
}

// SearchRepositoriesOptions defines the options for searching repositories
type SearchRepositoriesOptions struct {
	Query   string `json:"query"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
}

// Validate validates the SearchRepositoriesOptions
func (o *SearchRepositoriesOptions) Validate() error {
	if o.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}

// ForkRepositoryOptions defines the options for forking a repository
type ForkRepositoryOptions struct {
	Owner        string `json:"owner"`
	Repo         string `json:"repo"`
	Organization string `json:"organization,omitempty"`
}

// Validate validates the ForkRepositoryOptions
func (o *ForkRepositoryOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Organization != "" {
		if _, err := common.ValidateOwnerName(o.Organization); err != nil {
			return err
		}
	}
	return nil
}

// CreateRepository creates a new GitHub repository
func CreateRepository(options *CreateRepositoryOptions, apiReqs *common.APIRequirements) (*common.GitHubRepository, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest("https://api.github.com/user/repos", "POST", options, apiReqs)
	if err != nil {
		return nil, err
	}

	var repo common.GitHubRepository
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

// SearchRepositories searches for GitHub repositories
func SearchRepositories(options *SearchRepositoriesOptions, apiReqs *common.APIRequirements) (*common.GitHubSearchResponse, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	if options.Page <= 0 {
		options.Page = 1
	}
	if options.PerPage <= 0 || options.PerPage > 100 {
		options.PerPage = 30
	}

	params := map[string]string{
		"q":        options.Query,
		"page":     strconv.Itoa(options.Page),
		"per_page": strconv.Itoa(options.PerPage),
	}

	url, err := common.BuildURL("https://api.github.com/search/repositories", params)
	if err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
	if err != nil {
		return nil, err
	}

	var searchResp common.GitHubSearchResponse
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// ForkRepository forks a GitHub repository
func ForkRepository(options *ForkRepositoryOptions, apiReqs *common.APIRequirements) (*common.GitHubRepository, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/forks", options.Owner, options.Repo)
	if options.Organization != "" {
		params := map[string]string{
			"organization": options.Organization,
		}
		var err error
		url, err = common.BuildURL(url, params)
		if err != nil {
			return nil, err
		}
	}

	resp, err := common.GitHubRequest(url, "POST", nil, apiReqs)
	if err != nil {
		return nil, err
	}

	var repo common.GitHubRepository
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}
