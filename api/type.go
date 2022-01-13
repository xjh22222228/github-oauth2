package api

import "time"

// Commit https://api.github.com/repos/xjh22222228/nav/commits?page=0&per_page=32229
type Commit struct {
	Sha    string `json:"sha"`
	NodeId string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string     `json:"name"`
			Email string     `json:"email"`
			Date  *time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			Url string `json:"url"`
		} `json:"tree"`
		Url          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	CommentsUrl string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		Url     string `json:"url"`
		HtmlUrl string `json:"html_url"`
	} `json:"parents"`
}

// UserRepo https://api.github.com/users/xjh22222228/repos
type UserRepo struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HtmlUrl          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ForksUrl         string      `json:"forks_url"`
	KeysUrl          string      `json:"keys_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	TeamsUrl         string      `json:"teams_url"`
	HooksUrl         string      `json:"hooks_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	EventsUrl        string      `json:"events_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BranchesUrl      string      `json:"branches_url"`
	TagsUrl          string      `json:"tags_url"`
	BlobsUrl         string      `json:"blobs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	TreesUrl         string      `json:"trees_url"`
	StatusesUrl      string      `json:"statuses_url"`
	LanguagesUrl     string      `json:"languages_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	CommitsUrl       string      `json:"commits_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	CommentsUrl      string      `json:"comments_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	ContentsUrl      string      `json:"contents_url"`
	CompareUrl       string      `json:"compare_url"`
	MergesUrl        string      `json:"merges_url"`
	ArchiveUrl       string      `json:"archive_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	IssuesUrl        string      `json:"issues_url"`
	PullsUrl         string      `json:"pulls_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	LabelsUrl        string      `json:"labels_url"`
	ReleasesUrl      string      `json:"releases_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitUrl           string      `json:"git_url"`
	SshUrl           string      `json:"ssh_url"`
	CloneUrl         string      `json:"clone_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorUrl        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxId string `json:"spdx_id"`
		Url    string `json:"url"`
		NodeId string `json:"node_id"`
	} `json:"license"`
	AllowForking  bool          `json:"allow_forking"`
	IsTemplate    bool          `json:"is_template"`
	Topics        []interface{} `json:"topics"`
	Visibility    string        `json:"visibility"`
	Forks         int           `json:"forks"`
	OpenIssues    int           `json:"open_issues"`
	Watchers      int           `json:"watchers"`
	DefaultBranch string        `json:"default_branch"`
}

type User struct {
	Login                   string      `json:"login"`
	Id                      int         `json:"id"`
	NodeId                  string      `json:"node_id"`
	AvatarUrl               string      `json:"avatar_url"`
	GravatarId              string      `json:"gravatar_id"`
	Url                     string      `json:"url"`
	HtmlUrl                 string      `json:"html_url"`
	FollowersUrl            string      `json:"followers_url"`
	FollowingUrl            string      `json:"following_url"`
	GistsUrl                string      `json:"gists_url"`
	StarredUrl              string      `json:"starred_url"`
	SubscriptionsUrl        string      `json:"subscriptions_url"`
	OrganizationsUrl        string      `json:"organizations_url"`
	ReposUrl                string      `json:"repos_url"`
	EventsUrl               string      `json:"events_url"`
	ReceivedEventsUrl       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 interface{} `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   string      `json:"email"`
	Hireable                interface{} `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         interface{} `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

type LikeCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type HourCount struct {
	Hour  int `json:"hour"`
	Count int `json:"count"`
}
