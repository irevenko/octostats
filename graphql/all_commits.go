package graphql

import (
	"context"
	"log"
	"time"

	"github.com/shurcooL/githubv4"
)

func AllCommits(client *githubv4.Client, user string, from int, to int) []CommitContributions {
	loc, _ := time.LoadLocation("Local")
	fromDate := time.Date(from, time.Month(1), 1, 0, 0, 0, 0, loc)
	toDate := time.Date(to, time.Month(1), 1, 0, 0, 0, 0, loc)

	variables := map[string]interface{}{
		"user": githubv4.String(user),
		"from": githubv4.DateTime{Time: fromDate},
		"to":   githubv4.DateTime{Time: toDate},
	}
	err := client.Query(context.Background(), &ContributionsQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return ContributionsQuery.User.ContributionsCollection.CommitContributionsByRepository
}
