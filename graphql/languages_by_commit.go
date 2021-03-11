package graphql

import (
	"fmt"

	h "../helpers"
	"github.com/shurcooL/githubv4"
)

func LanguagesByCommit(client *githubv4.Client, user string, from int, to int) {
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

	fmt.Println(langsSlice)
	fmt.Println(numsSlice)

	languagesCommit := h.CountLanguagesCommit(langsSlice, numsSlice)

	fmt.Println(languagesCommit)
	//langs, count := h.SortMap(languagesCommit)
	// cast githubv4.Int to normal int
	//return langs, count

	//fmt.Println(h.CountIntDuplicates())
}
