package github

import (
	"context"

	h "../helpers"
	"github.com/google/go-github/github"
)

func StarsPerLanguage(ctx context.Context, client *github.Client, allRepos []*github.Repository) (languages []string, starsNum []int) {
	var starsSlice []int
	var langsSlice []string

	for _, v := range allRepos {
		starsNum := *v.StargazersCount
		starsSlice = append(starsSlice, starsNum)

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
