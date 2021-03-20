package rest

import (
	"github.com/google/go-github/v33/github"
	h "github.com/irevenko/octostats/helpers"
)

func MostForkedRepos(client *github.Client, allRepos []*github.Repository) (repoNames []string, repoForks []float64) {
	var forksSlice []float64
	var namesSlice []string

	for _, v := range allRepos {
		forks := *v.ForksCount
		name := *v.Name
		forksSlice = append(forksSlice, float64(forks))
		namesSlice = append(namesSlice, name)
	}

	forksNum := h.CalcStarsOrForks(namesSlice, forksSlice)
	names, forks := h.SortMap(forksNum)
	return names, forks
}
