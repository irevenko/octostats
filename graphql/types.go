package graphql

import "github.com/shurcooL/githubv4"

type contributions struct {
	TotalCount githubv4.Int
}

type repository struct {
	NameWithOwner   githubv4.String
	PrimaryLanguage struct {
		Name githubv4.String
	}
}

type commitContributions struct {
	Contributions contributions
	Repository    repository
}

type issueContributions struct {
	Contributions contributions
	Repository    repository
}

type pullRequestContributions struct {
	Contributions contributions
	Repository    repository
}

type pullRequestReviewContributions struct {
	Contributions contributions
	Repository    repository
}

type ContributionsCollection struct {
	CommitContributionsByRepository            []commitContributions
	IssueContributionsByRepository             []issueContributions
	PullRequestContributionsByRepository       []pullRequestContributions
	PullRequestReviewContributionsByRepository []pullRequestReviewContributions
}

type AggregatedContributionsCollection struct {
	Repository             string
	CommitCount            int
	IssueCount             int
	PullRequestCount       int
	PullRequestReviewCount int
}
