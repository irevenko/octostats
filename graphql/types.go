package graphql

import "github.com/shurcooL/githubv4"

type Contributions struct {
	TotalCount int
}

type Repository struct {
	NameWithOwner   string
	PrimaryLanguage struct {
		Name string
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

type User struct {
	Login           string
	Name            string
	AvatarURL       string
	Location        string
	Company         string
	Email           string
	TwitterUsername string
	WebsiteURL      string
	Bio             string
	Status          struct {
		Emoji   string
		Message string
	}
	CreatedAt githubv4.DateTime
	Followers struct {
		TotalCount int
	}
	Following struct {
		TotalCount int
	}
	StarredRepositories struct {
		TotalCount int
	}
	Projects struct {
		TotalCount int
	}
	Packages struct {
		TotalCount int
	}
	Watching struct {
		TotalCount int
	} `graphql:"watching(privacy: PUBLIC)"`
	Gists struct {
		TotalCount int
	} `graphql:"gists(privacy: PUBLIC)"`
	Repositories struct {
		TotalCount int
	} `graphql:"repositories(privacy: PUBLIC)"`
	Organizations struct {
		TotalCount int
	} `graphql:"organizations"`
	SponsorshipsAsSponsor struct {
		TotalCount int
	} `graphql:"sponsorshipsAsSponsor"`
	SponsorshipsAsMaintainer struct {
		TotalCount int
	} `graphql:"sponsorshipsAsMaintainer"`
}

type Organization struct {
	Login           string
	Name            string
	AvatarURL       string
	Location        string
	Email           string
	TwitterUsername string
	WebsiteURL      string
	Description     string
	CreatedAt       githubv4.DateTime
	Projects        struct {
		TotalCount int
	}
	Packages struct {
		TotalCount int
	}
	Repositories struct {
		TotalCount int
	} `graphql:"repositories(privacy: PUBLIC)"`
	MembersWithRole struct {
		TotalCount int
	} `graphql:"membersWithRole(first: 100)"`
	SponsorshipsAsSponsor struct {
		TotalCount int
	} `graphql:"sponsorshipsAsSponsor"`
	SponsorshipsAsMaintainer struct {
		TotalCount int
	} `graphql:"sponsorshipsAsMaintainer"`
}
