package github

import (
	"context"

	h "../helpers"
	"github.com/google/go-github/github"
)

func MostStarredRepos(ctx context.Context, client *github.Client, allRepos []*github.Repository) (repoNames []string, repoStars []int) {
	var starsSlice []int
	var namesSlice []string

	for _, v := range allRepos {
		stars := *v.StargazersCount
		name := *v.Name
		starsSlice = append(starsSlice, stars)
		namesSlice = append(namesSlice, name)
	}

	starsNum := h.CalcStarsOrForks(namesSlice, starsSlice)
	names, stars := h.SortMap(starsNum)
	return names, stars
}
