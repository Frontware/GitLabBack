package gitlab

import (
	"strconv"
	"time"
)

// Project basic GitLab project struct
type Project struct {
	ID                   int        `json:"id"`
	Description          string     `json:"description"`
	DefaultBranch        string     `json:"default_branch"`
	TagList              []string   `json:"tag_list"`
	Archived             bool       `json:"archived"`
	Visibility           string     `json:"visibility"`
	SSHURLToRepo         string     `json:"ssh_url_to_repo"`
	HTTPURLToRepo        string     `json:"http_url_to_repo"`
	Web                  string     `json:"web_url"`
	Name                 string     `json:"name"`
	NameWithNamespace    string     `json:"name_with_namespace"`
	Path                 string     `json:"path"`
	PathWithNamespace    string     `json:"path_with_namespace"`
	IssuesEnabled        bool       `json:"issues_enabled"`
	MergeRequestEnabled  bool       `json:"merge_request_enabled"`
	WikiEnabled          bool       `json:"wiki_enabled"`
	JobsEnabled          bool       `json:"jobs_enabled"`
	SnippetsEnabled      bool       `json:"snippets_enabled"`
	CreatedAt            *time.Time `json:"created_at"`
	LastActivityAt       *time.Time `jons:"last_activity_at"`
	SharedRunnersEnabled bool       `json:"shared_runners_enabled"`
	CreatorID            int        `json:"creator_id"`
	Namespace            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Path string `json:"path"`
		Kind string `json:"kind"`
	} `json:"namespace"`
	Avatar           string `json:"avatar"`
	Star             int    `json:"star_count"`
	Fork             int    `json:"fork_count"`
	OpenIssues       int    `json:"open_issues"`
	PublicJobs       bool   `json:"public_jobs"`
	SharedWithGroups []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	RequestAccessEnabled bool `json:"request_access_enabled"`
}

// ListProjects requests list of projects for a group
func (c *Client) ListProjects(id int) ([]Project, error) {
	strID := strconv.Itoa(id)
	req, err := c.NewRequest("GET", "groups/"+strID+"/projects")
	if err != nil {
		return nil, err
	}

	projects := []Project{}

	err = c.Do(req, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
