package operations

import (
	"encoding/json"
	"fmt"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// CreateBranchOptions defines the options for creating a branch
type CreateBranchOptions struct {
	Owner      string `json:"owner"`
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	FromBranch string `json:"from_branch"`
}

// Validate validates the CreateBranchOptions
func (o *CreateBranchOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if _, err := common.ValidateBranchName(o.Branch); err != nil {
		return err
	}
	if _, err := common.ValidateBranchName(o.FromBranch); err != nil {
		return err
	}
	return nil
}

// CreateBranchFromRef creates a new branch in a GitHub repository
func CreateBranchFromRef(options *CreateBranchOptions, apiReqs *common.APIRequirements) (*common.GitHubBranch, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	// First get the source branch to get the SHA
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches/%s", options.Owner, options.Repo, options.FromBranch)
	resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error getting source branch: %w", err)
	}

	var sourceBranch common.GitHubBranch
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &sourceBranch); err != nil {
		return nil, err
	}

	// Now create the new branch as a reference
	refURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs", options.Owner, options.Repo)
	refData := map[string]string{
		"ref": fmt.Sprintf("refs/heads/%s", options.Branch),
		"sha": sourceBranch.Commit.SHA,
	}

	_, err = common.GitHubRequest(refURL, "POST", refData, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error creating branch: %w", err)
	}

	// Verify branch was created
	branchURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches/%s", options.Owner, options.Repo, options.Branch)
	branchResp, err := common.GitHubRequest(branchURL, "GET", nil, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("branch might have been created but verification failed: %w", err)
	}

	var newBranch common.GitHubBranch
	jsonData, err = json.Marshal(branchResp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &newBranch); err != nil {
		return nil, err
	}

	return &newBranch, nil
}
