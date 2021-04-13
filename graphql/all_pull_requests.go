package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
)

func AllPullRequests(client *githubv4.Client, user string, fromYear int, toYear int) ([]PullRequestContributions, error) {
	loc, _ := time.LoadLocation("Local")
	_, month, day := time.Now().Date()
	var m int = int(month)
	var d int = int(day)
	fromDate := time.Date(fromYear, time.Month(m), d, 0, 0, 0, 0, loc)
	toDate := time.Date(toYear, time.Month(m), d, 0, 0, 0, 0, loc)

	variables := map[string]interface{}{
		"user": githubv4.String(user),
		"from": githubv4.DateTime{Time: fromDate},
		"to":   githubv4.DateTime{Time: toDate},
	}

	err := client.Query(context.Background(), &ContributionsQuery, variables)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get pull requests for %s: %w", user, err)
	}

	return ContributionsQuery.User.ContributionsCollection.PullRequestContributionsByRepository, nil
}
