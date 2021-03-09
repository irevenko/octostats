package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetStarsPerLanguage(ctx context.Context, client *github.Client, allRepos []*github.Repository) ([]string, []int) {
	var starsSlice []int
	var languagesSlice []string

	for _, v := range allRepos {
		stars := *v.StargazersCount
		starsSlice = append(starsSlice, stars)

		if v.Language != nil {
			lang := *v.Language
			languagesSlice = append(languagesSlice, lang)
		} else {
			lang := "No Language"
			languagesSlice = append(languagesSlice, lang)
		}
	}

	stars := calcStarsOrForks(languagesSlice, starsSlice)
	l, n := sortMap(stars)
	return l, n
}
