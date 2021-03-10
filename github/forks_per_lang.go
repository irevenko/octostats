package github

import (
	"context"

	h "../helpers"
	"github.com/google/go-github/github"
)

func ForksPerLanguage(ctx context.Context, client *github.Client, allRepos []*github.Repository) (languages []string, forksNum []int) {
	var forksSlice []int
	var langsSlice []string

	for _, v := range allRepos {
		forksNum := *v.ForksCount
		forksSlice = append(forksSlice, forksNum)

		if v.Language != nil {
			lang := *v.Language
			langsSlice = append(langsSlice, lang)
		} else {
			lang := "No Language"
			langsSlice = append(langsSlice, lang)
		}
	}

	forks := h.CalcStarsOrForks(langsSlice, forksSlice)
	langs, count := h.SortMap(forks)
	return langs, count
}
