package operations

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/metoro-io/github-mcp-server-go/common"
)

// GetFileContentsOptions defines options for getting file contents
type GetFileContentsOptions struct {
	Owner string `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo  string `json:"repo" jsonschema:"description=The name of the repository containing the file"`
	Path  string `json:"path" jsonschema:"description=The path to the file or directory within the repository"`
	Ref   string `json:"ref,omitempty" jsonschema:"description=The name of the commit/branch/tag. Default: the repository's default branch (usually main)"`
}

// Validate validates the GetFileContentsOptions
func (o *GetFileContentsOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Path == "" {
		return fmt.Errorf("path is required")
	}
	return nil
}

// CreateOrUpdateFileOptions defines options for creating or updating a file
type CreateOrUpdateFileOptions struct {
	Owner     string         `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo      string         `json:"repo" jsonschema:"description=The name of the repository where the file will be created or updated"`
	Path      string         `json:"path" jsonschema:"description=The path to the file within the repository"`
	Message   string         `json:"message" jsonschema:"description=The commit message for the file creation or update"`
	Content   string         `json:"content" jsonschema:"description=The new content of the file as a string"`
	Branch    string         `json:"branch,omitempty" jsonschema:"description=The branch name to commit to. Default: the repository's default branch (usually main)"`
	SHA       string         `json:"sha,omitempty" jsonschema:"description=The blob SHA of the file being replaced if updating an existing file"`
	Committer *CommitterInfo `json:"committer,omitempty" jsonschema:"description=Information about the committer. If omitted the authenticated user's information is used"`
	Author    *CommitterInfo `json:"author,omitempty" jsonschema:"description=Information about the author. If omitted the committer information is used"`
}

// CommitterInfo represents author/committer information
type CommitterInfo struct {
	Name  string `json:"name" jsonschema:"description=The name of the author or committer"`
	Email string `json:"email" jsonschema:"description=The email of the author or committer"`
}

// Validate validates the CreateOrUpdateFileOptions
func (o *CreateOrUpdateFileOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if o.Path == "" {
		return fmt.Errorf("path is required")
	}
	if o.Message == "" {
		return fmt.Errorf("commit message is required")
	}
	if o.Content == "" {
		return fmt.Errorf("content is required")
	}
	if o.Branch != "" {
		if _, err := common.ValidateBranchName(o.Branch); err != nil {
			return err
		}
	}
	return nil
}

// PushFilesOptions defines options for pushing multiple files
type PushFilesOptions struct {
	Owner   string               `json:"owner" jsonschema:"description=The username or organization name that owns the repository"`
	Repo    string               `json:"repo" jsonschema:"description=The name of the repository where the files will be pushed"`
	Branch  string               `json:"branch" jsonschema:"description=The branch name to commit to"`
	Message string               `json:"message" jsonschema:"description=The commit message for this push operation"`
	Files   []PushFileDefinition `json:"files" jsonschema:"description=Array of files to create update or delete in this push operation"`
	BaseSHA string               `json:"base_sha,omitempty" jsonschema:"description=The SHA of the base commit to apply changes to. Default: latest commit on the specified branch"`
}

// PushFileDefinition represents a file to push
type PushFileDefinition struct {
	Path    string `json:"path" jsonschema:"description=The path to the file within the repository"`
	Content string `json:"content" jsonschema:"description=The content of the file as a string. Required unless Delete is true"`
	Delete  bool   `json:"delete,omitempty" jsonschema:"description=Whether to delete this file. If true Content is not required"`
}

// Validate validates the PushFilesOptions
func (o *PushFilesOptions) Validate() error {
	if _, err := common.ValidateOwnerName(o.Owner); err != nil {
		return err
	}
	if _, err := common.ValidateRepositoryName(o.Repo); err != nil {
		return err
	}
	if _, err := common.ValidateBranchName(o.Branch); err != nil {
		return err
	}
	if o.Message == "" {
		return fmt.Errorf("commit message is required")
	}
	if len(o.Files) == 0 {
		return fmt.Errorf("at least one file is required")
	}
	for i, file := range o.Files {
		if file.Path == "" {
			return fmt.Errorf("path is required for file at index %d", i)
		}
		if !file.Delete && file.Content == "" {
			return fmt.Errorf("content is required for non-deleted file at index %d", i)
		}
	}
	return nil
}

// GetFileContents gets the contents of a file from a GitHub repository
func GetFileContents(options *GetFileContentsOptions, apiReqs *common.APIRequirements) (interface{}, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", options.Owner, options.Repo, options.Path)
	if options.Ref != "" {
		params := map[string]string{
			"ref": options.Ref,
		}
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

	// Handle both single file and directory responses
	switch content := resp.(type) {
	case []interface{}:
		// This is a directory
		var fileList []common.FileContent
		jsonData, err := json.Marshal(content)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonData, &fileList); err != nil {
			return nil, err
		}
		return fileList, nil
	case map[string]interface{}:
		// This is a file
		var fileContent common.FileContent
		jsonData, err := json.Marshal(content)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonData, &fileContent); err != nil {
			return nil, err
		}

		// If the file is binary, just return it as is
		if fileContent.Encoding != "base64" {
			return fileContent, nil
		}

		// Decode the content if it's base64 encoded
		decodedContent, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(fileContent.Content, "\n", ""))
		if err != nil {
			return nil, fmt.Errorf("error decoding base64 content: %w", err)
		}
		fileContent.Content = string(decodedContent)
		return fileContent, nil
	default:
		return nil, fmt.Errorf("unexpected response type: %T", resp)
	}
}

// CreateOrUpdateFile creates or updates a file in a GitHub repository
func CreateOrUpdateFile(options *CreateOrUpdateFileOptions, apiReqs *common.APIRequirements) (*common.FileContent, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	// First, check if the file exists to get its SHA (for update)
	if options.SHA == "" {
		getOptions := &GetFileContentsOptions{
			Owner: options.Owner,
			Repo:  options.Repo,
			Path:  options.Path,
			Ref:   options.Branch,
		}
		existingFile, err := GetFileContents(getOptions, apiReqs)
		if err == nil {
			// File exists, get its SHA
			if fileContent, ok := existingFile.(common.FileContent); ok {
				options.SHA = fileContent.SHA
			}
		}
		// If the file doesn't exist, that's fine - we'll create it
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", options.Owner, options.Repo, options.Path)

	// Encode the content as base64
	content := base64.StdEncoding.EncodeToString([]byte(options.Content))

	requestBody := map[string]interface{}{
		"message": options.Message,
		"content": content,
	}

	if options.Branch != "" {
		requestBody["branch"] = options.Branch
	}

	if options.SHA != "" {
		requestBody["sha"] = options.SHA
	}

	if options.Committer != nil {
		requestBody["committer"] = options.Committer
	}

	if options.Author != nil {
		requestBody["author"] = options.Author
	}

	resp, err := common.GitHubRequest(url, "PUT", requestBody, apiReqs)
	if err != nil {
		return nil, err
	}

	respMap, ok := resp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", resp)
	}

	contentObj, ok := respMap["content"]
	if !ok {
		return nil, fmt.Errorf("content not found in response")
	}

	var fileContent common.FileContent
	jsonData, err := json.Marshal(contentObj)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonData, &fileContent); err != nil {
		return nil, err
	}

	return &fileContent, nil
}

// PushFiles pushes multiple files to a GitHub repository in a single commit
func PushFiles(options *PushFilesOptions, apiReqs *common.APIRequirements) (interface{}, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	// First, get the latest commit SHA for the branch
	baseSHA := options.BaseSHA
	if baseSHA == "" {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s",
			options.Owner, options.Repo, options.Branch)
		resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
		if err != nil {
			return nil, fmt.Errorf("error getting branch reference: %w", err)
		}

		respMap, ok := resp.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected response type: %T", resp)
		}

		objectMap, ok := respMap["object"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("object not found in response")
		}

		sha, ok := objectMap["sha"].(string)
		if !ok {
			return nil, fmt.Errorf("sha not found in response")
		}
		baseSHA = sha
	}

	// Get the base tree
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s",
		options.Owner, options.Repo, baseSHA)
	resp, err := common.GitHubRequest(url, "GET", nil, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error getting commit: %w", err)
	}

	respMap, ok := resp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", resp)
	}

	treeMap, ok := respMap["tree"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("tree not found in response")
	}

	baseTreeSHA, ok := treeMap["sha"].(string)
	if !ok {
		return nil, fmt.Errorf("tree sha not found in response")
	}

	// Create a new tree with the changes
	var treeItems []map[string]interface{}
	for _, file := range options.Files {
		if file.Delete {
			// To delete a file, we omit the content and set the sha to null
			treeItems = append(treeItems, map[string]interface{}{
				"path": file.Path,
				"mode": "100644",
				"type": "blob",
				"sha":  nil,
			})
		} else {
			// For new or updated files, we include the content
			treeItems = append(treeItems, map[string]interface{}{
				"path":    file.Path,
				"mode":    "100644",
				"type":    "blob",
				"content": file.Content,
			})
		}
	}

	// Create a tree
	createTreeURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees",
		options.Owner, options.Repo)
	createTreeBody := map[string]interface{}{
		"base_tree": baseTreeSHA,
		"tree":      treeItems,
	}

	treeResp, err := common.GitHubRequest(createTreeURL, "POST", createTreeBody, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error creating tree: %w", err)
	}

	treeRespMap, ok := treeResp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected tree response type: %T", treeResp)
	}

	newTreeSHA, ok := treeRespMap["sha"].(string)
	if !ok {
		return nil, fmt.Errorf("new tree sha not found in response")
	}

	// Create a commit
	createCommitURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits",
		options.Owner, options.Repo)
	createCommitBody := map[string]interface{}{
		"message": options.Message,
		"tree":    newTreeSHA,
		"parents": []string{baseSHA},
	}

	commitResp, err := common.GitHubRequest(createCommitURL, "POST", createCommitBody, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error creating commit: %w", err)
	}

	commitRespMap, ok := commitResp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected commit response type: %T", commitResp)
	}

	newCommitSHA, ok := commitRespMap["sha"].(string)
	if !ok {
		return nil, fmt.Errorf("new commit sha not found in response")
	}

	// Update the reference
	updateRefURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s",
		options.Owner, options.Repo, options.Branch)
	updateRefBody := map[string]interface{}{
		"sha": newCommitSHA,
	}

	_, err = common.GitHubRequest(updateRefURL, "PATCH", updateRefBody, apiReqs)
	if err != nil {
		return nil, fmt.Errorf("error updating reference: %w", err)
	}

	// Return the commit data
	return commitResp, nil
}
