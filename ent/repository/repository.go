// Code generated by ent, DO NOT EDIT.

package repository

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the repository type in the database.
	Label = "repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldPrivate holds the string denoting the private field in the database.
	FieldPrivate = "private"
	// FieldHTMLURL holds the string denoting the html_url field in the database.
	FieldHTMLURL = "html_url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFork holds the string denoting the fork field in the database.
	FieldFork = "fork"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldArchiveURL holds the string denoting the archive_url field in the database.
	FieldArchiveURL = "archive_url"
	// FieldAssigneesURL holds the string denoting the assignees_url field in the database.
	FieldAssigneesURL = "assignees_url"
	// FieldBlobsURL holds the string denoting the blobs_url field in the database.
	FieldBlobsURL = "blobs_url"
	// FieldBranchesURL holds the string denoting the branches_url field in the database.
	FieldBranchesURL = "branches_url"
	// FieldCollaboratorsURL holds the string denoting the collaborators_url field in the database.
	FieldCollaboratorsURL = "collaborators_url"
	// FieldCommentsURL holds the string denoting the comments_url field in the database.
	FieldCommentsURL = "comments_url"
	// FieldCommitsURL holds the string denoting the commits_url field in the database.
	FieldCommitsURL = "commits_url"
	// FieldCompareURL holds the string denoting the compare_url field in the database.
	FieldCompareURL = "compare_url"
	// FieldContentsURL holds the string denoting the contents_url field in the database.
	FieldContentsURL = "contents_url"
	// FieldContributorsURL holds the string denoting the contributors_url field in the database.
	FieldContributorsURL = "contributors_url"
	// FieldDeploymentsURL holds the string denoting the deployments_url field in the database.
	FieldDeploymentsURL = "deployments_url"
	// FieldDownloadsURL holds the string denoting the downloads_url field in the database.
	FieldDownloadsURL = "downloads_url"
	// FieldEventsURL holds the string denoting the events_url field in the database.
	FieldEventsURL = "events_url"
	// FieldForksURL holds the string denoting the forks_url field in the database.
	FieldForksURL = "forks_url"
	// FieldGitCommitsURL holds the string denoting the git_commits_url field in the database.
	FieldGitCommitsURL = "git_commits_url"
	// FieldGitRefsURL holds the string denoting the git_refs_url field in the database.
	FieldGitRefsURL = "git_refs_url"
	// FieldGitTagsURL holds the string denoting the git_tags_url field in the database.
	FieldGitTagsURL = "git_tags_url"
	// FieldGitURL holds the string denoting the git_url field in the database.
	FieldGitURL = "git_url"
	// FieldIssueCommentURL holds the string denoting the issue_comment_url field in the database.
	FieldIssueCommentURL = "issue_comment_url"
	// FieldIssueEventsURL holds the string denoting the issue_events_url field in the database.
	FieldIssueEventsURL = "issue_events_url"
	// FieldIssuesURL holds the string denoting the issues_url field in the database.
	FieldIssuesURL = "issues_url"
	// FieldKeysURL holds the string denoting the keys_url field in the database.
	FieldKeysURL = "keys_url"
	// FieldLabelsURL holds the string denoting the labels_url field in the database.
	FieldLabelsURL = "labels_url"
	// FieldLanguagesURL holds the string denoting the languages_url field in the database.
	FieldLanguagesURL = "languages_url"
	// FieldMergesURL holds the string denoting the merges_url field in the database.
	FieldMergesURL = "merges_url"
	// FieldMilestonesURL holds the string denoting the milestones_url field in the database.
	FieldMilestonesURL = "milestones_url"
	// FieldNotificationsURL holds the string denoting the notifications_url field in the database.
	FieldNotificationsURL = "notifications_url"
	// FieldPullsURL holds the string denoting the pulls_url field in the database.
	FieldPullsURL = "pulls_url"
	// FieldReleasesURL holds the string denoting the releases_url field in the database.
	FieldReleasesURL = "releases_url"
	// FieldSSHURL holds the string denoting the ssh_url field in the database.
	FieldSSHURL = "ssh_url"
	// FieldStargazersURL holds the string denoting the stargazers_url field in the database.
	FieldStargazersURL = "stargazers_url"
	// FieldStatusesURL holds the string denoting the statuses_url field in the database.
	FieldStatusesURL = "statuses_url"
	// FieldSubscribersURL holds the string denoting the subscribers_url field in the database.
	FieldSubscribersURL = "subscribers_url"
	// FieldSubscriptionURL holds the string denoting the subscription_url field in the database.
	FieldSubscriptionURL = "subscription_url"
	// FieldTagsURL holds the string denoting the tags_url field in the database.
	FieldTagsURL = "tags_url"
	// FieldTeamsURL holds the string denoting the teams_url field in the database.
	FieldTeamsURL = "teams_url"
	// FieldTreesURL holds the string denoting the trees_url field in the database.
	FieldTreesURL = "trees_url"
	// FieldCloneURL holds the string denoting the clone_url field in the database.
	FieldCloneURL = "clone_url"
	// FieldMirrorURL holds the string denoting the mirror_url field in the database.
	FieldMirrorURL = "mirror_url"
	// FieldHooksURL holds the string denoting the hooks_url field in the database.
	FieldHooksURL = "hooks_url"
	// FieldSvnURL holds the string denoting the svn_url field in the database.
	FieldSvnURL = "svn_url"
	// FieldHomepage holds the string denoting the homepage field in the database.
	FieldHomepage = "homepage"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldForksCount holds the string denoting the forks_count field in the database.
	FieldForksCount = "forks_count"
	// FieldStargazersCount holds the string denoting the stargazers_count field in the database.
	FieldStargazersCount = "stargazers_count"
	// FieldWatchersCount holds the string denoting the watchers_count field in the database.
	FieldWatchersCount = "watchers_count"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldDefaultBranch holds the string denoting the default_branch field in the database.
	FieldDefaultBranch = "default_branch"
	// FieldOpenIssuesCount holds the string denoting the open_issues_count field in the database.
	FieldOpenIssuesCount = "open_issues_count"
	// FieldIsTemplate holds the string denoting the is_template field in the database.
	FieldIsTemplate = "is_template"
	// FieldTopics holds the string denoting the topics field in the database.
	FieldTopics = "topics"
	// FieldHasIssuesEnabled holds the string denoting the has_issues_enabled field in the database.
	FieldHasIssuesEnabled = "has_issues_enabled"
	// FieldHasProjects holds the string denoting the has_projects field in the database.
	FieldHasProjects = "has_projects"
	// FieldHasWiki holds the string denoting the has_wiki field in the database.
	FieldHasWiki = "has_wiki"
	// FieldHasPages holds the string denoting the has_pages field in the database.
	FieldHasPages = "has_pages"
	// FieldHasDownloads holds the string denoting the has_downloads field in the database.
	FieldHasDownloads = "has_downloads"
	// FieldHasDiscussions holds the string denoting the has_discussions field in the database.
	FieldHasDiscussions = "has_discussions"
	// FieldArchived holds the string denoting the archived field in the database.
	FieldArchived = "archived"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldPushedAt holds the string denoting the pushed_at field in the database.
	FieldPushedAt = "pushed_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSubscribersCount holds the string denoting the subscribers_count field in the database.
	FieldSubscribersCount = "subscribers_count"
	// FieldNetworkCount holds the string denoting the network_count field in the database.
	FieldNetworkCount = "network_count"
	// FieldForks holds the string denoting the forks field in the database.
	FieldForks = "forks"
	// FieldOpenIssues holds the string denoting the open_issues field in the database.
	FieldOpenIssues = "open_issues"
	// FieldWatchers holds the string denoting the watchers field in the database.
	FieldWatchers = "watchers"
	// FieldLicense holds the string denoting the license field in the database.
	FieldLicense = "license"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeIssues holds the string denoting the issues edge name in mutations.
	EdgeIssues = "issues"
	// EdgePullRequests holds the string denoting the pull_requests edge name in mutations.
	EdgePullRequests = "pull_requests"
	// Table holds the table name of the repository in the database.
	Table = "repositories"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "repositories"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_repositories"
	// IssuesTable is the table that holds the issues relation/edge.
	IssuesTable = "issues"
	// IssuesInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesInverseTable = "issues"
	// IssuesColumn is the table column denoting the issues relation/edge.
	IssuesColumn = "repository_issues"
	// PullRequestsTable is the table that holds the pull_requests relation/edge.
	PullRequestsTable = "pull_requests"
	// PullRequestsInverseTable is the table name for the PullRequest entity.
	// It exists in this package in order to avoid circular dependency with the "pullrequest" package.
	PullRequestsInverseTable = "pull_requests"
	// PullRequestsColumn is the table column denoting the pull_requests relation/edge.
	PullRequestsColumn = "repository_pull_requests"
)

// Columns holds all SQL columns for repository fields.
var Columns = []string{
	FieldID,
	FieldNodeID,
	FieldName,
	FieldFullName,
	FieldPrivate,
	FieldHTMLURL,
	FieldDescription,
	FieldFork,
	FieldURL,
	FieldArchiveURL,
	FieldAssigneesURL,
	FieldBlobsURL,
	FieldBranchesURL,
	FieldCollaboratorsURL,
	FieldCommentsURL,
	FieldCommitsURL,
	FieldCompareURL,
	FieldContentsURL,
	FieldContributorsURL,
	FieldDeploymentsURL,
	FieldDownloadsURL,
	FieldEventsURL,
	FieldForksURL,
	FieldGitCommitsURL,
	FieldGitRefsURL,
	FieldGitTagsURL,
	FieldGitURL,
	FieldIssueCommentURL,
	FieldIssueEventsURL,
	FieldIssuesURL,
	FieldKeysURL,
	FieldLabelsURL,
	FieldLanguagesURL,
	FieldMergesURL,
	FieldMilestonesURL,
	FieldNotificationsURL,
	FieldPullsURL,
	FieldReleasesURL,
	FieldSSHURL,
	FieldStargazersURL,
	FieldStatusesURL,
	FieldSubscribersURL,
	FieldSubscriptionURL,
	FieldTagsURL,
	FieldTeamsURL,
	FieldTreesURL,
	FieldCloneURL,
	FieldMirrorURL,
	FieldHooksURL,
	FieldSvnURL,
	FieldHomepage,
	FieldLanguage,
	FieldForksCount,
	FieldStargazersCount,
	FieldWatchersCount,
	FieldSize,
	FieldDefaultBranch,
	FieldOpenIssuesCount,
	FieldIsTemplate,
	FieldTopics,
	FieldHasIssuesEnabled,
	FieldHasProjects,
	FieldHasWiki,
	FieldHasPages,
	FieldHasDownloads,
	FieldHasDiscussions,
	FieldArchived,
	FieldDisabled,
	FieldVisibility,
	FieldPushedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSubscribersCount,
	FieldNetworkCount,
	FieldForks,
	FieldOpenIssues,
	FieldWatchers,
	FieldLicense,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "repositories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_repositories",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// Visibility values.
const (
	VisibilityPublic   Visibility = "public"
	VisibilityPrivate  Visibility = "private"
	VisibilityInternal Visibility = "internal"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPublic, VisibilityPrivate, VisibilityInternal:
		return nil
	default:
		return fmt.Errorf("repository: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Repository queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByPrivate orders the results by the private field.
func ByPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivate, opts...).ToFunc()
}

// ByHTMLURL orders the results by the html_url field.
func ByHTMLURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLURL, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFork orders the results by the fork field.
func ByFork(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFork, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByArchiveURL orders the results by the archive_url field.
func ByArchiveURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchiveURL, opts...).ToFunc()
}

// ByAssigneesURL orders the results by the assignees_url field.
func ByAssigneesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAssigneesURL, opts...).ToFunc()
}

// ByBlobsURL orders the results by the blobs_url field.
func ByBlobsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlobsURL, opts...).ToFunc()
}

// ByBranchesURL orders the results by the branches_url field.
func ByBranchesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBranchesURL, opts...).ToFunc()
}

// ByCollaboratorsURL orders the results by the collaborators_url field.
func ByCollaboratorsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollaboratorsURL, opts...).ToFunc()
}

// ByCommentsURL orders the results by the comments_url field.
func ByCommentsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentsURL, opts...).ToFunc()
}

// ByCommitsURL orders the results by the commits_url field.
func ByCommitsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommitsURL, opts...).ToFunc()
}

// ByCompareURL orders the results by the compare_url field.
func ByCompareURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompareURL, opts...).ToFunc()
}

// ByContentsURL orders the results by the contents_url field.
func ByContentsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContentsURL, opts...).ToFunc()
}

// ByContributorsURL orders the results by the contributors_url field.
func ByContributorsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContributorsURL, opts...).ToFunc()
}

// ByDeploymentsURL orders the results by the deployments_url field.
func ByDeploymentsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeploymentsURL, opts...).ToFunc()
}

// ByDownloadsURL orders the results by the downloads_url field.
func ByDownloadsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadsURL, opts...).ToFunc()
}

// ByEventsURL orders the results by the events_url field.
func ByEventsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventsURL, opts...).ToFunc()
}

// ByForksURL orders the results by the forks_url field.
func ByForksURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForksURL, opts...).ToFunc()
}

// ByGitCommitsURL orders the results by the git_commits_url field.
func ByGitCommitsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGitCommitsURL, opts...).ToFunc()
}

// ByGitRefsURL orders the results by the git_refs_url field.
func ByGitRefsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGitRefsURL, opts...).ToFunc()
}

// ByGitTagsURL orders the results by the git_tags_url field.
func ByGitTagsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGitTagsURL, opts...).ToFunc()
}

// ByGitURL orders the results by the git_url field.
func ByGitURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGitURL, opts...).ToFunc()
}

// ByIssueCommentURL orders the results by the issue_comment_url field.
func ByIssueCommentURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssueCommentURL, opts...).ToFunc()
}

// ByIssueEventsURL orders the results by the issue_events_url field.
func ByIssueEventsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssueEventsURL, opts...).ToFunc()
}

// ByIssuesURL orders the results by the issues_url field.
func ByIssuesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssuesURL, opts...).ToFunc()
}

// ByKeysURL orders the results by the keys_url field.
func ByKeysURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeysURL, opts...).ToFunc()
}

// ByLabelsURL orders the results by the labels_url field.
func ByLabelsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLabelsURL, opts...).ToFunc()
}

// ByLanguagesURL orders the results by the languages_url field.
func ByLanguagesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguagesURL, opts...).ToFunc()
}

// ByMergesURL orders the results by the merges_url field.
func ByMergesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMergesURL, opts...).ToFunc()
}

// ByMilestonesURL orders the results by the milestones_url field.
func ByMilestonesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMilestonesURL, opts...).ToFunc()
}

// ByNotificationsURL orders the results by the notifications_url field.
func ByNotificationsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotificationsURL, opts...).ToFunc()
}

// ByPullsURL orders the results by the pulls_url field.
func ByPullsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPullsURL, opts...).ToFunc()
}

// ByReleasesURL orders the results by the releases_url field.
func ByReleasesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleasesURL, opts...).ToFunc()
}

// BySSHURL orders the results by the ssh_url field.
func BySSHURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSSHURL, opts...).ToFunc()
}

// ByStargazersURL orders the results by the stargazers_url field.
func ByStargazersURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStargazersURL, opts...).ToFunc()
}

// ByStatusesURL orders the results by the statuses_url field.
func ByStatusesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusesURL, opts...).ToFunc()
}

// BySubscribersURL orders the results by the subscribers_url field.
func BySubscribersURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscribersURL, opts...).ToFunc()
}

// BySubscriptionURL orders the results by the subscription_url field.
func BySubscriptionURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionURL, opts...).ToFunc()
}

// ByTagsURL orders the results by the tags_url field.
func ByTagsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagsURL, opts...).ToFunc()
}

// ByTeamsURL orders the results by the teams_url field.
func ByTeamsURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamsURL, opts...).ToFunc()
}

// ByTreesURL orders the results by the trees_url field.
func ByTreesURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTreesURL, opts...).ToFunc()
}

// ByCloneURL orders the results by the clone_url field.
func ByCloneURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCloneURL, opts...).ToFunc()
}

// ByMirrorURL orders the results by the mirror_url field.
func ByMirrorURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMirrorURL, opts...).ToFunc()
}

// ByHooksURL orders the results by the hooks_url field.
func ByHooksURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHooksURL, opts...).ToFunc()
}

// BySvnURL orders the results by the svn_url field.
func BySvnURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSvnURL, opts...).ToFunc()
}

// ByHomepage orders the results by the homepage field.
func ByHomepage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHomepage, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByForksCount orders the results by the forks_count field.
func ByForksCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForksCount, opts...).ToFunc()
}

// ByStargazersCount orders the results by the stargazers_count field.
func ByStargazersCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStargazersCount, opts...).ToFunc()
}

// ByWatchersCount orders the results by the watchers_count field.
func ByWatchersCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWatchersCount, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByDefaultBranch orders the results by the default_branch field.
func ByDefaultBranch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultBranch, opts...).ToFunc()
}

// ByOpenIssuesCount orders the results by the open_issues_count field.
func ByOpenIssuesCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpenIssuesCount, opts...).ToFunc()
}

// ByIsTemplate orders the results by the is_template field.
func ByIsTemplate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTemplate, opts...).ToFunc()
}

// ByHasIssuesEnabled orders the results by the has_issues_enabled field.
func ByHasIssuesEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasIssuesEnabled, opts...).ToFunc()
}

// ByHasProjects orders the results by the has_projects field.
func ByHasProjects(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasProjects, opts...).ToFunc()
}

// ByHasWiki orders the results by the has_wiki field.
func ByHasWiki(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasWiki, opts...).ToFunc()
}

// ByHasPages orders the results by the has_pages field.
func ByHasPages(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasPages, opts...).ToFunc()
}

// ByHasDownloads orders the results by the has_downloads field.
func ByHasDownloads(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasDownloads, opts...).ToFunc()
}

// ByHasDiscussions orders the results by the has_discussions field.
func ByHasDiscussions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasDiscussions, opts...).ToFunc()
}

// ByArchived orders the results by the archived field.
func ByArchived(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchived, opts...).ToFunc()
}

// ByDisabled orders the results by the disabled field.
func ByDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabled, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByPushedAt orders the results by the pushed_at field.
func ByPushedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPushedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySubscribersCount orders the results by the subscribers_count field.
func BySubscribersCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscribersCount, opts...).ToFunc()
}

// ByNetworkCount orders the results by the network_count field.
func ByNetworkCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetworkCount, opts...).ToFunc()
}

// ByForks orders the results by the forks field.
func ByForks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForks, opts...).ToFunc()
}

// ByOpenIssues orders the results by the open_issues field.
func ByOpenIssues(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpenIssues, opts...).ToFunc()
}

// ByWatchers orders the results by the watchers field.
func ByWatchers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWatchers, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByIssuesCount orders the results by issues count.
func ByIssuesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuesStep(), opts...)
	}
}

// ByIssues orders the results by issues terms.
func ByIssues(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPullRequestsCount orders the results by pull_requests count.
func ByPullRequestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPullRequestsStep(), opts...)
	}
}

// ByPullRequests orders the results by pull_requests terms.
func ByPullRequests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPullRequestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newIssuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IssuesTable, IssuesColumn),
	)
}
func newPullRequestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PullRequestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PullRequestsTable, PullRequestsColumn),
	)
}
