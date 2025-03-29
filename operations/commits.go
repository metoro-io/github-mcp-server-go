package operations

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// ListCommitsOptions defines options for listing commits
type ListCommitsOptions struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Branch  string `json:"branch,omitempty"`
	Path    string `json:"path,omitempty"`
	Since   string `json:"since,omitempty"` // ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	Until   string `json:"until,omitempty"` // ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
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
