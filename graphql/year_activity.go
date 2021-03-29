package graphql

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func YearActivity(client *githubv4.Client, user string) (dates []string, contribs []float64, err error) {
	variables := map[string]interface{}{
		"user":          githubv4.String(user),
		"repoCount":     githubv4.Int(100),
		"languageCount": githubv4.Int(100),
	}

	clientErr := client.Query(context.Background(), &YearActivityQuery, variables)
	if clientErr != nil {
		err = fmt.Errorf("Couldn't get year activity for %s: %w", user, clientErr)
		return
	}

	var datesSlice []string
	var contribsSlice []float64

	for _, v := range YearActivityQuery.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, week := range v.ContributionDays {
			if week.Date != "" {
				datesSlice = append(datesSlice, week.Date)
				contribsSlice = append(contribsSlice, float64(week.ContributionCount))
			}
		}
	}

	return datesSlice, contribsSlice, nil
}
