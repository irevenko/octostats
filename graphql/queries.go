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
	} `graphql:"user(login: $user)"`
}

var OrganizationQuery struct {
	Organization struct {
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
	} `graphql:"organization(login: $user)"`
}
