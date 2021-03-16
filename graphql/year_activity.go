package graphql

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

func YearActivity(client *githubv4.Client, user string) (dates []string, contribs []float64) {
	variables := map[string]interface{}{
		"user":          githubv4.String(user),
		"repoCount":     githubv4.Int(100),
		"languageCount": githubv4.Int(100),
	}

	err := client.Query(context.Background(), &YearActivityQuery, variables)
	if err != nil {
		log.Fatal(err)
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

	return datesSlice, contribsSlice
}
