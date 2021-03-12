package rest

import (
	"github.com/google/go-github/v33/github"
	h "github.com/irevenko/octostats/helpers"
)

func MostStarredRepos(client *github.Client, allRepos []*github.Repository) (repoNames []string, repoStars []int) {
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
