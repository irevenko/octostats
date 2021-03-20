package rest

import (
	"github.com/google/go-github/v33/github"
	h "github.com/irevenko/octostats/helpers"
)

func StarsPerLanguage(client *github.Client, allRepos []*github.Repository) (languages []string, starsNum []float64) {
	var starsSlice []float64
	var langsSlice []string

	for _, v := range allRepos {
		starsNum := *v.StargazersCount
		starsSlice = append(starsSlice, float64(starsNum))

		if v.Language != nil {
			lang := *v.Language
			langsSlice = append(langsSlice, lang)
		} else {
			lang := "No Language"
			langsSlice = append(langsSlice, lang)
		}
	}

	stars := h.CalcStarsOrForks(langsSlice, starsSlice)
	langs, count := h.SortMap(stars)
	return langs, count
}
