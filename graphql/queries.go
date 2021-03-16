package graphql

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
