package rest

import (
	h "../helpers"
	"github.com/google/go-github/github"
)

func LanguagesByRepo(client *github.Client, allRepos []*github.Repository) (languages []string, occurrences []int) {
	var langsSlice []string

	for _, v := range allRepos {
		if v.Language != nil {
			lang := *v.Language
			langsSlice = append(langsSlice, lang)
		} else {
			lang := "No Language"
			langsSlice = append(langsSlice, lang)
		}
	}

	mostUsedLangs := h.CountDuplicates(langsSlice)
	langs, occurrs := h.SortMap(mostUsedLangs)

	return langs, occurrs
}