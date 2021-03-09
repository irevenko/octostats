package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetMostStarredRepos(ctx context.Context, client *github.Client, allRepos []*github.Repository) ([]string, []int) {
	var starsSlice []int
	var namesSlice []string

	for _, v := range allRepos {
		stars := *v.StargazersCount
		name := *v.Name
		starsSlice = append(starsSlice, stars)
		namesSlice = append(namesSlice, name)
	}

	sorted := calcStarsOrForks(namesSlice, starsSlice)
	l, n := sortMap(sorted)
	return l, n
}
