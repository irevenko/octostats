package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
)

func AllContributions(client *githubv4.Client, user string, fromYear int, toYear int) (c ContributionsCollection, err error) {
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

	clientErr := client.Query(context.Background(), &ContributionsQuery, variables)
	if clientErr != nil {
		return c, fmt.Errorf("Couldn't get contributions for %s: %w", user, clientErr)
	}

	return ContributionsQuery.User.ContributionsCollection, nil
}
