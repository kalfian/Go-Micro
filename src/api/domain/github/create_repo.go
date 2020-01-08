package github

/*
{
	"name": "Hello-World",
	"description": "Thi is furst repo",
	"homepage": "https://github.com",
	"private": false,
	"has_issues": true,
	"has_projects": true,
	"has_wiki": true
}
*/

//CreateRepoRequest ... parameter to create repos
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// ResponseCreateRepo ... Response when repo created from server
type ResponseCreateRepo struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Fullname    string          `json:"fullname"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

// RepoOwner ... some response from ResponseCreateRepo
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermissions ... some response from ResponseCreateRepo
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"pull"`
	HasPush bool `json:"push"`
}
