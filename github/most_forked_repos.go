package github

import (
	"context"

	h "../helpers"
	"github.com/google/go-github/github"
)

func MostForkedRepos(ctx context.Context, client *github.Client, allRepos []*github.Repository) (repoNames []string, repoForks []int) {
	var forksSlice []int
	var namesSlice []string

	for _, v := range allRepos {
		forks := *v.ForksCount
		name := *v.Name
		forksSlice = append(forksSlice, forks)
		namesSlice = append(namesSlice, name)
	}

	forksNum := h.CalcStarsOrForks(namesSlice, forksSlice)
	names, forks := h.SortMap(forksNum)
	return names, forks
}
