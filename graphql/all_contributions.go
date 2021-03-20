package graphql

import (
	"context"
	"log"
	"time"

	"github.com/shurcooL/githubv4"
)

func AllContributions(client *githubv4.Client, user string, fromYear int, toYear int) ContributionsCollection {
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
		log.Fatal(err)
	}

	return ContributionsQuery.User.ContributionsCollection
}
