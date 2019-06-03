package gitlab

// Group basic GitLab group struct
type Group struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Path                  string `json:"path"`
	Description           string `json:"description"`
	Visibility            string `json:"visibility"`
	LFSEnable             bool   `json:"lfs_enabled"`
	Avatar                string `json:"avatar_url"`
	Web                   string `json:"web_url"`
	AccessEnable          bool   `json:"request_access_enabled"`
	FullName              string `json:"full_name"`
	FullPath              string `json:"full_path"`
	FileTemplateProjectID int    `json:"file_template_project_id"`
	ParentID              int    `json:"parent_id"`
}

// ListGroups requests list of GitLab groups
func (c *Client) ListGroups() ([]Group, error) {
	req, err := c.NewRequest("GET", "groups")
	if err != nil {
		return nil, err
	}

	groups := []Group{}

	err = c.Do(req, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
