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

type User struct {
	Login           githubv4.String
	Name            githubv4.String
	AvatarURL       githubv4.URI
	Location        githubv4.String
	Company         githubv4.String
	Email           githubv4.String
	TwitterUsername githubv4.String
	WebsiteURL      githubv4.URI
	Bio             githubv4.String
	Status          struct {
		Emoji   string
		Message string
	}
	CreatedAt githubv4.DateTime
	Followers struct {
		TotalCount githubv4.Int
	}
	Following struct {
		TotalCount githubv4.Int
	}
	StarredRepositories struct {
		TotalCount githubv4.Int
	}
	Projects struct {
		TotalCount githubv4.Int
	}
	Packages struct {
		TotalCount githubv4.Int
	}
	Watching struct {
		TotalCount githubv4.Int
	} `graphql:"watching(privacy: PUBLIC)"`
	Gists struct {
		TotalCount githubv4.Int
	} `graphql:"gists(privacy: PUBLIC)"`
	Repositories struct {
		TotalCount int
	} `graphql:"repositories(privacy: PUBLIC)"`
	Organizations struct {
		Nodes []struct {
			Name string
		}
	} `graphql:"organizations(first: 100)"`
	SponsorshipsAsSponsor struct {
		Nodes []struct {
			Sponsorable struct {
				SponsorsListing struct {
					Slug string
				}
			}
		}
	} `graphql:"sponsorshipsAsSponsor(first: 100)"`
	SponsorshipsAsMaintainer struct {
		Nodes []struct {
			SponsorEntity struct {
				User struct{ Login string } `graphql:"... on User"`
			}
		}
	} `graphql:"sponsorshipsAsMaintainer(first: 100)"`
}

type Organization struct {
	Login           githubv4.String
	Name            githubv4.String
	AvatarURL       githubv4.URI
	Location        githubv4.String
	Email           githubv4.String
	TwitterUsername githubv4.String
	WebsiteURL      githubv4.URI
	Description     githubv4.String
	CreatedAt       githubv4.DateTime
	Projects        struct {
		TotalCount githubv4.Int
	}
	Packages struct {
		TotalCount githubv4.Int
	}
	Repositories struct {
		TotalCount int
	} `graphql:"repositories(privacy: PUBLIC)"`
	MembersWithRole struct {
		TotalCount githubv4.Int
	} `graphql:"membersWithRole(first: 100)"`
}
