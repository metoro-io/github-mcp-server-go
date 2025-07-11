package operations

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// GetTagsOptions defines the options for getting repository tags
type GetTagsOptions struct {
	Owner   string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo    string `json:"repo" jsonschema:"description=The name of the repository"`
	Search  string `json:"search,omitempty" jsonschema:"description=Optional search query to filter tags. Performs case-insensitive fuzzy matching on tag names"`
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number of the results to fetch. Default: 1"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of results per page. Default: 30. Maximum: 100"`
}

// Validate validates the GetTagsOptions
func (o *GetTagsOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	return nil
}

// GetTags fetches all tags for a GitHub repository
func GetTags(options *GetTagsOptions, apiReqs *common.APIRequirements) ([]string, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/tags", options.Owner, options.Repo)

	params := make(map[string]string)
	if options.Page > 0 {
		params["page"] = fmt.Sprintf("%d", options.Page)
	}
	if options.PerPage > 0 {
		params["per_page"] = fmt.Sprintf("%d", options.PerPage)
	}

	fullURL, err := common.BuildURL(url, params)
	if err != nil {
		return nil, err
	}

	resp, err := common.GitHubRequest(fullURL, "GET", nil, apiReqs)
	if err != nil {
		return nil, err
	}

	var refs []common.GitHubRef
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &refs); err != nil {
		return nil, err
	}

	refToRet := make([]string, 0, len(refs))
	for _, ref := range refs {
		// Extract tag name from ref (e.g., "refs/tags/v1.0.0" -> "v1.0.0")
		tagName := strings.TrimPrefix(ref.Ref, "refs/tags/")

		// Apply search filter if provided
		if options.Search != "" {
			// Case-insensitive fuzzy matching
			if !strings.Contains(strings.ToLower(tagName), strings.ToLower(options.Search)) {
				continue
			}
		}

		refToRet = append(refToRet, tagName)
	}

	return refToRet, nil
}
