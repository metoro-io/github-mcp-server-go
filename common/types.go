package common

import "time"

// GitHubUser represents a GitHub user or organization
type GitHubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	HTMLURL           string `json:"html_url"`
	GravatarID        string `json:"gravatar_id"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	PrivateGists      int    `json:"private_gists"`
	TotalPrivateRepos int    `json:"total_private_repos"`
	OwnedPrivateRepos int    `json:"owned_private_repos"`
	DiskUsage         int    `json:"disk_usage"`
	Collaborators     int    `json:"collaborators"`
}

// GitHubRepository represents a GitHub repository
type GitHubRepository struct {
	ID               int               `json:"id"`
	NodeID           string            `json:"node_id"`
	Name             string            `json:"name"`
	FullName         string            `json:"full_name"`
	Private          bool              `json:"private"`
	Owner            GitHubUser        `json:"owner"`
	HTMLURL          string            `json:"html_url"`
	Description      string            `json:"description"`
	Fork             bool              `json:"fork"`
	URL              string            `json:"url"`
	ForksURL         string            `json:"forks_url"`
	KeysURL          string            `json:"keys_url"`
	CollaboratorsURL string            `json:"collaborators_url"`
	TeamsURL         string            `json:"teams_url"`
	HooksURL         string            `json:"hooks_url"`
	IssueEventsURL   string            `json:"issue_events_url"`
	EventsURL        string            `json:"events_url"`
	AssigneesURL     string            `json:"assignees_url"`
	BranchesURL      string            `json:"branches_url"`
	TagsURL          string            `json:"tags_url"`
	BlobsURL         string            `json:"blobs_url"`
	GitTagsURL       string            `json:"git_tags_url"`
	GitRefsURL       string            `json:"git_refs_url"`
	TreesURL         string            `json:"trees_url"`
	StatusesURL      string            `json:"statuses_url"`
	LanguagesURL     string            `json:"languages_url"`
	StargazersURL    string            `json:"stargazers_url"`
	ContributorsURL  string            `json:"contributors_url"`
	SubscribersURL   string            `json:"subscribers_url"`
	SubscriptionURL  string            `json:"subscription_url"`
	CommitsURL       string            `json:"commits_url"`
	GitCommitsURL    string            `json:"git_commits_url"`
	CommentsURL      string            `json:"comments_url"`
	IssueCommentURL  string            `json:"issue_comment_url"`
	ContentsURL      string            `json:"contents_url"`
	CompareURL       string            `json:"compare_url"`
	MergesURL        string            `json:"merges_url"`
	ArchiveURL       string            `json:"archive_url"`
	DownloadsURL     string            `json:"downloads_url"`
	IssuesURL        string            `json:"issues_url"`
	PullsURL         string            `json:"pulls_url"`
	MilestonesURL    string            `json:"milestones_url"`
	NotificationsURL string            `json:"notifications_url"`
	LabelsURL        string            `json:"labels_url"`
	ReleasesURL      string            `json:"releases_url"`
	DeploymentsURL   string            `json:"deployments_url"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	PushedAt         time.Time         `json:"pushed_at"`
	GitURL           string            `json:"git_url"`
	SSHURL           string            `json:"ssh_url"`
	CloneURL         string            `json:"clone_url"`
	SvnURL           string            `json:"svn_url"`
	Homepage         string            `json:"homepage"`
	Size             int               `json:"size"`
	StargazersCount  int               `json:"stargazers_count"`
	WatchersCount    int               `json:"watchers_count"`
	Language         string            `json:"language"`
	HasIssues        bool              `json:"has_issues"`
	HasProjects      bool              `json:"has_projects"`
	HasDownloads     bool              `json:"has_downloads"`
	HasWiki          bool              `json:"has_wiki"`
	HasPages         bool              `json:"has_pages"`
	ForksCount       int               `json:"forks_count"`
	Archived         bool              `json:"archived"`
	Disabled         bool              `json:"disabled"`
	OpenIssuesCount  int               `json:"open_issues_count"`
	License          interface{}       `json:"license"`
	Forks            int               `json:"forks"`
	OpenIssues       int               `json:"open_issues"`
	Watchers         int               `json:"watchers"`
	DefaultBranch    string            `json:"default_branch"`
	Permissions      *Permissions      `json:"permissions,omitempty"`
	Parent           *GitHubRepository `json:"parent,omitempty"`
	Source           *GitHubRepository `json:"source,omitempty"`
}

// Permissions represents repository permissions
type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

// GitHubBranch represents a branch in a GitHub repository
type GitHubBranch struct {
	Name      string    `json:"name"`
	Commit    CommitRef `json:"commit"`
	Protected bool      `json:"protected"`
}

// CommitRef represents a reference to a commit
type CommitRef struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// FileContent represents content of a file in a GitHub repository
type FileContent struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	SHA         string `json:"sha"`
	URL         string `json:"url"`
	GitURL      string `json:"git_url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
}

// GitHubSearchResponse represents a search response from GitHub
type GitHubSearchResponse struct {
	TotalCount        int                `json:"total_count"`
	IncompleteResults bool               `json:"incomplete_results"`
	Items             []GitHubRepository `json:"items"`
}

// GitHubIssue represents an issue in a GitHub repository
type GitHubIssue struct {
	URL               string          `json:"url"`
	RepositoryURL     string          `json:"repository_url"`
	LabelsURL         string          `json:"labels_url"`
	CommentsURL       string          `json:"comments_url"`
	EventsURL         string          `json:"events_url"`
	HTMLURL           string          `json:"html_url"`
	ID                int             `json:"id"`
	NodeID            string          `json:"node_id"`
	Number            int             `json:"number"`
	Title             string          `json:"title"`
	User              GitHubUser      `json:"user"`
	Labels            []Label         `json:"labels"`
	State             string          `json:"state"`
	Locked            bool            `json:"locked"`
	Assignee          *GitHubUser     `json:"assignee"`
	Assignees         []GitHubUser    `json:"assignees"`
	Milestone         *Milestone      `json:"milestone"`
	Comments          int             `json:"comments"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	ClosedAt          *time.Time      `json:"closed_at"`
	AuthorAssociation string          `json:"author_association"`
	Body              string          `json:"body"`
	TimelineURL       string          `json:"timeline_url"`
	PullRequest       *PullRequestRef `json:"pull_request,omitempty"`
}

// Label represents a label on a GitHub issue
type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

// Milestone represents a milestone in a GitHub repository
type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int        `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      GitHubUser `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	State        string     `json:"state"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DueOn        *time.Time `json:"due_on"`
	ClosedAt     *time.Time `json:"closed_at"`
}

// PullRequestRef is a reference to a pull request
type PullRequestRef struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}

// GitHubPullRequest represents a pull request in a GitHub repository
type GitHubPullRequest struct {
	URL                 string       `json:"url"`
	ID                  int          `json:"id"`
	NodeID              string       `json:"node_id"`
	HTMLURL             string       `json:"html_url"`
	DiffURL             string       `json:"diff_url"`
	PatchURL            string       `json:"patch_url"`
	IssueURL            string       `json:"issue_url"`
	Number              int          `json:"number"`
	State               string       `json:"state"`
	Locked              bool         `json:"locked"`
	Title               string       `json:"title"`
	User                GitHubUser   `json:"user"`
	Body                string       `json:"body"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
	ClosedAt            *time.Time   `json:"closed_at"`
	MergedAt            *time.Time   `json:"merged_at"`
	MergeCommitSHA      string       `json:"merge_commit_sha"`
	Assignee            *GitHubUser  `json:"assignee"`
	Assignees           []GitHubUser `json:"assignees"`
	RequestedReviewers  []GitHubUser `json:"requested_reviewers"`
	RequestedTeams      []Team       `json:"requested_teams"`
	Labels              []Label      `json:"labels"`
	Milestone           *Milestone   `json:"milestone"`
	Draft               bool         `json:"draft"`
	CommitsURL          string       `json:"commits_url"`
	ReviewCommentsURL   string       `json:"review_comments_url"`
	ReviewCommentURL    string       `json:"review_comment_url"`
	CommentsURL         string       `json:"comments_url"`
	StatusesURL         string       `json:"statuses_url"`
	Head                PRRef        `json:"head"`
	Base                PRRef        `json:"base"`
	AuthorAssociation   string       `json:"author_association"`
	AutoMerge           interface{}  `json:"auto_merge"`
	Merged              bool         `json:"merged"`
	Mergeable           bool         `json:"mergeable"`
	Rebaseable          bool         `json:"rebaseable"`
	MergeableState      string       `json:"mergeable_state"`
	Comments            int          `json:"comments"`
	ReviewComments      int          `json:"review_comments"`
	MaintainerCanModify bool         `json:"maintainer_can_modify"`
	Commits             int          `json:"commits"`
	Additions           int          `json:"additions"`
	Deletions           int          `json:"deletions"`
	ChangedFiles        int          `json:"changed_files"`
}

// Team represents a team in a GitHub organization
type Team struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
}

// PRRef represents a reference to a branch in a pull request
type PRRef struct {
	Label string           `json:"label"`
	Ref   string           `json:"ref"`
	SHA   string           `json:"sha"`
	User  GitHubUser       `json:"user"`
	Repo  GitHubRepository `json:"repo"`
}

// GitHubCommit represents a commit in a GitHub repository
type GitHubCommit struct {
	SHA         string      `json:"sha"`
	NodeID      string      `json:"node_id"`
	Commit      CommitData  `json:"commit"`
	URL         string      `json:"url"`
	HTMLURL     string      `json:"html_url"`
	CommentsURL string      `json:"comments_url"`
	Author      *GitHubUser `json:"author"`
	Committer   *GitHubUser `json:"committer"`
	Parents     []CommitRef `json:"parents"`
}

// CommitData represents commit data
type CommitData struct {
	Author       CommitAuthor `json:"author"`
	Committer    CommitAuthor `json:"committer"`
	Message      string       `json:"message"`
	Tree         CommitRef    `json:"tree"`
	URL          string       `json:"url"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}

// CommitAuthor represents the author or committer of a commit
type CommitAuthor struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

// Verification represents the verification status of a commit
type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

// GitHubSearchCodeResponse represents a code search response from GitHub
type GitHubSearchCodeResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []CodeResult `json:"items"`
}

// CodeResult represents a single code search result
type CodeResult struct {
	Name       string           `json:"name"`
	Path       string           `json:"path"`
	SHA        string           `json:"sha"`
	URL        string           `json:"url"`
	GitURL     string           `json:"git_url"`
	HTMLURL    string           `json:"html_url"`
	Repository GitHubRepository `json:"repository"`
}

// GitHubSearchIssuesResponse represents an issue search response from GitHub
type GitHubSearchIssuesResponse struct {
	TotalCount        int           `json:"total_count"`
	IncompleteResults bool          `json:"incomplete_results"`
	Items             []GitHubIssue `json:"items"`
}

// GitHubSearchUsersResponse represents a user search response from GitHub
type GitHubSearchUsersResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []GitHubUser `json:"items"`
}
