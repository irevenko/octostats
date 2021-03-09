package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetMostForkedRepos(ctx context.Context, client *github.Client, allRepos []*github.Repository) ([]string, []int) {
	var forksSlice []int
	var namesSlice []string

	for _, v := range allRepos {
		forks := *v.ForksCount
		name := *v.Name
		forksSlice = append(forksSlice, forks)
		namesSlice = append(namesSlice, name)
	}

	sorted := calcStarsOrForks(namesSlice, forksSlice)
	l, n := sortMap(sorted)
	return l, n
}
