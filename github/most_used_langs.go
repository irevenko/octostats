package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetMostUsedLanguages(ctx context.Context, client *github.Client, allRepos []*github.Repository) (languages []string, occurences []int) {
	var languagesSlice []string

	for _, v := range allRepos {
		if v.Language != nil {
			lang := *v.Language
			languagesSlice = append(languagesSlice, lang)
		} else {
			lang := "No Language"
			languagesSlice = append(languagesSlice, lang)
		}
	}

	mostUsedLangs := countLanguages(languagesSlice)
	langs, occurs := sortMap(mostUsedLangs)

	return langs, occurs
}
