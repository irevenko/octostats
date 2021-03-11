package graphql

var ContributionsQuery struct {
	User struct {
		ContributionsCollection ContributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $user)"`
}
