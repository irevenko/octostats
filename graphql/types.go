package graphql

import "github.com/shurcooL/githubv4"

type Contributions struct {
	TotalCount githubv4.Int
}

type Repository struct {
	NameWithOwner   githubv4.String
	PrimaryLanguage struct {
		Name githubv4.String
	}
}

type CommitContributions struct {
	Contributions Contributions
	Repository    Repository
}

type IssueContributions struct {
	Contributions Contributions
	Repository    Repository
}

type PullRequestContributions struct {
	Contributions Contributions
	Repository    Repository
}

type pullRequestReviewContributions struct {
	Contributions Contributions
	Repository    Repository
}

type ContributionsCollection struct {
	CommitContributionsByRepository            []CommitContributions
	IssueContributionsByRepository             []IssueContributions
	PullRequestContributionsByRepository       []PullRequestContributions
	PullRequestReviewContributionsByRepository []pullRequestReviewContributions
}

type AggregatedContributionsCollection struct {
	Repository             string
	CommitCount            int
	IssueCount             int
	PullRequestCount       int
	PullRequestReviewCount int
}

type Nodes struct {
	PrimaryLanguage struct {
		Name string
	}
	Watchers struct {
		TotalCount int
	}
	StarGazers struct {
		TotalCount int
	} `graphql:"stargazers"`
	Name      string
	ForkCount int
	Languages struct {
		TotalCount int
		Nodes      []Language
	} `graphql:"languages(first: $languageCount)"`
}

type Language struct {
	Name string
}

type ContributionCalendar struct {
	Weeks []Weeks
}

type Weeks struct {
	ContributionDays []ContributionDays
}

type ContributionDays struct {
	Date              string
	ContributionCount int
}
