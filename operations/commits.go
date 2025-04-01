package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// ListCommitsOptions defines options for listing commits
type ListCommitsOptions struct {
	Owner   string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo    string `json:"repo" jsonschema:"description=The name of the repository to list commits from"`
	Branch  string `json:"branch,omitempty" jsonschema:"description=The branch name or commit SHA to list commits from. Default: the repository's default branch"`
	Path    string `json:"path,omitempty" jsonschema:"description=Only commits containing changes to this file path will be returned"`
	Since   string `json:"since,omitempty" jsonschema:"description=Only commits after this date will be returned. ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ"`
	Until   string `json:"until,omitempty" jsonschema:"description=Only commits before this date will be returned. ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ"`
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
}

// Validate validates the ListCommitsOptions
func (o *ListCommitsOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Branch != "" {
		if _, err := common.ValidateBranchName(o.Branch); err != nil {
			return err
		}
	}
	return nil
}

// ListCommits lists commits in a GitHub repository
func ListCommits(options *ListCommitsOptions, apiReqs *common.APIRequirements) ([]common.GitHubCommit, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", options.Owner, options.Repo)

	params := make(map[string]string)
	if options.Branch != "" {
		params["sha"] = options.Branch
	}
	if options.Path != "" {
		params["path"] = options.Path
	}
	if options.Since != "" {
		params["since"] = options.Since
	}
	if options.Until != "" {
		params["until"] = options.Until
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

	var commits []common.GitHubCommit
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &commits); err != nil {
		return nil, err
	}

	return commits, nil
}
