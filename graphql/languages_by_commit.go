package graphql

import (
	h "github.com/irevenko/octostats/helpers"
	"github.com/shurcooL/githubv4"
)

func LanguagesByCommit(client *githubv4.Client, user string, from int, to int) (languages []string, commitsNum []githubv4.Int) {
	commits := AllCommits(client, user, from, to)

	var langsSlice []string
	var numsSlice []githubv4.Int

	for i, v := range commits {
		if v.Repository.PrimaryLanguage.Name == "" {
			langsSlice = append(langsSlice, "No Language")
		} else {
			langsSlice = append(langsSlice, string(commits[i].Repository.PrimaryLanguage.Name))
		}

		numsSlice = append(numsSlice, commits[i].Contributions.TotalCount)
	}

	languagesCommit := h.CountLanguagesCommit(langsSlice, numsSlice)
	langs, count := h.SortIntv4Map(languagesCommit)

	return langs, count
}
