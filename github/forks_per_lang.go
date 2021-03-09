package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetForksPerLanguage(ctx context.Context, client *github.Client, allRepos []*github.Repository) ([]string, []int) {
	var forksSlice []int
	var languagesSlice []string

	for _, v := range allRepos {
		forks := *v.ForksCount
		forksSlice = append(forksSlice, forks)

		if v.Language != nil {
			lang := *v.Language
			languagesSlice = append(languagesSlice, lang)
		} else {
			lang := "No Language"
			languagesSlice = append(languagesSlice, lang)
		}
	}

	forks := calcStarsOrForks(languagesSlice, forksSlice)
	l, n := sortMap(forks)
	return l, n
}
