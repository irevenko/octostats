package graphql

import "github.com/shurcooL/githubv4"

var ContributionsQuery struct {
	User struct {
		ContributionsCollection ContributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $user)"`
}

var YearActivityQuery struct {
	User struct {
		Repositories struct {
			Nodes []Nodes
		} `graphql:"repositories(first: $repoCount, ownerAffiliations: OWNER)"`
		ContributionsCollection struct {
			ContributionCalendar ContributionCalendar
		} `graph:"contributionsCollection"`
	} `graphql:"user(login: $user)"`
}

var UserQuery struct {
	User struct {
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
	} `graphql:"user(login: $user)"`
}

var OrganizationQuery struct {
	Organization struct {
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
	} `graphql:"organization(login: $user)"`
}
